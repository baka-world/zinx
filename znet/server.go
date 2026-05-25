package znet

import (
	"errors"
	"fmt"
	"net"
	"time"
	"zinx/ziface"
)

type Server struct {
	Name string

	IPVersion string

	IP string

	Port int

	Router ziface.IRouter
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("server add router success!")
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	// Callback display
	fmt.Printf("[Conn Handle] CallbackToClient ...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("[Conn Handle] CallbackToClient Write Error", err)
		return errors.New("write Error")
	}
	return nil
}
func (s *Server) Start() {
	fmt.Printf("[START] %s Server listener at IP: %s, Port: %d, is starting\n", s.Name, s.IP, s.Port)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Printf("[START] Error resolving tcp address: %s\n", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("[START] Error starting tcp listener: %s\n", err)
			return
		}

		fmt.Printf("[START] Server listening at IP: %s, Port: %d\n", s.IP, s.Port)

		// TODO Should autogen ID
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("[START] Error accepting tcp connection: %s\n", err)
			}

			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Printf("[STOP] %s Server listener at IP: %s, Port: %d\n", s.Name, s.IP, s.Port)
}

func (s *Server) Serve() {
	s.Start()

	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
		Router:    nil,
	}
	return s
}
