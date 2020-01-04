// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 测试参数

import (
	"testing"
	"time"
)

func TestExpire(t *testing.T) {
// 可以设置测试参数 测试用例最大时间，超过这个时间将panic ;cmd: go test -timeout=1s -v
time.Sleep(2*time.Second)

}
