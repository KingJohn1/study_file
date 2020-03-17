// Copyright (c) liuhao. 2012-2050. All rights reserved.
package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx2 context.Context) {
		fmt.Println("enter gorountine")
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("done ,", ctx2.Err())
				return
			default:
				time.Sleep(500*time.Millisecond)
				time.AfterFunc(500*time.Millisecond, func() {

					fmt.Println("work")
				})
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1*time.Second)
}
