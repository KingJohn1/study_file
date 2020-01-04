// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"fmt"
	"prcitf/util"
	"testing"
)

func Test_Pipeline(t *testing.T) {

	p := func(ch <-chan int) <-chan int {
		out := make(chan int)

		go func() {
			defer close(out)
			a := <-ch
			out <- a * a
		}()
		return out
	}

	var ch = make(chan int, 1)

	go func() {
		ch <- 2
		a := <-p(ch)
		ch <- a
	}()
	util.SleepLit()
	fmt.Println(<-p(ch))

}
