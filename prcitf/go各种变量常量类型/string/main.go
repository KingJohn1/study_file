// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main


import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("liuhao",`liu\s`))
	fmt.Println(strings.Join([]string{"liu","hao"}," "))
	fmt.Println(strings.Index("liuhao","hao"))
	fmt.Println(strings.Repeat("liuhao",1))
	fmt.Println(strings.Replace("liuhao","hao","haohao",-1))
	fmt.Println(strings.Trim("liuhaoliu","liu"))
	fmt.Println(strings.Fields("liu hao ni hao ya"))

var str []byte

str=strconv.AppendBool(str,true)
str=strconv.AppendInt(str,12,10)
str=strconv.AppendFloat(str,11.11,'g',11111111,64)
os.Stdout.Write(str)
fmt.Println()

	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

}
