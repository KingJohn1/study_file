// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"reflect"
)

func main(){
	var a=struct{
		aa int
	}{1}
	a.aa=1

	fmt.Println(a,reflect.TypeOf(a).PkgPath(),reflect.TypeOf(a).PkgPath()=="",reflect.TypeOf(a).Name())
	fmt.Println(reflect.TypeOf(a).String())

	var A = struct {
		bb int
	}{1}
	A.bb=1
	fmt.Println(A,reflect.TypeOf(A).PkgPath(),reflect.TypeOf(A).PkgPath()=="")

	var b struct{
		aa int
	}= struct{ aa int }{aa: 2}
b.aa=1
}
