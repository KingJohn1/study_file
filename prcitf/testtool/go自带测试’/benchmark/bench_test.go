// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package benchmark

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"
	_ "runtime/pprof"
	"testing"
)

func testedFunc() {
	fmt.Sprintf("hello")
}

func Benchmark学习(b *testing.B) {

	//测试函数结束后，打印测试结果
	b.ReportAllocs()
	b.SetParallelism(10)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("内存分配", m.Alloc)
	b.RunParallel(func(pb *testing.PB) {
		// Each goroutine has its own bytes.Buffer.
		var buf bytes.Buffer
		for pb.Next() {
			// The loop body is executed b.N times total across all goroutines.
			buf.Reset()

		}
	})
}

func Test_Benchmark(t*testing.T) {
	benchmarkResult := testing.Benchmark(func(b *testing.B) {
		templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// Each goroutine has its own bytes.Buffer.
			var buf bytes.Buffer
			for pb.Next() {
				// The loop body is executed b.N times total across all goroutines.
				buf.Reset()
				templ.Execute(&buf, "World")
			}
		})
	})

	// fmt.Printf("%8d\t%10d ns/op\t%10d B/op\t%10d allocs/op\n", benchmarkResult.N, benchmarkResult.NsPerOp(), benchmarkResult.AllocedBytesPerOp(), benchmarkResult.AllocsPerOp())
	fmt.Printf("%s\t%s\n", benchmarkResult.String(), benchmarkResult.MemString())
}