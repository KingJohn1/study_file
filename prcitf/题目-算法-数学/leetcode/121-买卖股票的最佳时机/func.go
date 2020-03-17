// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _21_买卖股票的最佳时机

func maxProfit(prices []int) int {
	var ringhtMax int
	var max int
	for i:=len(prices)-1;i>=0;i--{
		if prices[i]>ringhtMax{
			ringhtMax=prices[i]
		}
		if ringhtMax-prices[i]>max{
			max=ringhtMax-prices[i]
		}
	}
	return max
}