// Copyright (c) liuhao. 2012-2050. All rights reserved.
package 取余mod

import (
	"fmt"
	"testing"
)

func add(a, b, c int) bool {

	r1:= (a+b)%c
	r2:=((a%c)+(b%c))%c
	if r1!=r2{
		fmt.Println(r1,r2,a,b,c,a%c,b%c,(a%c)+(b&c))
		return false
	}
	return true
}

var data = [][]int{
	{
		1, 2, 3,
	},
	{
		11, 2, 5,
	},
	{
		1234123, 13412, 12432,
	},
}

// 注意减法时 得到结果可能小于0
func Test_加法(t *testing.T) {
	// (a+b) mod c==((a mod c)+(b mod c)) mod c
	for i, v := range data {
		if ! add(v[0], v[1], v[2]) {
			t.Error(i)
		}
	}
}

func multi(a, b, c int) bool {
	return (a*(b%c))%c == (a*b)%c
}

func Test_乘法(t *testing.T) {
	for i, v := range data {
		if !multi(v[0], v[1], v[2]) {
			t.Error(i)
		}
	}
}
