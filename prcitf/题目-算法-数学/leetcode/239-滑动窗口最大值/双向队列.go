// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _39_滑动窗口最大值

import (
	"fmt"
)

type DoubleSidedList struct {
	begin int // 如果没有值 为-1
	end   int
	data  []int
}

func newDoubleSidedList(l int) *DoubleSidedList {
	list := new(DoubleSidedList)
	list.data = make([]int, l)
	list.begin = -1
	list.end = -1
	return list
}

func (l *DoubleSidedList) pushBack(i int) {
	init := false
	if l.begin == -1 {
		l.begin = 0
		init = true
	}
	l.end = (l.end + 1) % len(l.data)
	l.data[l.end] = i
	if !init && l.end == l.begin {
		fmt.Println("list is full")
	}
}

func (l *DoubleSidedList) pushFront(i int) {
	init := false
	if l.end == -1 {
		l.end = 0
		l.begin = 0
		init = true
	}
	if !init {
		l.begin = (l.begin + len(l.data) - 1) % len(l.data)
	}
	l.data[l.begin] = i
	if !init && l.end == l.begin {
		fmt.Println("list is full")
	}
}

func (l *DoubleSidedList) peekFront() int {
	if l.begin == -1 {
		return -1
	}
	return l.data[l.begin]
}

func (l *DoubleSidedList) peekBack() int {
	if l.end == -1 {
		return -1
	}
	return l.data[l.end]
}

func (l *DoubleSidedList) popFront() int {
	if l.begin == -1 {
		fmt.Println("list is empty")
		return -1
	}

	i := l.data[l.begin]
	if l.begin == l.end {
		l.begin, l.end = -1, -1
	} else {
		l.begin = (l.begin + 1) % len(l.data)
	}
	return i
}

func (l *DoubleSidedList) popBack() int {
	if l.end == -1 {
		fmt.Println("list is empty")
		return -1
	}

	i := l.data[l.end]
	if l.begin == l.end {
		l.begin, l.end = -1, -1
	} else {
		l.end = (l.end + len(l.data) - 1) % len(l.data)
	}
	return i
}

type PriotityList struct {
	indexs    *DoubleSidedList //存储下标
	maxLength int
	data      []int
}

func newPriotityList(list *DoubleSidedList, maxLength int, data []int) *PriotityList {
	return &PriotityList{
		indexs:    list,
		maxLength: maxLength,
		data:      data,
	}
}

func (p *PriotityList) initPush() {
	for i := 0; i < p.maxLength; i++ {
		p.PushIndex(i)
	}
}

func (p *PriotityList) PushIndex(index int) {
	for p.indexs.peekBack() != -1 {
		if p.data[ p.indexs.peekBack()] <= p.data[index] {
			p.indexs.popBack()
		} else {
			break
		}
	}
	p.indexs.pushBack(index)

	p.checkAndClearExpiredIndex()
}

func (p *PriotityList) PeekFrontValue() int {
	return p.data[p.indexs.peekFront()]
}

func (p *PriotityList) checkAndClearExpiredIndex() {
	if p.indexs.peekFront() == -1 {
		return
	}

	begin, end := p.indexs.peekFront(), p.indexs.peekBack()
	if end-begin+1 > p.maxLength {
		p.indexs.popFront()
	}
}

// 执行用时 :
//20 ms , 在所有 golang 提交中击败了 97.77% 的用户
//内存消耗 :
//6.3 MB , 在所有 golang 提交中击败了 100.00% 的用户
func maxSlidingWindow(nums []int, k int) []int {
	var maxWindow []int
	list := newPriotityList(newDoubleSidedList(k), k, nums)
	list.initPush()
	if len(nums)!=0{
		maxWindow = append(maxWindow, list.PeekFrontValue())
	}

	for i := k; i < len(nums); i++ {
		list.PushIndex(i)
		maxWindow = append(maxWindow, list.PeekFrontValue())
	}
	return maxWindow
}
