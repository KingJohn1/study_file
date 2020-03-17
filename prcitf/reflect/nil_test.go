// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package rr

import (
	"github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func Test_Nil(t *testing.T) {
	type testNil *int
	convey.Convey(
		"nil 不等于 testNil(nil) ",t,func(){
			convey.ShouldBeFalse(reflect.DeepEqual(nil,testNil(nil)))
		},
		)
}
