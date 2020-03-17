// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main_test

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

func Test_Cas实现自旋锁(t *testing.T) {
	var a int32

	var mutex sync.WaitGroup
	mutex.Add(10)

	for i := 0; i < 10; i++ {
		go func(s string) {
			// lock
			for ; !atomic.CompareAndSwapInt32(&a, 0, 1); {
				runtime.Gosched()
			}
			// process
			fmt.Print("begin ")
			fmt.Print(s)
			fmt.Println(" end ")
			// unlock
			for ; !atomic.CompareAndSwapInt32(&a, 1, 0); {

			}
			mutex.Done()
		}("lock" + strconv.Itoa(i))
	}

	mutex.Wait()

	// 输出可能会改变.begin 处理 end的打印顺序，说明没有并发，自旋锁成功
	//begin lock0 end
	//begin lock9 end
	//begin lock4 end
	//begin lock5 end
	//begin lock1 end
	//begin lock6 end
	//begin lock2 end
	//begin lock7 end
	//begin lock3 end
	//begin lock8 end
}
