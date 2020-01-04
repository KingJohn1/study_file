// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _37_复数乘法

import (
	"fmt"
	"strconv"
	"strings"
)

//输入: "1+-1i", "1+-1i"
//输出: "0+-2i"
//解释: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。

var useGoApiComplex = false

func complexNumberMultiply(a string, b string) string {
	if useGoApiComplex {
		return complexNumberMultiplyByGoApi(a, b)
	}
	return complexNumberMultiplyNotByGoApi(a, b)
}

// 暂不实现，因为多次一举
func complexNumberMultiplyByGoApi(a string, b string) string {

	return ""
}

func complexNumberMultiplyNotByGoApi(a string, b string) string {
	var a1, a2, b1, b2 int
	fmt.Fscanf(strings.NewReader(a), "%d+%di", &a1, &a2)
	fmt.Fscanf(strings.NewReader(b), "%d+%di", &b1, &b2)
	fmt.Println(a1, a2, b1, b2)
	r := a1*b1 - a2*b2
	i := a1*b2 + a2*b1
	return strings.Join([]string{strconv.Itoa(r),"+",strconv.Itoa(i)+"i"},"")
}
