// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main1(){
	fmt.Println(os.Args)
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("数字")
	} else {
		fmt.Println("不是数字")
	}
	// !匹配汉字
	if m, _ := regexp.MatchString("^\\p{Han}+$", "liuhao刘号大声道"); !m {
	   return
	}
}

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func main() {

	src1 := []byte(`
        call hello alice
        hello bob
        call hello eve
    `)
	pat := regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res2 := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src1, -1) {
		res2 = pat.Expand(res2, []byte("$cmd('$arg')\n"), src1, s)
	}
	fmt.Println(string(res2))

	re1 := regexp.MustCompile("ab?")
	fmt.Println(re1.FindStringIndex("t1blettab1"))
	fmt.Println(re1.FindStringIndex("foo") == nil)
	resp, err := http.Get("http://127.0.0.1:8080/")
	if err != nil {
		fmt.Println("http get error.")
	}
	fmt.Println("resp ",resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}