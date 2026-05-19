package znet

import (
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

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("[START] Error accepting tcp connection: %s\n", err)
			}

			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Printf("[START] Error reading tcp connection: %s\n", err)
						continue
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Printf("[START] Error writing tcp connection: %s\n", err)
						continue
					}
				}
			}()
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
	}
	return s
}
