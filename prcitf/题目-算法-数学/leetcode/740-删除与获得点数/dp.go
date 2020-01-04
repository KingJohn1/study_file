// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _40_删除与获得点数

// 将数字的值作为索引存入pointsPerNum。 删除一个数值为a，则值为a-1 a+2都会被删除。
// 因此最后选择的数字一定不相邻，并且间隔最多为2。因为如果为3 ，则最后会剩余数，仍可以继续删除 获得更大的值

// 递推解法
func deleteAndEarnByDp(nums []int) int{
	var dp =make([]int,10001)
	var maxNum =0
	for _, v := range nums {
		dp[v] += v
		if v> maxNum {
			maxNum =v
		}
	}
	for i:=2;i<= maxNum;i++{
		dp[i]=max(dp[i-2]+dp[i],dp[i-1])
	}
	return dp[maxNum]
}

// 递归解法
func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 各个点数的总点数，index 为点数 value 为总点数;确保最后一位是0
	pointsPerNum := make([]int, 10002)
	dp:=make([]int,10002)
	for _, v := range nums {
		pointsPerNum[v] += v
	}
	//fmt.Println(pointsPerNum)

	var begin = 0
	var sum = 0
	for index := 1; index < len(pointsPerNum); index++ {
		if begin == 0 {
			if pointsPerNum[index] == 0 {
				continue
			}
			begin = index
			continue
		}
		//fmt.Println("begin not 0",index)
		if pointsPerNum[index] == 0 {
			sum += getMastEarn(pointsPerNum[begin:index],dp,begin)
			begin = 0
		}
	}
	return sum
}

func getMastEarn(points []int,dp []int,index int)(result int) {
	//fmt.Println("points",points)
	if len(points) == 0 {
		return 0
	}
	if len(points) == 1 {
		return points[0]
	}
	if dp[index]!=0{
		return dp[index]
	}

	result = points[0]
	if len(points) > 2 {
		result = result+ getMastEarn(points[2:],dp,index+2)
	}

	result= max(getMastEarn(points[1:],dp,index+1), result)
	dp[index]=result
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
