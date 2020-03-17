package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":120"
	serverAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println(serverAddr, err)
	if err != nil {
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp4", serverAddr)
	if err != nil {
		fmt.Println("listen tcp err,err:", err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		fmt.Println("listener accept one conn")
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		go func() {
			fmt.Println("开始返回输出")
			daytime := time.Now().String()
			n, e := conn.Write([]byte(daytime))
			fmt.Println(n, e)
			conn.Close()
		}()
	}
}
