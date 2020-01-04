package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func ExampleScanner_emptyFinalToken() {
	// Comma-separated list; last entry is empty.
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// Define a split function that separates on commas.
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	// Scan.
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	// Output: "1" "2" "3" "4" ""
}

func Test_科学计数法表示的浮点数的读入(t *testing.T) {

	var a float64
	r:=strings.NewReader("2.1112E-6")
	fmt.Fscanf(r,"%f",&a)
	fmt.Println(a)
	fmt.Printf("%f %[1]g %[1]e %[1]b",a)
	//output:0.000002 2.1112e-06 2.111200e-06 4984930059317197p-71
}

func Test_含有空格的读入(t *testing.T) {
	var s string
	n, err := fmt.Fscanf(strings.NewReader("  123 123 kdaf"), "%s", &s)
	fmt.Println(n,err)
	fmt.Printf("%#v \n",s)// output123

	var d int
	n, err = fmt.Fscanf(strings.NewReader("  123 123 kdaf"), "%d", &d)
	fmt.Println(n,err)
	fmt.Printf("%#v",d)// output123
}

func Test_读入的整数大小越界(t *testing.T) {
	var d int8
	n,err:=fmt.Fscanf(strings.NewReader("258"),"%d",&d)
	// 越界会丢弃为0
	fmt.Println(n,err)//output:0 integer overflow on token 258
	fmt.Println(d)

	n,err=fmt.Fscan(strings.NewReader(" 25"),&d)
	// 越界会丢弃为0
	fmt.Println(n,err)
	fmt.Println(d)

	var dd int16
	n,err=fmt.Fscanf(strings.NewReader("-258111111111111111"),"%d",&dd)
	// 越界会丢弃为0
	fmt.Println(n,err)//output:0 integer overflow on token 258
	fmt.Println(dd)
}

func Test_读入的浮点数大小越界(t *testing.T) {
	var f float64
	n,err:=fmt.Fscanf(strings.NewReader("12235489266.341234567567"),"%f",&f)
	fmt.Println(n,err)
	fmt.Printf("%f",f)
}