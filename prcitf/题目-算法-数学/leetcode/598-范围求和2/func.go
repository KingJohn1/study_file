// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _98_范围求和2

// 由于是求最大元素的个数，而非最大元素，因此不需要用暴力法。只需求 重叠矩形的面积即可。

func maxCount(m int, n int, ops [][]int) int {
	min1, min2 := m, n
	for i := 0; i < len(ops); i++ {
		if ops[i][0]<min1{
			min1=ops[i][0]
		}
		if ops[i][1]<min2{
			min2=ops[i][1]
		}
	}
	return min1*min2
}
