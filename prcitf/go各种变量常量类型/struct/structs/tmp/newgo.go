package tmp

type A struct {

	b
}
type b struct {

}

type Itf interface {
	Get()b
}
func (a b)Get()b{
	return a
}
