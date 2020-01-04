// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _224_最大相等频率

import "testing"

func TestName(t *testing.T) {
	var data = []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{2, 2, 1, 1, 5, 3, 3, 5},
			expect: 7,
		},
		{
			nums:   []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5},
			expect: 13,
		},
		{
			nums:   []int{1, 1, 1, 2, 2, 2},
			expect: 5,
		},
		{
			nums:   []int{10, 2, 8, 9, 3, 8, 1, 5, 2, 3, 7, 6},
			expect: 8,
		},
		{
			nums:   nil,
			expect: 0,
		},
		{
			nums:   []int{10},
			expect: 1,
		},
		{
			nums:   []int{1, 1, 1, 1,},
			expect: 4,
		},
		{
			nums:   []int{2, 2, 2, 2, 4,6},
			expect: 5,
		},
		{
			nums:   []int{2, 2, 2, 3, 4, 5},
			expect: 4,
		},
		{
			nums:   []int{2, 2, 2, 2, 1, 2, 2, 3},
			expect: 7,
		},
	}
	for i, v := range data {

		if maxEqualFreq(v.nums) != v.expect {
			t.Error("index", i, "error", "nums:", v.nums, "expect:", v.expect, "get:", maxEqualFreq(v.nums))
		}
	}
}
