// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _67_字符串的排列

// 记录Source字符串的各个字母个数，然后循环s2 统计source的长度下 各字母个数，如果相等则存在排序
//时间复杂度：L1+(L2-L1)*26
//空间复杂度：O(1)。常数级的空间。

// 速度超过golang提交 80%
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var countS, countT [26]int
	for _, v := range s1 {
		countS[v-'a']++
	}

	for i := 0; i < len(s2); i++ {
		countT[s2[i]-'a']++

		if i >= len(s1) {
			countT[s2[i-len(s1)]-'a']--
		}
		if countT == countS {
			return true
		}
	}

	return false
}

// 速度超过golang提交 100% 即下面方法六
func checkInclusion1(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	count := 0
	var countS, countT [26]int

	for i := 0; i < len(s1); i++ {
		countS[s1[i]-'a']++
		countT[s2[i]-'a']++
	}

	for i := 0; i < len(countS); i++ {
		if countS[i] == countT[i] {
			count++
		}
	}

	for i := len(s1); i < len(s2); i++ {
		if count == 26 {
			return true
		}
		if s2[i-len(s1)] == s2[i] {
			continue
		}

		countT[s2[i]-'a']++
		if countT[s2[i]-'a'] == countS[s2[i]-'a'] {
			count++
		} else if countT[s2[i]-'a'] == countS[s2[i]-'a']+1 {
			count--
		}

		countT[s2[i-len(s1)]-'a']--
		if countT[s2[i-len(s1)]-'a'] == countS[s2[i-len(s1)]-'a'] {
			count++
		} else if countT[s2[i-len(s1)]-'a']+1 == countS[s2[i-len(s1)]-'a'] {
			count--
		}


	}

	return count == 26
}

// 网上看到一个优秀的方法
/*
方法六 优化的滑动窗口 [通过]:
算法

上一种方法可以优化，如果不是比较每个更新的 s2maps2map 的哈希表的所有元素，而是对应于 s2s2 考虑的每个窗口，我们会跟踪先前哈希表中已经匹配的元素数量当我们向右移动窗口时，只更新匹配元素的数量。

为此，我们维护一个 countcount 变量，该变量存储字符数（26个字母表中的数字），这些字符在 s1s1 中具有相同的出现频率，当前窗口在 s2s2 中。当我们滑动窗口时，如果扣除最后一个元素并添加新元素导致任何字符的新频率匹配，我们将 countcount 递增1.如果不是，我们保持 countcount 完整。但是，如果添加频率相同的字符（添加和删除之前）相同的字符，现在会导致频率不匹配，这会通过递减相同的 countcount 变量来考虑。如果在移动窗口后，countcount 的计算结果为26，则表示所有字符的频率完全匹配。所以，我们立即返回一个True。

Java
public class Solution {
    public boolean checkInclusion(String s1, String s2) {
        if (s1.length() > s2.length())
            return false;
        int[] s1map = new int[26];
        int[] s2map = new int[26];
        for (int i = 0; i < s1.length(); i++) {
            s1map[s1.charAt(i) - 'a']++;
            s2map[s2.charAt(i) - 'a']++;
        }
        int count = 0;
        for (int i = 0; i < 26; i++)
            if (s1map[i] == s2map[i])
                count++;
        for (int i = 0; i < s2.length() - s1.length(); i++) {
            int r = s2.charAt(i + s1.length()) - 'a', l = s2.charAt(i) - 'a';
            if (count == 26)
                return true;
            s2map[r]++;
            if (s2map[r] == s1map[r])
                count++;
            else if (s2map[r] == s1map[r] + 1)
                count--;
            s2map[l]--;
            if (s2map[l] == s1map[l])
                count++;
            else if (s2map[l] == s1map[l] - 1)
                count--;
        }
        return count == 26;
    }
}
复杂度分析

时间复杂度：L1+(L2-L1)
空间复杂度：O(1)。常数级的空间。

作者：LeetCode
链接：https://leetcode-cn.com/problems/permutation-in-string/solution/zi-fu-chuan-de-pai-lie-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
