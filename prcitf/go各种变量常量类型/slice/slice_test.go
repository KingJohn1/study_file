// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package slice

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"log"
	"math"
	"testing"
	"unsafe"
)

var hasPanic = func() {
	if err := recover(); err != nil {
		fmt.Println("has panic ")
	} else {
		log.Fatal("no panic,excepet has panic")
	}
}

func Test_SliceIndexType(t *testing.T) {

	convey.Convey(
		"len(a)-1 可能会panic", t, func() {
			defer hasPanic()
			var a = make([]int, 0)
			fmt.Println(a[len(a)-1])
		},
	)
	convey.Convey(
		"range 的 index 类型", t, func() {
			var a = make([]int, 1)
			for i := range a {
				convey.So(fmt.Sprintf("%T", i), convey.ShouldEqual, "int")
			}
		},
	)
	convey.Convey(
		"slice index type 可以是int也可以是uint8 uint32等 ", t, func() {
			var index int
			var index1 uint8
			var a = make([]int, math.MaxUint8)
			a[index] = 1
			a[index1] = 2
			fmt.Printf("%T  ,%[1]v  ", index1-1)
			var b int
			// index1-b uint8与int 无法运算，index1-1 可以运算运算结果index1的类型uint8；理解下字面常量为可变的万能整数 类型
			fmt.Printf("%T  ,%T  ,%T ", 1, index-b, index-1)
			a[index1-2] = 1
		},
	)
}

func Test_SliceLenCap(t *testing.T) {

	convey.Convey(
		"index 超过 len 小于 cap  会panic", t, func() {
			defer hasPanic()
			var a = make([]int, 1, 10)
			a[2] = 1
		},
	)
	convey.Convey(
		"通过旧切片生成新切片，新切片的容量不在len内在cap内  不会panic", t, func() {
			var a = make([]int, 5, 10)
			_ = a[4:9]
		},
	)
	convey.Convey(
		"通过旧切片生成新切片，新切片的容量不是新切片的len，而是新切片的第一个元素位置到底层数组cap的最后一个位置的长度", t, func() {
			var a = make([]int, 5, 10)
			var index1 = 1
			b := a[index1:]
			convey.So(len(b), convey.ShouldEqual, len(a)-index1)
			convey.So(cap(b), convey.ShouldNotEqual, len(b))
			convey.So(cap(b), convey.ShouldEqual, cap(a)-index1)
		},
	)
}

// gophere 2018 滴滴 中说 数组可以省对象分配，所以测试下区别
func Test_切片和数组内存布局是否相同(t *testing.T) {
	convey.Convey(
		"数组取地址直接得到第一个数字，具体原因还得什么了解", t, func() {
			var slice = make([]byte, 5, 10)
			slice[0] = 1
			var array = [5]byte{1, 2, 3, 4, 5,}
			// 0xc000048680 0xc0000540f0
			//0xc000066150 0xc000066150
			// todo 切片内存布局 data len cap；为和array 直接指向了数据 类似c、c++ 中的数组内存结构
			fmt.Println((unsafe.Pointer)(&slice), (unsafe.Pointer)(&slice[0]))
			fmt.Println((unsafe.Pointer)(&array), (unsafe.Pointer)(&array[0]))
			fmt.Println(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&array)) + 5)))
		},
	)
}

func TestRange(t *testing.T) {
	convey.Convey(
		"range的语法是只确认一个范围，不会动态根据切片内容改变长度，所以range时不要增删元素", t, func() {
			defer hasPanic()
			var a = []int{1, 2, 3}
			for i := range a {
				if a[i] == 1 {
					fmt.Println("n")
					continue
				}
				a = nil
			}
		},
	)
}

func Test_Append(t *testing.T) {
	convey.Convey("append 函数没有左值 会语法错误。append函数不会修改入参", t, func() {
		var a = make([]int, 1, 10)
		a[0] = 1
		b := append(a, 2)
		fmt.Println(a, b)
		//[1] [1 2]
	})

	convey.Convey("", t, func() {
		var a = make([]int, 0, 10)
		var f = func(a []int) {
			a = append(a, 1)
			fmt.Println(a)
		}
		f(a)
		fmt.Println(a)
	})
}

func Test_Print(t *testing.T) {
	var a = []byte{10, 2, 18, 16}
	fmt.Printf("%# X", a) //0X0A 0X02 0X12 0X10
}

func Test_三个点(t *testing.T) {
	var a []int
	var b = make([]int, 0, 0, )
	f := func(i ...int) {
		for _, v := range i {
			fmt.Println(v)
		}
	}
	f(a...)
	f(b...)
	a=append(a,1)
	fmt.Println()
	f(a...)
	fmt.Printf("%#x",281474976710658)
}
