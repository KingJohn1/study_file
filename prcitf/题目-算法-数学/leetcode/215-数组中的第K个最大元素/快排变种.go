// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _15_数组中的第K个最大元素

import "sort"

func findKthLargestByQuickSort(nums []int, k int) int {
	return quickSelectKthSmallLest(nums, len(nums)-k, 0, len(nums)-1)
}

func quickSelectKthSmallLest(nums []int, k, leftPos, rightPos int) int {

	if leftPos > rightPos || k < leftPos || k > rightPos || leftPos < 0 || rightPos >= len(nums) {
		return -1
	}

	if leftPos == rightPos && leftPos == k {
		return nums[k]
	}

	// 选最左边为基准，也可以选其他数为基准 选完后与最左边交换。这里选最左边为比较基准

	base := nums[leftPos]
	basePos := leftPos
	leftCursor, rightCursor := leftPos, rightPos
	for {
		for ; rightCursor > leftCursor; rightCursor-- {
			if nums[rightCursor] < base {
				break
			}
		}
		for ; leftCursor < rightCursor; leftCursor++ {
			if nums[leftCursor] > base {
				break
			}
		}
		if leftCursor == rightCursor {
			nums[leftCursor], nums[basePos] = nums[basePos], nums[leftCursor]
			break
		}
		nums[leftCursor], nums[rightCursor] = nums[rightCursor], nums[leftCursor]
	}
	midCurson := leftCursor
	//fmt.Println(midCurson,k,leftPos,rightPos,nums)
	if midCurson == k {
		return nums[k]
	}
	if midCurson > k {
		return quickSelectKthSmallLest(nums, k, leftPos, midCurson-1)
	}
	return quickSelectKthSmallLest(nums, k, midCurson+1, rightPos)
}

func findKthLargestBySort(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
