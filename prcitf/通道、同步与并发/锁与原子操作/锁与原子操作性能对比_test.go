// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"sync"
	"sync/atomic"
	"testing"
)

// 在将once的实现由锁改为cas后发现性能下降很多，因此测试下不同竞争情况下锁与cas的性能 -通道、同步与并发/once/重写once函数/改为cas实现_test.go
// 后发现cas的性能高于mutex。猜测once性能高的原因是 经常读值是否为1，是因为load的性能高于cas
var a, c uint32
var m sync.Mutex
//并发4 100000000	        17.3 ns/op	       0 B/op	       0 allocs/op
func Benchmark_Cas(b *testing.B) {
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if !atomic.CompareAndSwapUint32(&a, 0, 1) {
			}
		}
	})
	b.StopTimer()
	fmt.Println(a)
}

 // cas的性能大于锁
func TestBenchMark_CasCompareMutex(t *testing.T) {

	for i := 0; i < 10; i++ {
		a,c=0,0
		parallelNums := int(math.Pow(2, float64(i)))
		benchmarkCas := testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			b.SetParallelism(parallelNums)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if !atomic.CompareAndSwapUint32(&a, 0, 1) {
					}
				}
			})
		})
		fmt.Printf("cas   parallelNums:%v\t%s\t%s\n", parallelNums, benchmarkCas.String(), benchmarkCas.MemString())
		benchmarkMutex :=  testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					m.Lock()
					if c == 0 {
						c++
					}
					m.Unlock()
				}
			})
		})
		fmt.Printf("mutex parallelNums:%v\t%s\t%s\n", parallelNums, benchmarkMutex.String(), benchmarkMutex.MemString())
		assert.True(t,a==1)
		assert.True(t,c==1)
	}
}

// cas的性能远小于loadint
func TestBenchMark_CasCompareLoadAndMutex(t *testing.T) {

	for i := 0; i < 10; i++ {
		a,c=0,0
		parallelNums := int(math.Pow(2, float64(i)))
		benchmarkCas := testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			b.SetParallelism(parallelNums)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if !atomic.CompareAndSwapUint32(&a, 0, 1) {
					}
				}
			})
		})
		fmt.Printf("cas   parallelNums:%v\t%s\t%s\n", parallelNums, benchmarkCas.String(), benchmarkCas.MemString())
		benchmarkMutex :=  testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					if atomic.LoadUint32(&c)==0{
						m.Lock()
						if c == 0 {
							c++
						}
						m.Unlock()
					}
				}
			})
		})
		fmt.Printf("mutex parallelNums:%v\t%s\t%s\n", parallelNums, benchmarkMutex.String(), benchmarkMutex.MemString())
		assert.True(t,a==1)
		assert.True(t,c==1)
	}

	//output:
	//cas   parallelNums:1	100000000	        17.9 ns/op	       0 B/op	       0 allocs/op
	//mutex parallelNums:1	2000000000	         0.81 ns/op	       0 B/op	       0 allocs/op
	//cas   parallelNums:2	100000000	        18.3 ns/op	       0 B/op	       0 allocs/op
	//mutex parallelNums:2	2000000000	         0.76 ns/op	       0 B/op	       0 allocs/op
	//cas   parallelNums:4	100000000	        17.5 ns/op	       0 B/op	       0 allocs/op
	//mutex parallelNums:4	2000000000	         0.76 ns/op	       0 B/op	       0 allocs/op
	//cas   parallelNums:8	100000000	        17.4 ns/op	       0 B/op	       0 allocs/op
	//mutex parallelNums:8	2000000000	         0.75 ns/op	       0 B/op	       0 allocs/op
	//cas   parallelNums:16	100000000	        17.6 ns/op	       0 B/op	       0 allocs/op
	//mutex parallelNums:16	2000000000	         0.79 ns/op	       0 B/op	       0 allocs/op
	//cas   parallelNums:32	100000000	        17.6 ns/op	       0 B/op	       0 allocs/op
	//mutex parallelNums:32	2000000000	         0.82 ns/op	       0 B/op	       0 allocs/op
	//cas   parallelNums:64	100000000	        18.1 ns/op	       0 B/op	       0 allocs/op
	//mutex parallelNums:64	2000000000	         0.73 ns/op	       0 B/op	       0 allocs/op
	//cas   parallelNums:128	100000000	        16.5 ns/op	       0 B/op	       0 allocs/op
	//mutex parallelNums:128	2000000000	         0.77 ns/op	       0 B/op	       0 allocs/op
	//cas   parallelNums:256	100000000	        17.8 ns/op	       0 B/op	       0 allocs/op
	//mutex parallelNums:256	1000000000	         0.80 ns/op	       0 B/op	       0 allocs/op
	//cas   parallelNums:512	100000000	        16.9 ns/op	       0 B/op	       0 allocs/op
	//mutex parallelNums:512	2000000000	         0.85 ns/op	       0 B/op	       0 allocs/op
}

//并发4 20000000	       131 ns/op	       0 B/op	       0 allocs/op
func Benchmark_Mutex(b *testing.B) {
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Lock()
			if a == 0 {
				a++
			}
			m.Unlock()
		}
	})
	b.StopTimer()
	fmt.Println(a)
}
