// Copyright (c) liuhao. 2012-2050. All rights reserved.
package 类型重定义

import (
	"fmt"
	"testing"
)

type A struct {
}

func (a A) print() {
	fmt.Println("i am is A")
}

type B A

func Test_根据已有类型新生成一个类型(t *testing.T) {
	var b B
	// B 不集成A的方法
	b.print()
	A(b).print()
}
