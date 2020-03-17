package main

import (
	"fmt"
	"unsafe"
)

func main(){
	// todo prt +1 其实是加8 bit.地址是 64位需加8
	b:=[]byte{1,2,3,4,5,6,7}//todo b 与bb地址逃逸，所以 b和bbbb 其实只隔了3个字，但是编译器加了一个空白字，不知道原因
    bbbb:=[]byte{1,2,3,4,8}
	fmt.Printf("%d \n",&b)
	fmt.Printf("%d \n",unsafe.Pointer(&b))
	fmt.Printf("%d \n",unsafe.Pointer(&bbbb))
	fmt.Println(uintptr(unsafe.Pointer(&b)))
	a:=(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&b))+8))
	fmt.Println(*a)
	c:=(*[]byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&b))))
	fmt.Println(*c)

	d:=(**int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&b))))
	fmt.Printf("d %x %x %x %d %[4]x\n",d,*d,**d,*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(*d))+1)))
	fmt.Printf("%x %x %x \n" ,*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(d))+8)),*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(d))+16)),
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(d))+40)))

	e:=(**int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&b))))
	fmt.Printf("e %x\n",**e)

	f:=(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(*d)) +16))
	fmt.Printf("f %X\n",*f)

	fmt.Printf("%X\n",*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(*d)) )))
	fmt.Printf("%X\n",*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*d)) +8)))
	fmt.Printf("%X\n",*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(*d)) -8)))
	fmt.Printf("%X\n",*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(*d)) +16)))
	fmt.Printf("%X\n",*(*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(*d)) -16)))
	testType()
}

func testType() {
	const (
		a = uint16(iota)
		b
		c
	)
	fmt.Println(a, b, c)
	fmt.Printf("%T,%T", a, b)
}
