package main

import "fmt"

type S struct {
    a int
    b int
    B Dd
}

type Dd struct {
    B int
    e E
}
type E int

func main() {


    var s =[]string{"1"," ","2"}
    fmt.Println(s)
    fmt.Printf("%v ,%+[1]v ,%#[1]v \n",s)
    ss:=S{}
    fmt.Println(ss)
    fmt.Printf("%v %+[1]v %#[1]v \n",ss)

}
