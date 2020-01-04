package main

import (
	"reflect"
	"fmt"
	"os"
)

func main()  {
	i:=1
	a:=reflect.TypeOf(i)
	j:=2
	b:=reflect.TypeOf(j)
	fmt.Println(a==b)

	fmt.Println(a)
	x:=2
	d:=reflect.ValueOf(x)
	fmt.Println(d.CanAddr())
	d=reflect.ValueOf(&x)
	fmt.Println(d.CanAddr())
	d=reflect.ValueOf(&x).Elem()
	fmt.Println(d.CanAddr())
	d.Set(reflect.ValueOf(4))
	fmt.Println(x)
	d.SetInt(1)
	t:=d.Addr().Interface().(*int)
	*t=23
	//fmt.Println(d)
	fmt.Println(d)
	stdou:=reflect.ValueOf(os.Stdout).Elem()
	fmt.Println(stdou.Type())
	fd:=stdou.FieldByName("fd")
	fmt.Println(fd)
	fmt.Println(fd.CanAddr(),fd.Addr().CanInterface(),fd.CanSet())
	//fd.Set(reflect.ValueOf(2))
}

