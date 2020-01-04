// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"sync"
	"testing"
)
// 不推荐值拷贝，但可以运行
func Test_CopyMutex(t *testing.T) {
	var m sync.Mutex
	var b sync.Mutex
	b=m
	b.Lock()
	b.Unlock()
}

type Temp struct {
	lock sync.Mutex
}

func (t *Temp) Lock() {
	t.lock.Lock()
}

func (t Temp) Unlock() {
	t.lock.Unlock()
}

func Test_ErrorCopyMutex(tt *testing.T) {
	t := Temp{lock: sync.Mutex{}}
	t.Lock()
	// 是值拷贝，不会解锁
	t.Unlock()
	t.Lock()
}


type Temp1 struct {
	lock sync.Mutex
}

func (t *Temp1) Lock() {
	t.lock.Lock()
}

func (t *Temp1) Unlock() {
	t.lock.Unlock()
}
func Test_ErrorCopyMutex1(tt *testing.T) {
	// copy后可能会出现重复lock的问题
	var a sync.Mutex
	var ch =make(chan int,0)
	var t Temp1
	go func() {
		a.Lock()
		t.lock=a
		ch<-1
	}()
	<-ch
	t.Lock()
	t.Unlock()
}