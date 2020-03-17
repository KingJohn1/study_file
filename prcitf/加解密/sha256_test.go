// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 加解密

import (
	"crypto/sha256"
	"fmt"
	"io"
	"testing"
)

func Test_Usage(t*testing.T){
	hash := sha256.New()
	hash.Write([]byte("liuhao"))
	fmt.Println(string( hash.Sum(nil)))

	hash1:=sha256.New()
	io.WriteString(hash1,"liuhao")
	var b =make([]byte,500)
	fmt.Println  (string( hash1.Sum(b)),string( b))
}
