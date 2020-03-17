// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _00_最长上升子序列_长度_

// 求最长上升子序列用dp或递归，求最长上升子序列长度是个经典问题，之前读过的挑战程序竞赛有对这个问题的解法。leetcode题解也有描述，二分加贪心

func lengthOfLIS(nums []int) int {
	var minTailLis = make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		minTailLis = fillNum(minTailLis, nums[i])
	}

	return len(minTailLis)
}

// append 不会改变原有数组，所以使用append 必须返回
func fillNum(minTailLis []int, num int) []int {
	if len(minTailLis) == 0 {
		minTailLis = append(minTailLis, num)
	}
	if num > minTailLis[len(minTailLis)-1] {
		minTailLis = append(minTailLis, num)
		return minTailLis
	}
	index := binarySearchFirstMore(minTailLis, 0, len(minTailLis)-1, num)
	if index != -1 {
		minTailLis[index] = num
	}
	//fmt.Println(minTailLis)
	return minTailLis
}

func binarySearchFirstMore(minTailLis []int, l, r, num int) int {

	if l > r {
		return -1
	}

	midIndex := (l + r) / 2
	midValue := minTailLis[midIndex]
	if midValue <= num {
		return binarySearchFirstMore(minTailLis, midIndex+1, r, num)
	}
	if midIndex == 0 || minTailLis[midIndex-1] < num {
		return midIndex
	}
	if minTailLis[midIndex-1] == num {
		return -1
	}
	return binarySearchFirstMore(minTailLis, l, midIndex-1, num)
}
