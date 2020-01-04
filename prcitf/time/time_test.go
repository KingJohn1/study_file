// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package time

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_时间格式(tt *testing.T) {
	var en = make(map[string]string)
	var cn = make(map[string]string)
	en["time_zone"] = "America/Chicago"
	cn["time_zone"] = "Asia/Shanghai"

	loc, _ := time.LoadLocation(cn["time_zone"])
	t := time.Now()
	t = t.In(loc)
	fmt.Println(t.Format(time.RFC3339), t.UTC(), t.Unix(), )
	fmt.Println(t.Truncate(time.Duration(t.Unix())))
}

func Test_After(t *testing.T) {
	fmt.Println(time.Now())
	tc := time.Tick(1000 * time.Millisecond)
	fmt.Println(time.Now())

	for {
		select {
		case <-tc:
			fmt.Println("tc")
		}
	}
}
func Test_AfterTimeFunc(t *testing.T) {
	var mu sync.WaitGroup
	fmt.Println(time.Now().UTC())

	mu.Add(1)
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("timeout:", time.Now().UTC())
		mu.Done()
	})
	mu.Wait()
}

func Test_TimeFormat(t*testing.T){
	tt, err := time.Parse(time.RFC1123Z, "11:06:39")
	fmt.Println(tt,err)
}
