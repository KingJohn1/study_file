// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package testconvey

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net"
	"testing"
	"time"
)


func TestStringSliceEqual(t *testing.T) {
	Convey("test liuhao1",t, func() {
		a:=[]string{"liuhao"}
		b:=[]string{"liu"}
		So(StringSliceEqual(a,b),ShouldBeFalse)
		So(StringSliceEqual(a,b),ShouldBeTrue)
	})
	Convey("TestStringSliceEqual should return false when a ＝= nil  && b != nil", t, func() {
		a := []string(nil)
		b := []string{}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})

}
var b =[]string{"liu"}

func TestStringSliceEqual_Stub(t *testing.T) {
	//defer
	var ip net.IP
	ip=[]byte{1}
	fmt.Println(ip)
	var bb []byte
	bb=net.IP{1}
	fmt.Println(bb)
	//fmt.Println(ip==bb)
	Convey("test liuhao1",t, func() {
		a:=[]string{"liuhao"}

		So(StringSliceEqual(a,b),ShouldBeFalse)
		So(StringSliceEqual(a,b),ShouldBeTrue)
	})

}
func TestA(t*testing.T){

	var b B
	c:=b.a
	c()
	c()
	d:=(*B).a
	d(&b)

	fmt.Printf("type :%T",c)

}

func TestBB(t *testing.T){
	var d func()
	a:=func ()func(){
	b:=1
d=func(){
		b++
		fmt.Println(b)
	}
return d
	}
	//d()
	a1:=a()
	a1()
	a1()
	a2:=a()
	a2()
	a2()
	a1()

}

type B struct {
b int
}

func(b *B) a(){
	b.b++
fmt.Printf("liu b:%v \n",b.b)
}

func TestB(t *testing.T){
	Convey(
		"callback的执行顺序",
		t,

		func() {
		var a =func(){
			time.Sleep(500*time.Millisecond)
			fmt.Println("sleep 500")
		}

		a()
		fmt.Println("stop")
		},

		)
}



type II struct {

}

func (a II)get(){

}

func (a II)Get(){

}

func TestC(t *testing.T){

}