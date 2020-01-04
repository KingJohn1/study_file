package main

import "fmt"

func main(){
	var m=make(map[int]int,10)
	for i:=0;i<10;i++{
		m[i]=i
	}

	var d=[]int{1,2,5}
	for i:=range m{
		for j:=0;j<3;j++{
			if i==d[j]{
				delete(m,i)
			}
		}
	}
	fmt.Printf("%+v",m)
}