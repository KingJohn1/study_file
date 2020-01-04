// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _39_滑动窗口最大值

import (
	"fmt"
	"testing"
)

func Test_双向队列(t *testing.T) {
	l := newDoubleSidedList(5)
	l.pushBack(1)
	l.pushBack(2)
	l.pushBack(3)
	l.pushBack(4)
	l.pushBack(5)
	fmt.Println(l.popBack(), l.popFront(), l.popBack(), l.popFront(), l.popBack())
	fmt.Println(l)
	l.pushFront(1)
	l.pushFront(2)
	fmt.Println(l.popBack(), l.popBack())
}

func Test_优先队列(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{
		1, 3, -1, -3, 5, 3, 6, 7,
	}, 3))
}
