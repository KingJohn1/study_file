// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"fmt"
	"testing"
)

func Test_ConvString2Slice(t *testing.T) {
	slice := StringToStringByteSlice("124")
	//slice[1] = 1 // pannic
	fmt.Println(slice)
}

func ExampleStringByteSliceToString() {
	var slice = []byte{'1', '2', '3'}
	s := StringByteSliceToString(slice)
	fmt.Println(s)
	slice[0] = '2'
	fmt.Println(s)
	//output:123
	//223
}
