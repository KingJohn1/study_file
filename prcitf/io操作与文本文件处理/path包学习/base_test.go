// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package path包学习

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func Test_PathSplit(t *testing.T) {
	dir, file := path.Split("io操作与文本文件处理/path包学习/base_test.go")
	fmt.Println(dir,file)

	open, err := os.Open(`D:\GO\src\prcitf\io操作与文本文件处理\path包学习\base_test.go`)
	defer open.Close()
	info, err := os.Stat(`\GO\src\prcitf\io操作与文本文件处理\path包学习\base_test.go`)
	fmt.Println("name",info.Name())

	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("name",open.Name())
}