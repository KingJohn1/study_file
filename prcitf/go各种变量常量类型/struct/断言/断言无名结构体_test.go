// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 断言

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type S struct {
	a int
}

func (s S) get() {}

func Test_无名结构体(t *testing.T) {

	var s = struct {
		a int
	}{1}

	var i interface{} = s
	// 断言失败
	_, ok := i.(S)
	assert.NotEqual(t, true, ok)

	var ss = S{1}
	// 匿名结构体与有名结构体 也可以相等.即便方法数相同
	assert.True(t, ss == s)
}
