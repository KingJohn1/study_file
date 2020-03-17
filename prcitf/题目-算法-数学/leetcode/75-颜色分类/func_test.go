// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _5_颜色分类

import (
	"fmt"
	"testing"
)

func Test_SortColors(t *testing.T) {
	var data = []int{1, 1, 1, 2, 0, 2, 1, 1, 0, 1, 2, 1, 1}
	sortColors(data)
	fmt.Println(data)
	tt= func() int {
		return 2
	}

	fmt.Println(tt())

}

var tt=func ()int{
	return 1
}

func decorate(f func(a int)) func(a int) {
	b := func(a int) {
		// func
		f(a)
	}
	return b
}
