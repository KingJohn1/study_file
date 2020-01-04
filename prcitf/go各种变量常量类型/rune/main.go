// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	s:="nihao,刘号"
	for i:=0;i<len(s);i++{
		fmt.Println(string( s[i]))
	}
	fmt.Println("rune 输出")
	for i:=0;i<len(s);{
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d \t %c \n",i,r,)
		i=i+size
	}
	utf8.EncodeRune()

}

