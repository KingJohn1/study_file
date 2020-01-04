// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _15_计算右侧小于当前元素的个数

import (
	"fmt"
	"github.com/prashantv/gostub"
	"testing"
)

func TestName(t *testing.T) {
	gostub.Stub(&useRecursive,true)
	fmt.Println(dontUseRecursiveProcess([]int{0,2,1}))
	a:=complex(1,1)
	b:=complex(1,0)
	c:=a*b
	fmt.Printf("%v",c)
}

func TestBst(t *testing.T) {
	gostub.Stub(&useRecursive,true)
	fmt.Println(countSmaller1([]int{2,2,1}))
}

func TestName1(t *testing.T) {
	fmt.Println(useRecursive)
}

func tt(data []int)(result []int){
	defer func() {
		fmt.Println(data,result)
	}()
	return data
}

func TestName2(t *testing.T) {
fmt.Println(tt([]int{1}))
}
