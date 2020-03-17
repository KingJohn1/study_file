// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _9_两数相除

import (
	"math"
)

//任何数字a都可以由1，2，4，8，16 ... 2的n次方  相加组成，a用二进制表示，第一位是2的0次方，第二位是2的1次方 ... 第n位是2的n次方；
// a的2进制对应为1则取出，则取出2的index的次放进行累加

// 为了避免越界使用64位内存,执行用时短，内存占用大
// 执行用时 : 0 ms, 在所有 golang 提交中击败了 100.00% 的用户
//内存消耗 : 2.6 MB, 在所有 golang 提交中击败了 8.00% 的用户
func divide_BigMember(dividend int, divisor int) int {
	//fmt.Println("input", dividend, divisor)
	var a, b = int64(dividend), int64(divisor)
	if a == 0 {
		return 0
	}
	isNegative := ((a < 0) && (b > 0)) || ((a > 0) && (b < 0))

	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	//fmt.Println("convert", a, b)

	if a < b {
		return 0
	}

	// degree[i]{x,y} x为2的i次方 y为x倍的b； i=0 x=1,y=b i=1,x=2,y=2b i=3,x=4,y=4b
	degree := make([][]int64, 0)
	degree = append(degree, []int64{1, b})
	for i := 1; ; i++ {
		count := degree[i-1][0] + degree[i-1][0]
		value := degree[i-1][1] + degree[i-1][1]
		degree = append(degree, []int64{count, value})
		if value > a {
			break
		}
	}
	//fmt.Println(degree)

	rest := a
	count := int64(0)
	//fmt.Print("choose:")
	for i := len(degree) - 2; i >= 0; i-- {
		rest -= degree[i][1]
		if rest < 0 {
			rest += degree[i][1]
			continue
		}
		count += degree[i][0]
		//	fmt.Print(" ", i, degree[i], rest, " ")
	}
	//fmt.Println("count", count)

	if count > math.MaxInt32 && !isNegative {
		count = math.MaxInt32
	}

	if count > math.MaxInt32+1 && isNegative {
		count = math.MaxInt32 + 1
	}

	if isNegative {
		count = -count
	}
	//fmt.Println(math.MaxInt32, math.MinInt32)
	return int(count)
}

func divide(a int, b int) int {
	if a == 0 {
		return 0
	}
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	if a == math.MinInt32 && b == 1 {
		return math.MinInt32
	}

	isNegative := ((a < 0) && (b > 0)) || ((a > 0) && (b < 0))

	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	if a < b {
		return 0
	}

	// degree[i]{x,y} x为2的i次方 y为x倍的b； i=0 x=1,y=b i=1,x=2,y=2b i=3,x=4,y=4b
	degree := make([][]int, 0)
	degree = append(degree, []int{1, b})
	for i := 1; ; i++ {
		count := degree[i-1][0] + degree[i-1][0]
		value := degree[i-1][1] + degree[i-1][1]
		degree = append(degree, []int{count, value})
		if value > a {
			break
		}
	}

	rest := a
	count := 0
	for i := len(degree) - 2; i >= 0; i-- {
		rest -= degree[i][1]
		if rest < 0 {
			rest += degree[i][1]
			continue
		}
		count += degree[i][0]
	}


	if isNegative {
		count = -count
	}
	return int(count)
}
