// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package rr

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Struct(t *testing.T) {
	type Sub struct {
		SubF int
	}
	type S struct {
		I int
		T bool
		B []*Sub
	}

	var s interface{}
	s = &S{
		I: 2,
		T: true,
		B: []*Sub{
			{
				SubF: 1,
			},
		},
	}

	if reflect.ValueOf(s).Kind() == reflect.Ptr {
		fmt.Println(reflect.ValueOf(s).Elem().Field(0))
	}

	var want map[string]interface{} = map[string]interface{}{
		"I": 2, "T": true, "B": []*Sub{
			{
				SubF: 1,
			},
		},
	}

	value1 := reflect.ValueOf(s).Elem()
	//tt := reflect.TypeOf(s).Elem()

	for i, v := range want {
		fmt.Println(i)
		fmt.Println(reflect.DeepEqual(value1.FieldByName(i).Interface(), v))
		//tt.Field(i)
	}

	a := S{I: 2}
	rv := reflect.ValueOf(a)
	rt := reflect.TypeOf(a)
	ok := rv.Interface().(S)
	fmt.Println(ok.I)

	reflect.PtrTo(rt)

}
