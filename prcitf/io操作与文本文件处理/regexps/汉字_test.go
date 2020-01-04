// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main_test

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_汉字(t *testing.T) {
	m, _ := regexp.MatchString("\\p{Han}+$", "liuhao刘号大声道")
	fmt.Println(m)

	m1, _ := regexp.MatchString("[\u4e00-\u9fa5]", "liuhao刘号大声道")
	fmt.Println(m1)


	//(?:  pattern)是非捕获型括号  匹配pattern，但不捕获匹配结果。
	//(pattern )是捕获型括号。  匹配pattern，匹配pattern并捕获结果,自动获取组号
	//(?<name> pattern )  匹配pattern，  匹配pattern并捕获结果，设置name为组名
	//    使用小括号指定一个子表达式后，匹配这个子表达式的文本(也就是此分组捕获的内容)可以在表达式或其它程序中作进一步的处理。默认情况下，每个捕获组会自动拥有一个组号，规则是：从左向右，以分组的左括号为标志，第一个出现的分组的组号为1，第二个为2，以此类推。
	//    如果正则表达式中同时存在普通捕获组和命名捕获组，那么捕获组的编号就要特别注意，编号的规则是先对普通捕获组进行编号，再对命名捕获组进行编号。
	//     为了避免括号太多使编号混乱，也为了避免无用的捕获提高效率，在不需要捕获只需要指定分组的地方就可以使用非捕获型括号。问题里的非捕获型括号就是为此使用的

	//运用正则捕获组
	compile := regexp.MustCompile(`(?P<liu>{Han}+$)`)
	fmt.Println(compile.FindStringSubmatch("liuhao刘号大声道"))
	fmt.Println(compile.SubexpNames())
	fmt.Println("")

	r := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)
	fmt.Printf("%#v\n", r.FindStringSubmatch(`2015 2015-05-27`))
	fmt.Printf("%#v\n", r.SubexpNames())

	// (?:  pattern)是非捕获型括号  匹配pattern，但不捕获匹配结果。todo 不知如何使用
	rr := regexp.MustCompile(`(?:<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)
	fmt.Printf("%#v\n", rr.FindStringSubmatch(`2015 2015-05-27`))
	fmt.Printf("%#v\n", rr.SubexpNames())
}
