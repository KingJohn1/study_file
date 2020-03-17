// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"reflect"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0: return nil
	case 1: return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir: reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		reflect.Select(cases)
	}()
	return orDone
}

func fanOutReflect(ch <-chan interface{}, out []chan interface{}) {
	go func() {
		defer func() {
			for i := 0; i < len(out); i++ {
				close(out[i])
			}
		}()
		cases := make([]reflect.SelectCase, len(out))
		for i := range cases {
			cases[i].Dir = reflect.SelectSend
		}

		for v := range ch {
			v := v
			for i := range cases {
				cases[i].Chan = reflect.ValueOf(out[i])
				cases[i].Send = reflect.ValueOf(v)
			}
			for _ = range cases { // for each channel
				chosen, _, _ := reflect.Select(cases)
				cases[chosen].Chan = reflect.ValueOf(nil)
			}
		}
	}()
}
