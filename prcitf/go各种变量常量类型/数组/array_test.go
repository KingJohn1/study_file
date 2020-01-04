// Copyright (c) liuhao. 2012-2050. All rights reserved.
package 数组

import (
	"fmt"
	"testing"
)

func Test_Init(t *testing.T) {
	// 数组只能有常量表达式初始化，并且同长度类型 可以直接比较值
	var a = 10
	var b [10]int
	var c [10]int
	var d [10]int
	b[1] = 1
	c[1] = 1
	fmt.Println(a, b == c, c == d)
}

// todo 注意数组与切片的传递规则
func Test_数组是是传递值而非传递指针(t *testing.T) {
	var a [10]int
	a[1] = 1
	testPoint(a)
	fmt.Println(a)
}
func testPoint(a [10]int) {
	a[1] = 2
	fmt.Println(a)
}
