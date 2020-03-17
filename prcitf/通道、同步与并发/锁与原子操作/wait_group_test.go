// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"fmt"
	"prcitf/util"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func Test_多次add(t *testing.T) {

	var count int32
	wg:= sync.WaitGroup{}

	for i:=0;i<100;i++{
		// add 可以多次累计，并不是一次赋值
		wg.Add(1)
		go func(){
			// 不可以在这里wg.add 因为gorounte可能没有执行add main就退出了 见下面用例
			atomic.AddInt32(&count,1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func Test_多次add_位置不对_效果增强(t *testing.T) {

	var count int32
	wg:= sync.WaitGroup{}

	for i:=0;i<100;i++{
		go func(){
			// 不可以在这里wg.add 因为gorounte可能没有执行add main就退出了
			// 为了效果加休眠。
			time.Sleep(1*time.Second)
			wg.Add(1)
			atomic.AddInt32(&count,1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func Test_多次add_位置不对(t *testing.T) {

	var count int32
	wg:= sync.WaitGroup{}

	for i:=0;i<100;i++{
		go func(){
			// 不可以在这里wg.add 因为gorounte可能没有执行add main就退出了
			wg.Add(1)
			atomic.AddInt32(&count,1)
			wg.Done()
		}()
	}
	// 每次打印可能都不相同，大概率为90多
	wg.Wait()
	fmt.Println(count)
}

func Test_可以多次wait(t *testing.T) {

	var count int32
	wg:= sync.WaitGroup{}

	for i:=0;i<100;i++{
		// add 可以多次累计，并不是一次赋值
		wg.Add(1)
		go func(){
			// 不可以在这里wg.add 因为gorounte可能没有执行add main就退出了
			atomic.AddInt32(&count,1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(count)
	// 可以多次wait，并不会阻塞。即wait可重用
	wg.Wait()
	fmt.Println(count)
}

func Test_多次Done_会panic(t *testing.T) {
	defer util.HasPanic()

	var count int32
	wg:= sync.WaitGroup{}

	for i:=0;i<100;i++{
		// add 可以多次累计，并不是一次赋值
		wg.Add(1)
		go func(){
			// 不可以在这里wg.add 因为gorounte可能没有执行add main就退出了
			atomic.AddInt32(&count,1)
			wg.Done()
		}()
	}
	wg.Done()
	wg.Wait()
	fmt.Println(count)
}

func Test_AddWait并发调用(t *testing.T) {

	// 可能会panic；因为test以退出 goroutine还持有wg的资源
	// todo 为什么是偶现panic 呢？不应该是必现吗
	var wg=sync.WaitGroup{}
	var a ,b int32

	for i:=0;i<100;i++{
		go func(){
			for{
			wg.Add(1)
			a++
			atomic.AddInt32(&b,1)
			wg.Done()
			}
		}()
	}

	for i:=0;i<100;i++{
		go func() {
			for  {
				wg.Wait()
			}
		}()
	}
	wg.Wait()
	fmt.Println(a,b)
	//util.SleepLit()
}