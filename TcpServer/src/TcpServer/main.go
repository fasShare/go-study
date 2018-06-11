package main

import (
	"fmt"
	"mtcp"
	"net"
	"time"
)

func NewConnection(conn net.Conn) {
	buffer := make([]byte, 512)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			conn.Close()
			fmt.Println("read:", err)
			break
		}
		// Write出错会返回什么内容，没有测试出来
		if _, err = conn.Write(buffer[:n]); err != nil {
			conn.Close()
			fmt.Println("write:", err)
		}
	}
}

func main() {
	fmt.Println("start main!")
	server := mtcp.TcpServer{}
	if err := server.Start("127.0.0.1:8899", NewConnection); err != nil {
		fmt.Println(err)
	}
	fmt.Println("exit main!")
}
