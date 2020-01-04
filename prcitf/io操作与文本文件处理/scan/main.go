// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	scanner:= bufio.NewScanner(os.Stdin)
	//scanner.Split()

	for scanner.Scan(){

		fmt.Println(scanner.Text())
	}
	fmt.Println("end")

	//var a ,b int
	//n, err := fmt.Scanf("%d %d", &a, &b)
	//fmt.Println(n,err)
	//fmt.Println(a,b)
	//
	//n, err = fmt.Scanf("%d %d", &a, &b)
	//fmt.Println(n,err)
	//fmt.Println(a,b)
	//
	//n, err = fmt.Scanf("%d %d", &a, &b)
	//fmt.Println(n,err)
	//fmt.Println(a,b)

}

