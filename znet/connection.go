package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn         *net.TCPConn
	ConnID       uint32
	isClosed     bool
	Router       ziface.IRouter
	ExitBuffChan chan bool
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader goroutine is starting ...")
	defer fmt.Println("Reader goroutine is shutting down ...")
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Recv buf err:", err)
			return
		}

		req := Request{
			conn: c,
			data: buf,
		}

		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}
func (c *Connection) Start() {
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//TODO Close Callback

	// Close socket link
	c.Conn.Close()

	// Notify other reader thread that Chan is closed
	c.ExitBuffChan <- true

	// Release channel
	close(c.ExitBuffChan)
}

func (c *Connection) GetConnection() net.TCPConn {
	return *c.Conn
}
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
