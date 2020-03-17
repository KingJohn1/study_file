// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main_test

import (
	"sync"
	"testing"
)

// 可重入锁（ReetrantLock），也叫做递归锁，指的是在同一线程内，
// 外层函数获得锁之后，内层递归函数仍然可以获取到该锁。换一种说法：同一个线程再次进入同步代码时，可以使用自己已获取到的锁。

// 执行此用例，会阻塞。因此sync.mutex是不可重入锁
func Test_Lock(t *testing.T) {
	var mutex sync.Mutex
	mutex.Lock()
	mutex.Lock()
}
