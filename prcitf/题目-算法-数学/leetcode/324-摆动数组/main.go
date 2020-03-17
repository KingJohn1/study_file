// Copyright (c) liuhao. 2012-2050. All rights reserved.
package pp

import (
	"fmt"
	"sort"
)

// 此算法无法解决 输入: nums = [1, 3, 2, 2, 3, 1] 将得到[1 3 2 2 1 3]
//输出: 一个可能的答案是 [2, 3, 1, 3, 1, 2]
func wiggleSort(nums []int) {
	defer func() {
		fmt.Println(nums)
	}()
	sort.Ints(nums)

	if len(nums) < 3 {
		return
	}

	if len(nums) == 3 {
		nums[1], nums[2] = nums[2], nums[1]
		return
	}

	rightP := len(nums) / 2
	if rightP%2 == 1 {
		rightP++
	}

	for leftP := 1; leftP < len(nums)/2; leftP = leftP + 2 {

		nums[leftP], nums[rightP] = nums[rightP], nums[leftP]

		rightP = rightP + 2
	}

	if (len(nums)-1)%2 == 1 && nums[len(nums)-1] < nums[len(nums)-2] {
		nums[len(nums)-2], nums[len(nums)-1] = nums[len(nums)-1], nums[len(nums)-2]
	}
	if (len(nums)-1)%2 == 0 && nums[len(nums)-1] > nums[len(nums)-2] {
		nums[len(nums)-2], nums[len(nums)-1] = nums[len(nums)-1], nums[len(nums)-2]
	}
}

func wiggleSortAppend(nums []int) {
	defer func() {
		fmt.Println(nums)
	}()
	sort.Ints(nums)

	if len(nums) < 2 {
		return
	}
	copyData := make([]int, len(nums), len(nums))

	copy(copyData, nums)
	even := (len(nums)+1)/2-1
	odd := len(nums)-1


	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			nums[i] = copyData[even]
			even--
		} else {
			nums[i] = copyData[odd]
			odd--
		}
	}
}
