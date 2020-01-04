package main

import "fmt"

type a interface {
	A(a)bool
}
type b interface {
	 a
	B(b)bool
}
type c int
type d struct{
	c

}
func main1() {
 var cc c=c(1)
 var aa a=cc
 fmt.Println(aa.A(aa))
var dd d
dd.c=c(1)
var bb b
bb=dd//a要实现b接口，b接口中还有接口c，a不必实现c，只需要a中有属性实现了接口c
fmt.Println(bb)
}
func (d)B(b)bool{
	return false
}
func (c) A(temp a)bool{//类型实现一个接口时，函数必须实现 比如接口 B（C)不可以实现B（D）
	return  false
}
