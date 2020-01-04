// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
	_ "net/http/pprof"
	"net/rpc"
	 "prcitf/网络/websever/rpc/internal"
)

func main() {
	arith := new(websever.Arith)

	err := rpc.Register(arith)
	if err != nil {
		fmt.Println(err)
		return
	}
	rpc.HandleHTTP()
	// 在go的net包中，监听的可以是地址也可是端口；注册的handle也可以是地址加路径也可以是地址
	err = http.ListenAndServe(websever.Addr, nil)
	if err!=nil{
		log.Fatal("err:%v ",err)
	}
}
