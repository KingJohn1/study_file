// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"net"
	"net/rpc"
	websever "prcitf/网络/websever/rpc/internal"
)

func main(){

	arith:=new(websever.Arith)
	rpc.Register(arith)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":555")
	if err!=nil{
		log.Fatal(err.Error())
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	for{
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err.Error())
		}
		go func(){
			rpc.ServeConn(conn)
		}()
	}

}
