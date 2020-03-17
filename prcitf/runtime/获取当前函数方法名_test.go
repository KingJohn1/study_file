// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package runtime

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"runtime"
	"testing"
)

func Test_获取函数方法名行号(t *testing.T) {
	methodName()
	测试读取方法名是否有效()
	convey.Convey(
		"", t, 测试读取方法名是否有效)

}

func Test_闭包不增加调用链数目(t *testing.T) {
	闭包与调用链数()
}

func 闭包与调用链数() {
	f := func() {
		funcName, file, line, ok := runtime.Caller(0)
		if (ok) {
			fmt.Println()
			fmt.Println("func name: " + runtime.FuncForPC(funcName).Name())
			fmt.Printf("file: %s, line: %d\n", file, line)
		}
	}
	f()
}

func 测试读取方法名是否有效() {

	// skip 0 代表当前函数，也是调用runtime.Caller的函数。1 代表上一层调用者，以此类推。
	funcName, file, line, ok := runtime.Caller(1)
	if (ok) {
		fmt.Println()
		fmt.Println("func name: " + runtime.FuncForPC(funcName).Name())
		fmt.Printf("file: %s, line: %d\n", file, line)
	}
}

func methodName() string {
	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if f == nil {
		return "unknown method"
	}
	fmt.Println(f.Name(), f.Entry())
	print(f.Name())
	return f.Name()
}
