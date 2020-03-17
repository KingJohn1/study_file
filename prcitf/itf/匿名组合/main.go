// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import "fmt"

type A struct {
}

type B struct {
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

func main() {

	// 无编译错误，但会断言错误
	var c interface{}
	c = C{}
	_, ok := c.(i)
	fmt.Println(ok)
	//d.get()
}

