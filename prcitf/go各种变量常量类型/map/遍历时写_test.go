// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"fmt"
	"testing"
)

// map在range中增加元素或写元素 range 的len也会改变,但不一定是实际值
func Test_遍历时写(t *testing.T) {
	var m = make(map[int]int)

	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4
	j:=4
	for i, v := range m {
		j++
		m[j]=j
		if j==100{
			break
		}
		fmt.Println(i, v)
	}
}
