// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package data

import (
	"fmt"
	"prcitf/testtool/testmonkey/data/internal"
)

func A(){
	//var s =&data.S{}
	//s.Get()
	//
	//var si *data.SI=new(data.SI)
	//var ss= SS{
	//
	//}
	//ss.Get()
	//*si=s
	var s=&SS{}
	s.Get()
}

type SS struct {
	internal.S
}
func (ss *SS)Pet(){
	fmt.Println("pet")
}

type SSS struct {

}
func (s *SSS)Get(){
	fmt.Println("get")
}