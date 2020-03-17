// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_优先级(t *testing.T) {

	// 语义是以元音开头或以ok结束。而不是 开头结束中间夹着 元音或ok
	compile := regexp.MustCompile(`^[aeiou]|ok$`)

	fmt.Println(compile.FindString("ab"), compile.FindString("dok"), compile.FindString("ok"))
}
