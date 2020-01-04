// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"fmt"
	"sync"
	"testing"
)

type M struct {
	ch chan struct{}
}

func NewMutex() M {
	return M{
		ch: make(chan struct{}, 1),
	}
}

func (m M) Lock() {
	m.ch <- struct{}{}
}

func (m M) UnLock() {
	select {
	case <-m.ch:
	default:
		panic("unlock of unlocked mutex")
	}
}

func (m M) TryLock() bool {
	select {
	case m.ch <- struct{}{}:
		return true
	default:

	}
	return false
}

func (m M) IsLock() bool {
	return len(m.ch) == 1
}

func Test_实现锁(t *testing.T) {
	var mg sync.WaitGroup
	mg.Add(1)
	m := NewMutex()
	go func() {
		if !m.TryLock(){
			m.UnLock()
		}
		m.Lock()
		fmt.Println(m.IsLock())
		m.UnLock()
		mg.Done()
	}()
	fmt.Println(m.TryLock())
	fmt.Println(m.IsLock())

	mg.Wait()
}
