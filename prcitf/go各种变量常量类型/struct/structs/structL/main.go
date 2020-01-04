package main

import "fmt"

func main() {


	b:=&s{1}//实例化结构体
	s1:=s{2}
	var x1,x2 Interface
	x1=b
	x2=s1
	x2=x1
	fmt.Println(x2)
	a:=make(chan int)
	a=make(chan int)
	close(a)
	//a<-2
	a=make(chan int)
	go func(){
		a<-1
	}()
	e:=<-a
	fmt.Println(e)
	var intt []int
	fmt.Println(len(intt))
	intt=append(intt,1)
	fmt.Println(intt)

	//fmt.Println(ti)

}
type  Interface interface {
	t() Interface
}
type s struct {
	a int
}
func (s1 s) t()Interface{
   t:=newt(s1)
   if true{
   	return &s1
   }
   return  t
}
func newt(s1 s)*s{
	return &s1
}
type I int