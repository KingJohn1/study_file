// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _27_基本计算器2

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//calculate("1+1+2+3*4/5 + 6+ 7- 3/2(()())")
}

func TestName1(t *testing.T) {
	var data = "1+2+(2*2*(2+1))+1"
//	fmt.Print(splitString(data))
	fmt.Println(calculate(data))
}

func Test1(t *testing.T) {
	var s *Stack = &Stack{}
	var ss *([]int)
	*ss = append(*ss, 1)
	s.isEmpty()
	s.push("1")
	fmt.Println(s)
}
