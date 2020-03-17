// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _75_爱吃香蕉的珂珂

import "math"

func minEatingSpeed(piles []int, H int) int {

	if len(piles) == 0 {
		return 0
	}

	initL, initR := getInitK(piles, H)

	return binarySearchLessThan(piles, initL, initR, H)
}

func binarySearchLessThan(piles []int, l, r int, h int) int {
	if l > r {
		return -1
	}

	midIndex := (l + r) / 2
	midValue := getH(piles, midIndex)
	if midValue > h {
		return binarySearchLessThan(piles, midIndex+1, r, h)
	}
	if midIndex == 1 {
		return midIndex
	}

	if leftValue := getH(piles, midIndex-1); leftValue > h {
		return midIndex
	}
	return binarySearchLessThan(piles, l, midIndex-1, h)
}

func getH(piles []int, k int) int {
	sum := 0
	for _, v := range piles {
		//sum += (v-1)/k + 1
		sum+=int(math.Ceil(float64(v)/float64(k))) //效率更高 在所有 golang 提交中击败了 97.20% 的用户，不用ceil函数排名倒数
	}
	return sum
}

func getInitK(piles []int, h int) (int, int) {
	max := 0
	sum := 0
	for _, v := range piles {
		if v > max {
			sum += v
			max = v
		}
	}

	min := 1
	if sum/h > 1 {
		min = sum / h
	}
	return min, max
}

func minEatingSpeed1(piles []int, H int) int {

	if len(piles) == 0 {
		return 0
	}

	initL, initR := getInitK(piles, H)

	l, r := initL, initR

	for {
		if l > r {
			return -1
		}
		midIndex := (l + r) / 2
		midValue := getH(piles, midIndex)
		if midValue > H {
			l = midIndex + 1
			continue
		}

		if midIndex == 1 {
			return midIndex
		}
		if leftValue := getH(piles, midIndex-1); leftValue > H {
			return midIndex
		}
		r = midIndex - 1
	}
}
