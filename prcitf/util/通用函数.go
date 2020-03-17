// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package util

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

var HasPanic = func() {
	if err := recover(); err != nil {
		// skip 0 代表当前函数，也是调用runtime.Caller的函数。1 代表上一层调用者，以此类推。
		pc, file, line, ok := runtime.Caller(1)
		// 返回已skip为基准的函数，调用深度 go修复mulit reader的mr：https://github.com/golang/go/commit/ccdca832c569727f7985966a3324421a69739f57
		//runtime.Callers(skip,pc)
		if ok {
			fmt.Println("expect panic and has panic happen", "func name: "+runtime.FuncForPC(pc).Name())
			fmt.Printf("file: %s, line: %d\n", file, line)

		}
	} else {
		log.Fatal("预期panic，但是没有panic")
	}
}

func SleepLit() {
	time.Sleep(100 * time.Millisecond)
}

