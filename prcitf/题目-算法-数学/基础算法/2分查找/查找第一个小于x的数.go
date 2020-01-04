// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package _分查找

import "fmt"

func binarySearchLess(data []int,target float64){
	index:=searchLess(data,0,len(data)-1,target)
	if index==-1{
		fmt.Println("notfound")
	}else {
		fmt.Println("found index:",index)
	}
}

func searchLess(data []int ,i,j int,target float64)int{

	if i>j{
		return -1
	}

	midIndex:=(i+j)/2
	midData:=data[midIndex]

	if float64(midData)>=target{
		fmt.Println(">= target")
	return 	searchLess(data,i,midIndex-1,target)
	}

	if float64(midData)<target{
		fmt.Println("<= target")
		if midIndex==j||float64(data[midIndex+1])>=target{
			fmt.Println("已是最优")
			return midIndex
		}
		return searchLess(data,midIndex+1,j,target)
	}

	return -1
}