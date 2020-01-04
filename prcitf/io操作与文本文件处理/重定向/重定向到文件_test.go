// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 重定向

import (
	"fmt"
	"os"
	"prcitf/util"
	"syscall"
	"testing"
)

func TestRedirect(t *testing.T) {
	redirectDst,err := getFile("liuhao.txt")
	if err!=nil{
		t.Error("getfile err",redirectDst)
	}
	os.Stdout = redirectDst
	util.SleepLit()
	fmt.Println("lalal")
	os.Stdout=os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	//t.SkipNow()
}

func getFile(s string) (*os.File,error) {
	file, err := os.OpenFile(s, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	if err == nil {
		file.Write([]byte("begin1"))
		file.WriteAt([]byte("下面是日1志\n"),0)
		return file,nil
	}
	fmt.Println(err)

	createFile, err := os.Create(s)
	if err==nil{
		createFile.Write([]byte("begin"))
		createFile.WriteAt([]byte("下面是日志\n"),20)
		return createFile,nil
	}
	fmt.Println("err happen",err)
	return nil,err
}
