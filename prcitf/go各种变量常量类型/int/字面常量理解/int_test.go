// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 字面常量理解

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_SliceIndexType(t *testing.T) {

	convey.Convey(
		"字面常量理解 ", t, func() {
			var index int
			var index1 uint8
			var b int

			// index1-b uint8与int 无法运算，index1-1 可以运算运算结果index1的类型uint8；理解下字面常量为可变的万能整数 类型
			fmt.Printf("%T  ,%[1]v  ", index1-1)
			fmt.Printf("%T  ,%T  ,%T ",1,index-b,index-1)
		},
	)
}

