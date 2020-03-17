// Copyright (c) liuhao. 2012-2050. All rights reserved.
package 重写once函数

import (
	"sync"
	"sync/atomic"
	"testing"
)

// todo 思考：这是once的实现，那么使用cas 会有问题吗

type Once struct {
	m    sync.Mutex
	done uint32
}
func (o *Once) Do(f func())  {
	if !atomic.CompareAndSwapUint32(&o.done, 0, 1) {
		return
	}
	f()
}


type one int

func (o *one) Increment() {
	*o++
}

func run(t *testing.T, once *Once, o *one, c chan bool) {
	once.Do(func() { o.Increment() })
	if v := *o; v != 1 {
		t.Errorf("once failed inside run: %d is not 1", v)
	}
	c <- true
}

func TestOnce(t *testing.T) {
	o := new(one)
	once := new(Once)
	c := make(chan bool)
	const N = 10
	for i := 0; i < N; i++ {
		go run(t, once, o, c)
	}
	for i := 0; i < N; i++ {
		<-c
	}
	if *o != 1 {
		t.Errorf("once failed outside run: %d is not 1", *o)
	}
}

func TestOncePanic(t *testing.T) {
	var once Once
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("Once.Do did not panic")
			}
		}()
		once.Do(func() {
			panic("failed")
		})
	}()

	once.Do(func() {
		t.Fatalf("Once.Do called twice")
	})
}

// output:BenchmarkOnce-4   	100000000	        19.0 ns/op
func BenchmarkOnce(b *testing.B) {
	var once Once
	f := func() {}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			once.Do(f)
		}
	})
}

// output:BenchmarkOriginOnce-4   	2000000000	         2.74 ns/op
// 原始实现性能更优；(cas的性能在竞争较小的情况下应该不如锁)todo 根据测试cas性能高于锁，但是 once的实现大部分代码走不进锁的处理，性能高的原因应该是load的性能高
func BenchmarkOriginOnce(b *testing.B) {
	var once sync.Once
	f := func() {}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			once.Do(f)
		}
	})
}

