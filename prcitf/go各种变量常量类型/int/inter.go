package main1

type B struct {

}
//以接口的视角，接口看不到这个函数，因此认为其结构体未实现接口
func(b*B)get(){}
func(b*B)Get(){}


type a interface {
	Geta()
}

type A interface {
	a
}


