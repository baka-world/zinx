package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client is Starting...")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Connect Error", err)
		return
	}
	for {
		_, err := conn.Write([]byte("hello Zinx"))
		if err != nil {
			fmt.Println("Write Error", err)
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read Error", err)
		}
		fmt.Printf("Server callback: %s, cnt: %d\n", buf[:cnt], cnt)
		time.Sleep(time.Second)
	}

}
