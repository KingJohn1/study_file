// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package _144递减元素使数组呈锯齿状

func min(a,b int)int{
	if a<b {
		return a
	}
	return b
}

func movesToMakeZigzag(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}

	// 尝试奇数大于相邻节点
	oddSum := 0
	// 偶数点应该小于周围点
	for i := 0; i < len(nums); i += 2 {
		if i-1<0{
			if nums[i+1]<=nums[i]{
				oddSum=oddSum+nums[i]-nums[i+1]+1
			}
		}
		if i+1>=len(nums){
			if nums[i-1]<=nums[i]{
				oddSum=oddSum+nums[i]-nums[i-1]+1
			}
		}
		if i-1>=0&&i+1<len(nums){
			if nums[i]>= min(nums[i-1],nums[i+1]){
				oddSum=oddSum+nums[i]-min(nums[i-1],nums[i+1])+1
			}
		}
	}

	// 尝试偶数大于相邻节点,奇数点应该小于周围点
	evenSum:=0
	for i := 1; i < len(nums); i += 2 {
		if i-1<0{
			if nums[i+1]<=nums[i]{
				evenSum=evenSum+nums[i]-nums[i+1]+1
			}
		}
		if i+1>=len(nums){
			if nums[i-1]<=nums[i]{
				evenSum=evenSum+nums[i]-nums[i-1]+1
			}
		}
		if i-1>=0&&i+1<len(nums){
			if nums[i]>= min(nums[i-1],nums[i+1]){
				evenSum=evenSum+nums[i]-min(nums[i-1],nums[i+1])+1
			}
		}
	}

	return min(oddSum,evenSum)
}
