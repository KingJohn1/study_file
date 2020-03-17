// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _0_数组中只出现一次的数

// 代码简洁 适用场景广。但是map的性能与异或比差距较大。详见 benchmark
func getAppearOnceNumsByStatistics(nums []int) []int {
	counts := make(map[int]bool, len(nums))

	for _, v := range nums {
		value, ok := counts[v]
		if !ok {
			counts[v] = true
		} else if value{
			counts[v] = false
		}
	}
	result := make([]int, 0, 2)
	for k, v := range counts {
		if v {
			result = append(result, k)
		}
	}

	return result
}
