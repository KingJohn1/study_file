// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main(){

	var ss=make([]byte,100)
	ss[0]='t'
	fmt.Println(string(ss),len(bytes.TrimSpace(ss)),string(bytes.TrimFunc(ss, func(r rune) bool {
		if r==0{
			return true
		}
		return false
	},
	),
	),

	len(bytes.TrimFunc(ss, func(r rune) bool {
		if r==0{
			return true
		}
		return false
	},
	),

		),
	)

	d := flag.Bool("d", false, "Whether or not to launch in the background(like a daemon)")
	flag.Parse()
	if *d {
		cmd := exec.Command("go",
			//"-close-fds",
			//"-addr", *addr,
			//"-call", *call,
		)
		serr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatalln(err)
		}
		err = cmd.Start()
		if err != nil {
			log.Fatalln(err)
		}
		s, err := ioutil.ReadAll(serr)
		s = bytes.TrimSpace(s)
		if bytes.HasPrefix(s, []byte("addr: ")) {
			fmt.Println(string(s))
			cmd.Process.Release()
		} else {
			log.Printf("unexpected response from MarGo: `%s` error: `%v`\n", s, err)
			cmd.Process.Kill()
		}
	}
}