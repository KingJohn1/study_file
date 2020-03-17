// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _40_不同的子序列2

// 难点在dp

// 定义：dp[i]表示s中0...i的所有子串,dp[i+1]=dp[i]+不重复部分

//两种场景
// 1：   如果 dp[i]内没有重复子串，且s[i+1]与s[0]...s[i]都不相等，则dp[i+1]=dp[i]*2+1(新生成的不重复的子串为 原子串追加s[i+1] 和s[i+1]单独作为子串)
// ，此时dp[i+1]也没有重复子串且是全集
// 2：   如果 dp[i]内没有重复子串，且s[i+1]与s[0]...s[i]有相等字符，与s[i+1]最近相等字符s[j](如abcbab,abcb是abcba最近以b结尾的)
// 则dp[i+1]=2*dp[i+1]-dp[j-1];原因为s[j]==s[i+1] 则s[j]追加到dp[j-1]的子串 等于s[i+1]追加到dp[j-1]的子串，且s[i+1]无法构成不重复子串，所以将此重复部分减掉
// 除了上述重复部分外，没有任何重复的子串。dp[j-1]与s[i+1]的重复部分已经删除，
// 如果存在删除则是 以s[j]...s[i]结尾的子串与s[i+1]拼接后与现有子串重复，
// 那说明s[j]...s[i]与s[0]...s[j-1]重复，因为重复意味着现有的子串与新生成子串全匹配 那么结尾的元素也一定相同，所以必然在 s[0]...s[j]
// 除去尾部即以s[j]...s[i]结尾的子串即现有元素一定在s[0]...s[j-1]中，则新增子串重复，则dp[i]内有重复元素，所以不存在删除则是 以s[j]...s[i]结尾的子串与s[i+1]拼接后与现有子串重复
// 所以 如果 dp[i]内没有重复子串 ...此描述条件下，dp[i+1]=2*dp[i+1]-dp[j-1]不存在重复元素且是全集（减掉了全集场景中所有重复元素）

// 证明递推dp正确：dp[1]第一个元素只有自身，一定没有重复元素，所以dp[2]按上述方法递推不会出现重复元素，所以dp[n]没有重复元素且是全集


//执行用时 : 0 ms, 在所有 golang 提交中击败了 100.00% 的用户
//内存消耗 : 2.3 MB, 在所有 golang 提交中击败了 100.00% 的用户
func distinctSubseqII(s string) int {
	var modValue = 1000000007
	if s == "" {
		return 0
	}
	// 第一位索引为1
	var lastIndex = make([]int, 26)
	var dp = make([]int, 2001)
	for i, v := range s {
		if dpIndex := lastIndex[v-'a']; dpIndex == 0 {
			dp[i+1] = (dp[i]*2 + 1)%modValue
		} else {
			dp[i+1] = (dp[i]*2 - dp[dpIndex-1])%modValue
		}

		dp[i+1]=dp[i+1]%modValue
		// 取mod后 dp[i+1]可能小于dp[i]
		if dp[i+1]<0{
			dp[i+1]+=modValue
		}

		lastIndex[v-'a'] = i + 1
		//fmt.Println(dp[i+1])
	}
	return dp[len(s)]%modValue
}
