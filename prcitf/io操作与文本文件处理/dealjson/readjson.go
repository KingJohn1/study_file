// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	ServerName string
	ServerIP   string
	JJ         string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","Jj":"liu" ,"serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Printf("%#v ", s)
	fmt.Println()

	type Server struct {
		// ID 不会导出到JSON中
		ID int `json:"-"`

		// ServerName 的值会进行二次JSON编码
		ServerName  string `json:"serverName"`
		ServerName2 int    `json:"serverName2,string"`

		// 如果 ServerIP 为空，则不输出到JSON串中
		ServerIP string `json:"serverIP,omitempty"`
	}

	ss := Server{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: 12333,
		ServerIP:    ``,
	}
	b, _ := json.Marshal(ss)
	os.Stdout.Write(b)
	fmt.Println(string(b[0]), string(b[1]))
	fmt.Println()

}
