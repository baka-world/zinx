package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	fmt.Println("TestClient start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("TestClient Dial err:", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello ZINX"))
		if err != nil {
			fmt.Println("TestClient Write err:", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("TestClient Read err:", err)
			return
		}
		fmt.Printf("Server call back: %s, cnt = %d \n", string(buf[:cnt]), cnt)
		time.Sleep(time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer("[Zinx V0.1]")

	go ClientTest()
	s.Serve()
}
