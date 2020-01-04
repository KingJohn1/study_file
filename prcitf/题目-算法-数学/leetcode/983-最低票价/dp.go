// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _83_最低票价

import (
	"math"
)

var dp []int
var cheap7th, cheap30th bool

func mincostTickets(days []int, costs []int) int {
	cheap7th,cheap30th=true,true

	// 边界条件
	if costs[0] == 0 || costs[1] == 0 || costs[2] == 0 {
		return 0
	}

	dp = make([]int, len(days)+1)

	if costs[1] >= costs[0]*7 {
		cheap7th = false
	}

	if costs[2] >= costs[0]*30 || costs[2] >= costs[1]*5 {
		cheap30th = false
	}

	return getMinCostTickets(days, costs)

}

func getMinCostTickets(days []int, costs []int) (result int) {
	defer func() {
		dp[len(days)] = result
	}()
	if len(days) == 0 {
		return 0

	}
	if dp[len(days)] != 0 {
		return dp[len(days)]
	}
	costs1th := costs[0] + getMinCostTickets(days[1:], costs)
	costs7th := math.MaxInt32
	costs30th := math.MaxInt32
	if cheap7th {
		costs7th = costs[1]
		for i := 0; i < len(days); i++ {
			if days[i]-days[0]+1 >7 {
				costs7th += getMinCostTickets(days[i:], costs)
				break
			}
		}
	}

	//输入：days = [1,2,3,4,5,6,7,8,9,10,30,31], costs = [2,7,15]
	//输出：17
	if cheap30th {
		costs30th = costs[2]
		for i := 0; i < len(days); i++ {
			if days[i]-days[0]+1 > 30 {
				costs30th += getMinCostTickets(days[i:], costs)
				break
			}
		}
	}

	return min(costs1th, costs7th, costs30th)

}

func min(i int, i2 int, i3 int)(r int) {

	if i < i2 {
		if i < i3 {
			return i
		}
		return i3
	}

	if i2 < i3 {
		return i2
	}
	return i3
}
