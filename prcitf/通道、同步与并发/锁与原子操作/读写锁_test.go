// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test_RWMutex(t *testing.T) {
	var rwMutex = sync.RWMutex{}
// 上写锁后 再读 会阻塞，符合预期
	rwMutex.Lock()
	fmt.Println("lock")
	rwMutex.RLock()
}
