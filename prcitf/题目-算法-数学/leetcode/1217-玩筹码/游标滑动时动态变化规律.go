// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _217_玩筹码

import "math"

//1 <= chips.length <= 100
//1 <= chips[i] <= 10^9


// 下面算法理解错了题意，下面算法是求 所有的点都移到某个中心点左右，
// 中心点左边的点都移动中心点左侧第一个点，中心点右边的点都移动到中心点右边的第一个点；
// 如果距离这个中心点的长度是偶数，则不花费点数，如果距离是奇数则每移动2个点花费当前的数值
// 奇数位 偶数位分别遍历
// 对奇数位（或偶数位）左右两边的变量分别求和；
// 中心节点右移动两位： 更新左边数统计和 ，左边的数移动到当前中心节点操作数为原操作数加当前左边数字统计和；右边操作数为原操作数和减去右边统一和 然后更新右侧数字统计和；
// 中心节点从一侧一次移动到另一侧，求最小值
func minCostToMoveChips(chips []int) int {
	min := math.MaxInt64
	if len(chips) <= 1 {
		return 0
	}

	leftOddSum := 0
	rightOddSum := 0
	leftEvenSum := 0
	rightEvenSum := 0
	leftOps := 0
	rightOps := 0
	rightOpsEven := 0
	for i := 0; i < len(chips); i++ {
		if i%2 == 0 {
			rightEvenSum += chips[i]
			rightOpsEven += rightEvenSum
		} else {
			rightOddSum += chips[i]
			rightOps += rightOddSum
		}
	}

	for even := 0; even < len(chips); even += 2 {
		leftNum := 0
		if even > 0 {
			leftNum = chips[even-1]
		}
		leftOddSum += leftNum
		leftOps += leftOddSum

		rightOps -= rightOddSum
		rightOddSum -= leftNum

		if sum := leftOps + rightOps; sum < min {
			min = sum
		}
	}

	leftOps = 0
	rightOps = rightOpsEven
	for odd := 1; odd < len(chips); odd += 2 {
		leftNum := chips[odd-1]
		leftOddSum += leftNum
		leftOps += leftEvenSum

		rightOps -= rightEvenSum
		rightEvenSum -= leftNum

		if sum := leftOps + rightOps; sum < min {
			min = sum
		}
	}
	return min
}
