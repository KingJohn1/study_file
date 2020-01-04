package main

import (
	"fmt"
	"net"
)

//
//import (
//	"errors"
//	"fmt"
//	"log"
//)
//
//func a()error{
//	fmt.Println(fmt.Errorf("liuhao"))
//	return errors.New("liuhao")
//}
//
//func main() {
//
//	//var a b
//	//a=A{}
//	//obj,ok:=a.(AA)
//	//fmt.Println(obj,ok)
//	//var a =make(map[string]int ,2)
//	//fmt.Println(len(a))
//	//a["1"]=1
//	//fmt.Println(len(a))
//	//var a=1
//	//var b=1
//	//switch{
//	//case a==b:fmt.Println("2222")
//	//case true:fmt.Println("111")
//	//}
//	//var bb b
//	//bb.get()
//	const a=1
//	a=2
//	fmt.Println(a)
//
//}
//type A struct{
//	AA
//}
//type AA struct {
//
//}
//
//type b interface {
//	get()
//}
//func (a *AA)get(){
//if a==nil{
//	log.Print("liuh")
//}
//}
////func (a*A)get()(n int ,nn int ){
////var b []int
////b=append(b,1)
////
////fmt.Println(b)
////}
////type B A

func main() {

	//a:=byte('$')
	//b:=byte('t')
	ip1:=net.IP{'$',byte('\t'),0200,0,byte('P'),010,byte('o'),0377,0,0,0,0,0,002,0,byte('b')}
	fmt.Println(IsGlobalUnicastByteIP(ip1))
	fmt.Println([]byte(ip1))

	ip:=net.ParseIP("2409:8000:5008:6FFF::2:62")
	fmt.Println([]byte(ip))
	if len(ip) == net.IPv6len {
		fmt.Println(IsGlobalUnicastByteIP(ip))
	}


}
func IsGlobalUnicastByteIP(ip []byte) bool {
	return net.IP(ip).IsGlobalUnicast()
}


