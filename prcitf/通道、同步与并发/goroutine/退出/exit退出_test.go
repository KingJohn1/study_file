// Copyright (c) liuhao. 2012-2050. All rights reserved.
package 退出

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_Exit(t *testing.T) {
	// exit退出会退出整个进程
	go func(){
		fmt.Println("enter ")
		os.Exit(0)
		fmt.Println("afte exit")
	}()

	time.Sleep(100*time.Millisecond)
	fmt.Println("main exit")
}
