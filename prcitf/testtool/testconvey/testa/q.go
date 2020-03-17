// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
)

type a struct {

}

func NewA()*a{
	return &a{}
}

type I interface {
//	get()
	Get()
}

func (a *a)Get(){
	fmt.Println("lalala")
}

func main(){

	//fmt.Println( reflect.TypeOf(NewA()))
	//
	//e:=NewA()
	////b:=reflect.TypeOf(e)
	//
	//a:=monkey.PatchInstanceMethod(reflect.TypeOf(e),"Get",func(_ interface{}){
	//	fmt.Println("liuhao")
	//})
	//e.Get()

	 var a uint =1

	 const b=1
	switch a {
	case b:fmt.Println(b)
	fmt.Printf("%T",b)
	}

	//defer a.Unpatch()
//
//a:="12"
//fmt.Println(a*2)
}