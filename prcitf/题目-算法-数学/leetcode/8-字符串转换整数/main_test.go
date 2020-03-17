// Copyright (c) liuhao. 2012-2050. All rights reserved.
package tt

import (
	"fmt"
	"strconv"
	"testing"
	"testing/quick"
)

var data = []struct {
	input  string
	output int
}{
	{
		input:  "1224",
		output: 1224,
	},
	{
		input:  "4193 with words",
		output: 4193,
	},
	{
		input:  "4193with words",
		output: 4193,
	},
	{
		input:  "with words 121",
		output: 0,
	},
	{
		input:  "-91283472332",
		output: -1 << 31,
	},
	{
		input:  "91283472332",
		output: 1<<31 - 1,
	},
}

func Test_利用FScanf处理(t *testing.T) {

	for i := range data {
		if myAtoi_scan函数处理(data[i].input) != data[i].output {
			t.Errorf("index i want %v ,got %v", data[i].output, myAtoi_scan函数处理(data[i].input))
		}
	}
}

func Test_循环读取处理(t *testing.T) {

	for i := range data {
		if myAtoi(data[i].input) != data[i].output {
			t.Errorf("index i want %v ,got %v", data[i].output, myAtoi(data[i].input))
		}
	}
}

func reportError(property string, err error, t *testing.T) {
	if err != nil {
		t.Errorf("%s: %s", property, err)
	}
}

func Test_Atoi(t *testing.T) {

	reportError("Atoi",quick.CheckEqual(myAtoi_scan函数处理,myAtoi,nil),t)
	reportError("Atoi",quick.Check(func(b int32)bool{
		s:=strconv.FormatInt(int64(b),10)
		//println(myAtoi_scan函数处理(s))
		return  myAtoi_scan函数处理(s)==myAtoi(s)
	},nil),t)
}

func Test_两种方式是否输出相等(t *testing.T) {
	// Test_Atoi显示-1097181211 两个函数不相等，后检查代码发现，没有atoi没有带符号位
	fmt.Println(myAtoi_scan函数处理("-1097181211"),myAtoi("-1097181211"))
}

func Test_leetcode输出失败用例(t *testing.T) {
	fmt.Println(myAtoi("  -0012a42"))
	fmt.Println(myAtoi_scan函数处理("20000000000000000000"))
}