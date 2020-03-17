// Copyright (c) liuhao. 2012-2050. All rights reserved.
package pp

import (
	"testing"
)

var nums = []int{1, 5, 1, 1, 6, 4}

var data = []struct {
	nums []int
}{
	{[]int{1, 5, 1, 1, 6, 4},},
	{[]int{1, 3, 2, 2, 3, 1},},
	{[]int{4, 5, 5, 6, },},
}

func Test_排序解决摆动数组(t *testing.T) {

	check := func(index int, nums []int) {
		for i := range nums {
			if i+1 != len(nums) {
				if i%2 == 0 {
					if  nums[i] >= nums[i+1]{
						t.Errorf("index :%v index :%v  bad " ,index,i)
					}
				} else {
					if  nums[i] <= nums[i+1]{
						t.Errorf("index :%v index :%v  bad " ,index,i)
					}
				}
			}
		}
	}
	for i, v := range data {
		wiggleSortAppend(v.nums)
		check(i,v.nums)
	}
}
