package unsafe

import (
	"fmt"
	"unsafe"
)

var x struct{
	a bool
	b int16
	c []int
}
func main(){
	fmt.Println(unsafe.Sizeof(x.c),unsafe.Alignof(x.b),unsafe.Alignof(x),unsafe.Offsetof(x.b))
}