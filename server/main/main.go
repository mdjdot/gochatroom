package main

import (
	"fmt"
	"github.com/mdjdot/gochatroom/server/processes"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8321")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	fmt.Println("服务端端口 8321 开始等待请求...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("net.Listen.Accept err=", err)
			return
		}
		processes.ProcessConn(conn)
	}
}
