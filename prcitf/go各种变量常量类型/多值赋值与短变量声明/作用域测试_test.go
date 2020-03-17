// Copyright (c) liuhao. 2012-2050. All rights reserved.
package 多值赋值与短变量声明

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_多值赋值(t *testing.T) {

	// 如果变量需要不同的类型要么隐私声明 要多写成多行
	var a, b = 1, "2312"
	var (
		c int    = 1
		d string = "213"
	)
	fmt.Println(a, b, c, d)
	// var a1 int,b1 int //编译错误

}

func Test_利用多值赋值声明变量(t *testing.T) {
	convey.Convey("短变量声明 左侧有两个变量，只要有一个未定义即有效", t, func() {
		var a = 1
		// 其实只声明了b
		a, b := 2, 2
		convey.So(fmt.Sprint(a, b), convey.ShouldEqual, "2 2")
	})

	convey.Convey("短变量声明 左侧有多个变量，总是尽量定义多个变量", t, func() {
	var a=1
	{
		a,b:=2,3
		convey.So(fmt.Sprint(a,b),convey.ShouldEqual,"2 3")
	}
	convey.So(fmt.Sprint(a),convey.ShouldEqual,"1")
	})
}
