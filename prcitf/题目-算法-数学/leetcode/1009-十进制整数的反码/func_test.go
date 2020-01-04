// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _009_十进制整数的反码

import (
	"fmt"
	"testing"
)

func Test_或操作(t *testing.T) {

	var a =01|010
	var b=1|10

	fmt.Println(a,b)

}

func Test_func(t *testing.T) {
	fmt.Println()
	fmt.Println(bitwiseComplement(5)==2)
	fmt.Println(bitwiseComplement(7)==0)
	fmt.Println(bitwiseComplement(10)==5)
}