package main

import "fmt"

type t struct {
	A int
	B int
}

var t1 = t{1, 2}
var t2 = t{1, 2,}
var t3 = t{
	1, 2,
}
var t4 = t{A: 1, B: 2}
var t5 = t{A: 1, B: 2,}
var t6 = t{
	A: 1,
	B: 2,
}

type A struct {
	a int
	B int
}

type B struct {
	a int
}
type C interface {
	fet()
}
type list int

func main() {
	list1 := 1
	list2 := list(list1)
	var a A
	var b B
	b = B(a)
	fmt.Println(b, a)
}
