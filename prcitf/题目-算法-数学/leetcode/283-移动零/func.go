// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _83_移动零

// 想象一片一片东西往一个方向压缩
func moveZeroes(nums []int)  {
	notZero:=0

	for i:=0;i<len(nums);i++{
		if nums[i]!=0{
			if notZero!=i{
				nums[notZero]=nums[i]
				nums[i]=0
			}
			notZero++
		}
	}
}