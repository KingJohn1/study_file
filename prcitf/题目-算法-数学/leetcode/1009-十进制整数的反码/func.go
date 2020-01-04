// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _009_十进制整数的反码

// 直接求解
func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}

	var result int
	for i := 0; N != 0; i++ {
		if N&1 == 0 {
			result += 1 << uint32(i)
		}
		N >>= 1
	}

	return result
}

// 利用性质 n与其反码等与 n的位数的最大数  比如1101 与0010 的合为1111 也即10000 -1

// sprintf %b 得到整数然后取反，然后利用strconv.ParseInt(),复杂性能差 得不偿失

// 预处理后可以运用取反：网上方法
// 用对数函数 log2 获得 N 最高位 1 的位置。
//然后对1左移后减1，获得比N长度小1的掩码。
//将 N 取反，再和掩码与运算。
//00000101 N
//00000011 mask
//00000010 (~N) & mask
//
//作者：hareyukai
//链接：https://leetcode-cn.com/problems/complement-of-base-10-integer/solution/-bi-mian-xun-huan-by-hareyukai/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
