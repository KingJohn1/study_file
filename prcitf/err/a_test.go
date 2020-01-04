package main

import (
	"testing"
	"fmt"
)

type D interface {
}

func TestA(t *testing.T) {
	//var a []int
	var b = make(map[int]string)
	b[1] = "222222"
	delete(b,0)
	fmt.Println(b)
	//
	//for a2 := range b {
	//
	//	fmt.Printf("a2:%v", a2)
	//	fmt.Println()
	//}
	//
	//var a = []int{4, 5}
	//for aa := range a {
	//
	//	fmt.Println(aa)
	//}
}
