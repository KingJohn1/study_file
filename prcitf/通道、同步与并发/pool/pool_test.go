// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func Test_Pool(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	var i int
	var pool = sync.Pool{
		New: func() interface{} {
			i++
			return i
		},
	}

	go func() {
		for i := 0; i < 100000; i++ {
			pool.Get()
		}
		wg.Done()
	}()
	wg.Wait()
	pool.Put("nihao")
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	pool.New= func() interface{} {
		return "world"
	}
	fmt.Println(pool.Get())


}

func Test_Pool内存泄漏(t *testing.T) {
	pool := sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

	processRequest := func(size int) {
		b := pool.Get().(*bytes.Buffer)
		time.Sleep(500 * time.Millisecond) // Simulate processing time
		b.Grow(size)
		pool.Put(b)
		time.Sleep(1 * time.Millisecond) // Simulate idle time
	}

	// Simulate a steady stream of infrequent large requests.
	go func() {
		for {
			processRequest(1 << 28) // 256MiB
		}
	}()

	// Simulate a storm of small requests.
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}

	// Continually run a GC and track the allocated bytes.
	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dK\n", i, stats.Alloc/1024)
		time.Sleep(time.Second)
		runtime.GC()
	}
}