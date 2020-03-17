// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main_test

import (
	"fmt"
	"testing"
)

type A struct {
	a int
	i
}

type B struct {
	a int
	i
}

type C struct {
	A
	B
}

func (a A) get() {

}

func (a B) get() {

}

type i interface {
	get()
}

func Test_匿名组合(t *testing.T) {
	//c := C{}
	//// 编译错误
	//c.get()
	//// 编译错误
	//var d i = C{}
	//d.get()
}

func Test_匿名组合1(t *testing.T) {

	// 无编译错误，但会断言失败
	var c interface{}
	c = C{}
	d, ok := c.(i)
	fmt.Println(ok, d)
	//d.get()
}
