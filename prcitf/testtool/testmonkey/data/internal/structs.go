// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package internal

import "fmt"

type S struct {

}

var a=1

type SI interface {
	Get()
}

func(s *S)Get(){
	fmt.Println("origin get",a)
}