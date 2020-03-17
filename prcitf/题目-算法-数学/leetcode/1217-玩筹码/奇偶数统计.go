// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _217_玩筹码

// chips的value是筹码的位置
func minCostToMoveChips1(chips []int) int {
	var evenSum,oddSum int
	for _,v:=range chips{
		if v%2==0{
			evenSum++
		}else{
			oddSum++
		}
	}
	if evenSum>oddSum{
		return oddSum
	}
	return evenSum
}