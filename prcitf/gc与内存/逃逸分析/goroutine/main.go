// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import "fmt"

func main(){
	var a int

	// b也逃逸，猜测goruntue的栈在heap中，因此gorountue全部逃逸
	go func(b int){
		var c int
		fmt.Println(b,c)
	}(a)
}