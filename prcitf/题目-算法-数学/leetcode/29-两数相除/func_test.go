// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _9_两数相除

import (
	"fmt"
	"math"
	"testing"
	"testing/quick"
)

//func myStructProperty(in MyStruct) bool { return in.X == 42 }
//
//func TestCheckProperty(t *testing.T) {
//	reportError("myStructProperty", quick.Check(myStructProperty, nil), t)
//}

func reportError(property string, err error, t *testing.T) {
	if err != nil {
		t.Errorf("%s: %s", property, err)
	}
}

func myDiv(a,b int32)bool{
	return divide(int(a),int(b))==int(a/b)
}

func Test_随机测试(t *testing.T) {
	reportError("error:",quick.Check(myDiv,nil),t)
}

func TestName(t *testing.T) {
	divide1(math.MinInt32,math.MaxInt32)
}

func divide1(dividend int, divisor int) int {
	var a, b = int64(dividend), int64(divisor)
	fmt.Println(a==math.MinInt32, b==math.MaxInt32)
	return 0
}
