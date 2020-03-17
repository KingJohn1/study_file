// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"net"
)
var udpServerAddr, err = net.ResolveUDPAddr("udp", ":120")
func main (){

	fmt.Println(udpServerAddr)
	udpConn, err := net.ListenUDP("udp", udpServerAddr)
	if err!=nil{
		fmt.Println(err)
		return
	}

	for {
		var buf [512]byte
		readFromUDP, addr, err := udpConn.ReadFromUDP(buf[0:])
		fmt.Println("收到请求",readFromUDP,addr,err)
		i, err := udpConn.WriteToUDP([]byte("欢迎连接服务器"),addr)
		fmt.Println(i,err)
	}

}
