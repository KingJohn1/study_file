// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package _分查找

import (
	"testing"
)

func Test_BinanrySearchEqual(t *testing.T) {

	var data =[]int{1,2,3,4,5,6,7,8}
	bSearchEqual(data,8)
}

func Test_binarySearchLess(t *testing.T) {
	var data =[]int{0,1,2,3,4,5,6,7,8}
	binarySearchLess( data,1.2)
}