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
	handleAPI    ziface.HandFunc
	ExitBuffChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI ziface.HandFunc) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		handleAPI:    callbackAPI,
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
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Recv buf err:", err)
			return
		}

		if err := c.handleAPI(c.Conn, buf[:cnt], cnt); err != nil {
			fmt.Println("ConnID", c.ConnID, "Handle func error")
			return
		}
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
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
