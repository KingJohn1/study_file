package main

import (
	"fmt"
)

type A struct {
 	a int
 	b string
 	B
 }

type BB interface {
	get()
}

 type B struct {
	 a int
	 b string
 }
func (b B)get(){
	b.a=1
	fmt.Println("hello ")
}

//func(a A)String()string{
//	return "liuhoa"
//}

func get()*B{
	return nil
}

 func main(){
//var a *A
//fmt.Printf("%T,%#[1]v",a)
// 	fmt.Println(a)
var a BB
a=get()
fmt.Println(a,a==nil)
fmt.Printf("%T %#[1]v %[1]v,%#[2]x",a,11)


//a:=A{}
//a.get()
//
//fmt.Println(a,a.B)

// 	a1:=A{
// 		1,"a1",
//	}
//	b1:=B{1,"a1"}
//fmt.Println(a1==A(b1))//output:true
//var a int =1
//for i:=1;i<3;i++{
//	switch  a {
//
//	case 1:
//		fmt.Println(222)
//		break
//		fmt.Println("3333")
//		break
//		fmt.Println("111")
//	case 2:
//	}
//}
//
//
//b2:=tmp.A{}
//var d interface{}=b2
//c:=d.(tmp.Itf)
//fmt.Println(c,"dddddddd")



 }