// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _40_不同的子序列2

// 复杂度较高
func distinctSubseqII_map过滤(s string) int {
	if s == "" {
		return 0
	}
	// substr[i]表示从s[i:]的所有子串，则substr[i-1] 为 （s[i] 和 s[i]拼接substr[i]每个子串）的不重复部分
	// 从右往左遍历将所有的子串存入
	var subStr = make(map[string]struct{})

	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		newSub := make([]string, 0)
		currenAlp := s[i : i+1]
		_, ok := subStr[currenAlp]
		if !ok {
			count++
			newSub = append(newSub, currenAlp)
		}

		for k := range subStr {
			_, ok := subStr[currenAlp+k]
			if !ok {
				count++
				newSub = append(newSub, currenAlp+k)
			}
		}
		for _, v := range newSub {
			subStr[v] = struct{}{}
		}
	}
	return count % modValue
}
