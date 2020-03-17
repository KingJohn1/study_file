// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a int64 = 1
	var b int16 = int16(a)
	fmt.Println(&a, &b)

	var c int32 = 1
	var d = int8(c)
	fmt.Println(&c, &d)
	//output:
	// 0xc000054080 0xc000054088
	//0xc00005408c 0xc0000540b0
	// 说明 强制转换会开辟新的空间，并对齐

	var aa = "1234"
	var bb = []byte(aa)
	fmt.Println(aa, bb)
	bb[1] = 'a'
	fmt.Println(aa, bb)
	// bb 修改aa未修改，说明bb重新拷贝内存

	cc := StringToStringByteSlice(aa)
	fmt.Println(len(cc), cap(cc))
	// cc[1]=1//会panic
	fmt.Printf("%s %s %p \n", aa, cc, &cc)

	cc = StringToStringByteSlice(string(bb))
	fmt.Println(len(cc), cap(cc))
	cc = cc[1:]
	cc[1] = 1
	fmt.Printf("%s %s \n", aa, cc)

	s := StringByteSliceToString(bb)
	fmt.Println(s)
	bb[1] = 'b'
	fmt.Println(s)
}

func StringToStringByteSlice(str string) []byte {
	strPtr := (*[2]uintptr)(unsafe.Pointer(&str))
	res := [3]uintptr{strPtr[0], strPtr[1], strPtr[1]}
	return *(*[]byte)(unsafe.Pointer(&res))
}

func StringByteSliceToString(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}
