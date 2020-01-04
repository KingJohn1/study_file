// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _77_汉明距离总和

// 暴力法 n的平方级别

// 对每个位 分别统计1的数量和0的数量的乘积 n级别时间复杂度

// 易读,不改变原有数据
func totalHammingDistance(nums []int) int {

	var sumBit1, sumBit0 = make([]int, 32), make([]int, 32)

	for i := 0; i < len(nums); i++ {
		tmp := 1
		for j := 0; j < 32; j++ {
			if tmp&nums[i] == 0 {
				sumBit0[j]++
			} else {
				sumBit1[j]++
			}
			tmp = tmp << 1
		}
	}

	var sum int
	for i := 0; i < len(sumBit0); i++ {
		sum += sumBit0[i] * sumBit1[i]
	}
	return sum
}

// 性能，不改变原有数据
func totalHammingDistance1(nums []int) int {

	var sum int
	tmp := 1
	for i := 0; i < 32; i++ {
		total := 0
		for j := 0; j < len(nums); j++ {
			if tmp&nums[j] == 0 {
				total++
			}
		}
		tmp <<= 1
		sum += total * (len(nums) - total)
	}

	return sum
}
