// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _35_分发糖果

import "fmt"

// 将分数化成折线图
// 如果没有相邻相等的，则折线图下顶点为1，上顶点为到左右顶点的距离最大值，上下顶点之间的点的值为从下顶点累加
// 比如 分数为 4 5 6 5 4 3 则值为 1 2 4 3 2 1
// 考虑相邻相等的情况；相邻相等的数的下一个数比当前数大则也认为是下顶点，如果下一个数比当前数小则认为是上顶点
// 其实相邻相等即意味着两侧完全没有关系，可以分别求解再求和；如 分数为1 2 2 3的值即分数为1 2 的值组合分数为 2 3的值

// 附加：求出没人得到的糖果
func candy(ratings []int) int {
	fmt.Println(ratings)
	var candys = make([]int, len(ratings))

	sum := 0
	lowPos := -1
	highPos := -1
	cal := func(i int) {
		if lowPos == -1 && highPos == -1 {
			candys[i] = 1
			sum += 1
		} else if lowPos != -1 && highPos == -1 {

			{
				var index int
				for j := lowPos; j <= i; j++ {
					index++
					candys[j] = index
				}
			}

			sum += (1 + (i - lowPos + 1)) * (i - lowPos + 1) / 2 // 求高公式
		} else if lowPos == -1 && highPos != -1 {
			{
				var index int
				for j := i; j >= highPos; j-- {
					index++
					candys[j] = index
				}
			}
			sum += (i - highPos + 1 + 1) * (i - highPos + 1) / 2
		} else {
			left := highPos - lowPos
			right := i - highPos
			if left > right {
				{
					var index int
					for j := lowPos; j <= highPos; j++ {
						index++
						candys[j] = index
					}
					index = 0
					for j := i; j < highPos; j-- {
						index++
						candys[j] = index
					}
				}

				sum += (1 + left + 1) * (left + 1) / 2
				sum += (1 + right) * right / 2
			} else {
				{
					var index int
					for j := lowPos; j < highPos; j++ {
						index++
						candys[j] = index
					}
					index = 0
					for j := i; j >= highPos; j-- {
						index++
						candys[j] = index
					}
				}
				sum += (1 + left) * left / 2
				sum += (1 + right + 1) * (right + 1) / 2
			}
		}
		lowPos = -1
		highPos = -1
	}
	for i := 0; i < len(ratings); i++ {
		if i+1 == len(ratings) {
			cal(i)
			break
		}
		if ratings[i] == ratings[i+1] {
			cal(i)
			continue
		}
		if ratings[i] < ratings[i+1] {
			if highPos != -1 {
				cal(i)
				sum--
				lowPos = i
			}
			if lowPos == -1 {
				lowPos = i
			}
		}
		if ratings[i] > ratings[i+1] {
			if highPos == -1 {
				highPos = i
			}
		}
	}
	fmt.Println(candys)
	return sum
}

// 执行用时 : 20 ms , 在所有 golang 提交中击败了 95.60% 的用户
//内存消耗 : 6 MB , 在所有 golang 提交中击败了 100.00% 的用户
// 基础；求糖果总数
func candy1(ratings []int) int {

	sum := 0
	lowPos := -1
	highPos := -1
	cal := func(i int) {
		if lowPos == -1 && highPos == -1 {
			sum += 1
		} else if lowPos != -1 && highPos == -1 {

			{
				var index int
				for j := lowPos; j <= i; j++ {
					index++
				}
			}

			sum += (1 + (i - lowPos + 1)) * (i - lowPos + 1) / 2 // 求高公式
		} else if lowPos == -1 && highPos != -1 {
			{
				var index int
				for j := i; j >= highPos; j-- {
					index++
				}
			}
			sum += (i - highPos + 1 + 1) * (i - highPos + 1) / 2
		} else {
			left := highPos - lowPos
			right := i - highPos
			if left > right {
				sum += (1 + left + 1) * (left + 1) / 2
				sum += (1 + right) * right / 2
			} else {
				sum += (1 + left) * left / 2
				sum += (1 + right + 1) * (right + 1) / 2
			}
		}
		lowPos = -1
		highPos = -1
	}
	for i := 0; i < len(ratings); i++ {
		if i+1 == len(ratings) {
			cal(i)
			break
		}
		if ratings[i] == ratings[i+1] {
			cal(i)
			continue
		}
		if ratings[i] < ratings[i+1] {
			if highPos != -1 {
				cal(i)
				sum--
				lowPos = i
			}
			if lowPos == -1 {
				lowPos = i
			}
		}
		if ratings[i] > ratings[i+1] {
			if highPos == -1 {
				highPos = i
			}
		}
	}

	return sum
}
