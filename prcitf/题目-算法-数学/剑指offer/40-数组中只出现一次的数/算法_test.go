// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _0_数组中只出现一次的数

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestXor(t *testing.T) {
	var a, b int
	a = -1
	b = -1
	fmt.Println(a ^ b) //output 0;相同的负数异或也等于0
	c := 1
	fmt.Println(a ^ c) // output -2;说明 负数以反码进行异或
}

func Test_位操作(t *testing.T) {
	var a = -2

	fmt.Println(a >> 1) //output -1
}

func Test_异或求解(t *testing.T) {
	var nums = []int{
		2, 4, 3, 6, 3, 2, 5, 5, -2, 19, -2, 19,
	}
	fmt.Println(getAppearOnceNumsByXor(nums))
	fmt.Println(getAppearOnceNumsByXor(numsRand))
}

func Test_统计求解(t *testing.T) {
	var nums = []int{
		2, 4, 3, 6, 3, 2, 5, 5, -2, 19, -2, 19,
	}
	fmt.Println(getAppearOnceNumsByStatistics(nums))
	fmt.Println(getAppearOnceNumsByStatistics(numsRand))
}

var nums1 = []int{2, 4, 3, 6, 3, 2, 5, 5, -2, 19, -2, 19,}

var nums = []int{
	2, 4, 3, 6, 3, 2, 110, 5, 5, -2, 19, 59610, -2, 19, 1110, 1110, 2020, 2020, -99, -99, 56789, 56789, 110, 110, 596, 596,
	45, 378, 958, 45, 110, 378, 958, 59610,
}
var numsRand []int

func init() {
	numsRand = nil
	for i := 0; i < 10000; i++ {
		numsRand = append(numsRand, rand.Int())
	}

	numsRand = append(numsRand, numsRand...)
	numsRand = append(numsRand, 1, 2)

}

//go test -bench=. -run=bench -benchmem -cpuprofile=cpu.out -memprofile=mem.out
//go test -bench=. -run=bench -benchmem
// 测试数据为nums1
//Benchmark_异或求解-4            10000000               220 ns/op             208 B/op          3 allocs/op
//Benchmark_统计求解-4             1000000              1190 ns/op             256 B/op          4 allocs/op
//测试数据为nums
//Benchmark_异或求解-4             3000000               422 ns/op             592 B/op          3 allocs/op
//Benchmark_统计求解-4              500000              3111 ns/op             784 B/op          4 allocs/op
// 测试数据随机20000个
//Benchmark_异或求解-4                5000            240820 ns/op          327696 B/op          3 allocs/op
//Benchmark_统计求解-4                1000           1918945 ns/op          385175 B/op          6 allocs/op
// todo 为什么异或求解比map快这么多？ pprof分析来看 异或运算和map比太快了

func Benchmark_异或求解(b *testing.B) {

	for i := 0; i < b.N; i++ {
		getAppearOnceNumsByXor(numsRand)
	}
}

func Benchmark_统计求解(b *testing.B) {

	for i := 0; i < b.N; i++ {
		getAppearOnceNumsByStatistics(numsRand)
	}
}
