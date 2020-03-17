// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package testmonkey

import (
	"fmt"
	"github.com/bouk/monkey"
	"reflect"
	"testing"
)

func Test(t *testing.T){
var ss S
	guard:= monkey.PatchInstanceMethod(reflect.TypeOf(ss),"Get", func( _ S){
		fmt.Println("liuhao replace")
	})
	defer guard.Unpatch()

	var s =&S{

	}
	s.Get()



//	fmt.Println(reflect.ValueOf(ss),reflect.TypeOf(ss))

}

func tt( s SS){

}

type S struct {
 SS
}

type SS struct {

}
func(s SS)Get(){
	fmt.Println("get")
}