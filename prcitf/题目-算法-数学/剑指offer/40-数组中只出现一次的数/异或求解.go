// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _0_数组中只出现一次的数

import (
	"fmt"
	_ "runtime/pprof"
)

// 这种方法只能求解出现次数为2的，不具有普遍性
func getAppearOnceNumsByXor(nums []int) []int {

	xorSum := xorSums(nums)
	if xorSum == 0 {
		fmt.Println("xorsum is 0")
		return nil
	}

	index := 0
	for {
		if xorSum&1 == 1 {
			break
		}
		index++
		xorSum = xorSum >> 1
	}

	// 筛选此位是0和1的数 组成两个数组

	nums1 := make([]int, 0, len(nums))
	nums2 := make([]int, 0, len(nums))

	for _, v := range nums {
		if (v>>uint32(index))&1 == 1 {
			nums1 = append(nums1, v)
		} else {
			nums2 = append(nums2, v)
		}
	}
	result := make([]int, 0, 2)

	return append(result, xorSums(nums1), xorSums(nums2))
}

func xorSums(nums []int) int {
	xorSum := 0
	for _, v := range nums {
		xorSum ^= v
	}
	return xorSum
}
