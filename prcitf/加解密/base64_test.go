// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 加解密

import (
	"encoding/base64"
	"fmt"
	"testing"
)

// base64 是对称加密；安全性低；适用安全要求不高的地方
func Test_Base64(t *testing.T) {
	var src = []byte("刘号，你好 sdfas dda ")
	s := base64.StdEncoding.EncodeToString(src)
	fmt.Println(s)
	fmt.Println(base64.RawStdEncoding.EncodeToString(src))
	fmt.Println(base64.RawStdEncoding.WithPadding('@').EncodeToString(src))
	// 解码不关注padding，随意设置合法值都可以解码
	decodeString, _ := base64.RawStdEncoding.WithPadding('@').DecodeString(s)
	fmt.Println(string(decodeString))
}
