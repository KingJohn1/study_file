// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"net"
	"strings"
)

var clientAddr, _ = net.ResolveUDPAddr("udp", ":110")

func main() {
	udpServerAddr, _ := net.ResolveUDPAddr("udp", "10.151.73.175:120")
	udpConn, _ := net.DialUDP("udp", clientAddr, udpServerAddr)
	udpConn.Write([]byte("请求连接"))
	fmt.Println("请求连接")

	var buff [1111]byte
	udpConn.Read(buff[0:])

	s := strings.Replace(string(buff[0:]), string(byte(0)), "", -1)
	fmt.Println(string(buff[0:]))
	fmt.Println(s,len(s))
}
