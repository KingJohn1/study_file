// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 测试参数

import (
	"flag"
	"testing"
)

func TestShort(t *testing.T) {
	flag.Parse()
	if testing.Short(){
		t.Skip("跳过")
	}
	t.FailNow()
}
