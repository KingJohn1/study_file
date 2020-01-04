package main

import "fmt"

func main(){
	//type M map[int]int
	//
	//var m=make(M)
	//m[1]=2
	//t:=new(map[int]int)
	//fmt.Printf("%[1]T,%#[1]v\n",t)
	////(*t)[2]=3
	//type s struct {
	//	a int
	//}
	//mm:=new(s)
	//fmt.Printf("%[1]T,%#[1]v\n",mm)
	//fmt.Println(t)
	//fmt.Println(m)

	//a:=uint64(1)
	//b:=uint32(1)
	//fmt.Println(a&b)
	//fmt.Printf("%T\n",a>>1)

	type t struct {
		A int
	}
	b:=[]t{
		{1},
	}
	fmt.Println(b[0].A)
	c:=map[int]t{1:{1}}
	d:=c[1]
	d.A=1
	fmt.Println(d)
}
