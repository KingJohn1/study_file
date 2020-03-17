// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _40_不同的子序列2

import "math"

// 此算法求解 由这些字母组成所有子串的种类，其中ab 与ba 是同一个子串

// 先统计各个字母的子串个数，不重复的字符串相当于排序后的字符串，因此后面的字母小于等于前面的字母
// a后可跟a,b，c，d，... “空”，
// b 后可跟b,c... ,"空"

// dp[a]=dp[b]+dp[c]...+1(空 即结束此子串)+ dp[a-1](a个数减一的排序数)
// dp[a-1]=dp[b]+dp[c]...+1(空)+ dp[a-2]
// 因此dp[a]=n(a的个数)*(dp[b]+dp[c]...+1(空))
// dp[b]=n(b的个数)*(dp[c]+dp[d]+...+1)

//因为结果可能很大，所以返回答案模 10^9 + 7.
//S 只包含小写字母。
//1 <= S.length <= 2000

var modValue = int(math.Pow(10, 9)) + 7

func distinctSubseqII_非本题目描述(S string) int {
	if S == "" {
		return 0
	}
	var countPerAlp [26]int

	// 利用余数加法乘法性质（题目-算法-数学/数学/取余mod/mod_test.go），保存取余结果
	var dp [26]int
	for _, v := range S {
		countPerAlp[v-'a']++
	}
	for i := 25; i >= 0; i-- {
		// dp[b]=n(b的个数)*(dp[c]+dp[d]+...+1)
		if countPerAlp[i] == 0 {
			continue
		}
		sumAfterCurrent := 1
		for j := i + 1; j < 26; j++ {
			sumAfterCurrent += dp[j]
		}
		dp[i] = countPerAlp[i] * sumAfterCurrent
		if dp[i] >= modValue {
			dp[i] = dp[i] % modValue
		}
	}
	sum := 0
	for i := 0; i < len(dp); i++ {
		sum += dp[i]
	}
	return sum % modValue
}
