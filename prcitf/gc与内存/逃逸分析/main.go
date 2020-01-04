// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

type test struct{}

// go的变量不止存储在堆和栈上，还有别的位置如 .bss .data .noptrdata .noptrbss。linux 内存模型，全局变量并且初始化了 存在data，全局变量为初始化存储在bss
// 这里a存在data 中因此不涉及逃逸与不逃逸

//Section Headers:
//  [Nr] Name              Type             Address           Offset
//       Size              EntSize          Flags  Link  Info  Align
//  [ 0]                   NULL             0000000000000000  00000000
//       0000000000000000  0000000000000000           0     0     0
//  [ 1] .text             PROGBITS         0000000000401000  00001000
//       0000000000081505  0000000000000000  AX       0     0     16
//  [ 2] .rodata           PROGBITS         0000000000483000  00083000
//       0000000000041a7e  0000000000000000   A       0     0     32
//  [ 3] .shstrtab         STRTAB           0000000000000000  000c4a80
//       000000000000010b  0000000000000000           0     0     1
//  [ 4] .typelink         PROGBITS         00000000004c4ba0  000c4ba0
//       0000000000000b48  0000000000000000   A       0     0     32
//  [ 5] .itablink         PROGBITS         00000000004c56e8  000c56e8
//       0000000000000040  0000000000000000   A       0     0     8
//  [ 6] .gosymtab         PROGBITS         00000000004c5728  000c5728
//       0000000000000000  0000000000000000   A       0     0     1
//  [ 7] .gopclntab        PROGBITS         00000000004c5740  000c5740
//       000000000004e207  0000000000000000   A       0     0     32
//  [ 8] .noptrdata        PROGBITS         0000000000514000  00114000
//       000000000000cbfc  0000000000000000  WA       0     0     32
//  [ 9] .data             PROGBITS         0000000000520c00  00120c00
//       0000000000006b10  0000000000000000  WA       0     0     32
//  [10] .bss              NOBITS           0000000000527720  00127720
//       000000000001c688  0000000000000000  WA       0     0     32
//  [11] .noptrbss         NOBITS           0000000000543dc0  00143dc0
//       0000000000002698  0000000000000000  WA       0     0     32
//  [12] .debug_abbrev     PROGBITS         0000000000547000  00128000
//       00000000000001b5  0000000000000000           0     0     1
//  [13] .debug_line       PROGBITS         00000000005471b5  001281b5
//       00000000000106e8  0000000000000000           0     0     1
//  [14] .debug_frame      PROGBITS         000000000055789d  0013889d
//       00000000000120dc  0000000000000000           0     0     1
//  [15] .debug_pubnames   PROGBITS         0000000000569979  0014a979

//go tool link -h 查看链接参数 go tool compile -h 查看编译参数

var a *int
var bbb = make([]int, 1)
var ccc = new([]int)

func main() {
	bbb = append(bbb, 1, 2, 3)
	t1M, _ := test1()
	t2M := test2()
	_ = 1

	b := 1
	a = &b

	println("t1M", &t1M, "t2M", &t2M)
}

// 局部变量切片不会逃逸，但切片所指向的具体内容可能逃逸到heap
// 如果不逃逸，即在栈上分配。a 调用b，a先入栈 b引用a的切片 不会逃逸，但a引用b的切片，栈的性质不能满足a的引用。所以逃逸到heap上
func test1() (test, []int) {
	t1 := test{}
	tslice := make([]int, 2)
	tslice = append(tslice, 1, 2, 3, 4)

	println("t1", &t1, &tslice)
	return t1, tslice
}

func test2() *test {
	t2 := test{}
	println("t2", &t2)
	return &t2
}

//D:\GO\src\prcitf\gc与内存\逃逸分析>go run -gcflags "-m -l"  main.go
//# command-line-arguments
//.\main.go:14:16: test1 &t1 does not escape
//.\main.go:21:9: &t2 escapes to heap
//.\main.go:19:2: moved to heap: t2
//.\main.go:20:16: test2 &t2 does not escape
//.\main.go:9:17: main &t1M does not escape
//.\main.go:9:30: main &t2M does not escape
//t1 0xc000031f58
//t2 0x4d69c0
