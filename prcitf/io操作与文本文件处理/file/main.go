// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"os"
)

func main(){

	os.MkdirAll("liuhao_test/liu/liu/liu",0777)
	err:=os.RemoveAll("liuhao_test")
	fmt.Println(err)
	userFile := "astaxie.txt"
	defer func() {
		os.RemoveAll(userFile)
	}()
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}

	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}

}
