// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package rr

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_一个类型的指针的名称是空(t *testing.T) {

	type A struct {
	}
	var a A
	fmt.Println(reflect.TypeOf(a).Name(), " / ", reflect.TypeOf(new(A)).Name()=="", " / ", reflect.Indirect(reflect.ValueOf(new(A))).Type().Name())

}
