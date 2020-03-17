package main

import "fmt"

type A interface {
	te(y int)
}

type  Aimp struct {

}

func(b Aimp)te(a int){

}

func main1(){
	var a =Aimp{}
	a.te(1)

	var b A
	b=a

	b.te(2)

	fmt.Println(b,"")
}