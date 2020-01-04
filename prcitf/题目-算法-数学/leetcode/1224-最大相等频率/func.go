// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _224_最大相等频率

//2 <= nums.length <= 10^5
//1 <= nums[i] <= 10^5

//执行用时 : 80 ms, 在所有 golang 提交中击败了 94.12% 的用户
//内存消耗 : 8.9 MB, 在所有 golang 提交中击败了 100.00% 的用户
func maxEqualFreq(nums []int) int {
	// 频率为index的数的个数
	numsPerHz := make([]int, 100001)
	// num value为index的频率
	hz := make([]int, 100001)

	// 出现的num的个数，即hz的非零项个数
	validHzCount := 0

	result := 0
	for i, v := range nums {
		if hz[v] == 0 {
			validHzCount++
		}
		hz[v]++

		// hz++后原频率的个数-1，先频率+1 ；原频率为0的为无效值，不会使用
		numsPerHz[hz[v]]++
		numsPerHz[hz[v]-1]--
		if validHz(hz, numsPerHz, i+1, validHzCount) {
			result = i + 1
		}
	}

	return result
}

func validHz(hz, numsPerHz []int, currentLenth, validHzCount int) bool {
	// 如果要满足以下任意一个条件则满足要求
	// 只有一种数
	// 或者 validecount 频率相等 并且为一
	// 或者 validHzCount-1 频率相等，另一个是比其频率加一；
	// 或者 validcount-1 频率相等 ，另一个是一

	if validHzCount == 1 {
		return true
	}

	// hz全部为1
	if numsPerHz[1] == validHzCount {
		return true
	}

	// validHzCount-1 频率相等，另一个是比其频率加一
	{
		if (currentLenth-1)%validHzCount == 0 {
			k := (currentLenth - 1) / validHzCount
			if numsPerHz[k] == validHzCount-1 && numsPerHz[k+1] == 1 {
				return true
			}
		}
	}

	// validcount-1 频率相等 ，另一个是一
	if (currentLenth-1)%(validHzCount-1) == 0 {
		k := (currentLenth - 1) / (validHzCount - 1)
		if numsPerHz[k] == validHzCount-1 && numsPerHz[1] == 1 {
			return true
		}
	}

	return false
}
