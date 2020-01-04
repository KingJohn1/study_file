package main

import (
	"fmt"
	"sort"
)

type B []int
type ByteSlice []byte

func (b ByteSlice) Len() int {
	return len(b)
}
func (b ByteSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByteSlice) Less(i, j int) bool {
	return b[i] < b[j]
}
func main() {

	var a []int = []int{6, 3, 4, 5}
	//a:=make([]int,4)
	sort.IntSlice(a).Sort()
	sort.Ints(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	fmt.Println(a)
	fmt.Printf("%V\n", a)

	c := B(a)
	fmt.Println(c)
	s := "34324dfg"
	b := []byte(s)
	sort.Sort(ByteSlice(b))
	fmt.Printf("%s\n", b)

}
func (b B) e() (int ) {
	return b[1]
}
