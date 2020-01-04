package main

import "fmt"

type I interface {
	t()
}
type Inter int
func (Inter)t(){

}
func main(){

	var a I=Inter(1)
	if a==Inter(1){
		fmt.Println(true)
	}
	a=nil
	fmt.Printf("%T %V",a)
	fmt.Println(a)
}