package main

import (
	"testing"
	"fmt"
)
const (
	temp1=iota
	temp2
)
func TestAAA(t *testing.T){
	type QQ int
	type WW uint32
	var qq QQ =1
	var ww WW=1
	fmt.Printf("%T\n",temp1)
	fmt.Println(qq==1,ww)
	fmt.Println(int(temp2))
}
