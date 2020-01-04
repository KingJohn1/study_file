// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":120")
	clientAddr, err := net.ResolveTCPAddr("tcp4", ":121")
	fmt.Println(tcpAddr, err)
	if err != nil {
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp4", clientAddr, tcpAddr)
	fmt.Println(conn, err)
	fmt.Println("连接")
	rsp := make([]byte, 111, 128)
	rsp = []byte(strings.Repeat("", 10))
	conn.Read(rsp)

	result, err := ioutil.ReadAll(conn)
	fmt.Println(string(result), err, string(rsp))
}
