package main

import "fmt"
type(
	T1 struct {
		T2
		T3
	}
	T2 struct {

	}
	T3 struct {

	}

	I interface {
		get()
	}
	II interface {
		get()
	}
)

//func(t*T1)get(){
//	fmt.Println(1)
//}
//
func (t*T2)get()  {
fmt.Println(2)
}

func main(){
	var t II  =&T2{}
	_,ok:=t.(I)
	fmt.Println(ok)
	//_,ok=(*t).(*I)
	//fmt.Println(ok)

	var i *uint8

	fmt.Printf("i %v,i type:%T",i,i);fmt.Printf("i %v,i type:%T",(*uint16)(i),i)

}

//func (t*T3)get(){
//	fmt.Println(3)
//}
//
//func main1(){
//	var a pp.A
//
//
//	var t T1
//	t.get()
//	t.T2.get()
//}