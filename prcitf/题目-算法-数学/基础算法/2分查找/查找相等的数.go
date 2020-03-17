// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package _分查找

import "fmt"

func bSearchEqual(data []int,target int) {
	index := binarySearchEqual(data, 0, len(data)-1, target)
	if index ==-1{
		fmt.Println("notfound")
	}else{
		fmt.Println("found ",target,"index:", index)
	}
}

func binarySearchEqual(data[]int,i ,j,target int)int{
	if i>j{
		return -1
	}

	midIndex:=(i+j)/2
	midData:=data[midIndex]

	if midData==target{
		return midIndex
	}else if midData<target{
		return binarySearchEqual(data,midIndex+1,j,target)
	}else {
		return binarySearchEqual(data,i,midIndex-1,target)
	}
}


