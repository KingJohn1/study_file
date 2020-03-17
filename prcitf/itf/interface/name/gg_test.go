package name

import (
	"testing"
	"fmt"
)

type AA struct{

}
type itf interface {
	name()
}
func (a AA)name(){
	fmt.Println("name")
}


const(
	test1=iota
	test2
)

type Int uint32
const (
	test3 Int=0

)

func TestName(t *testing.T) {
	var i itf=AA{}
	i.name()
fmt.Println(test1==test3)
}
