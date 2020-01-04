// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package packet_test

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	packet "prcitf/io操作与文本文件处理/pb"
	"testing"
)

func Test_序列化结果(t *testing.T) {

	var (
		s = "liuhao body"
		a = &packet.StringMessage{
			Header: &packet.Header{
				MessageId:            nil,
				Topic:                nil,
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Body:                 &s,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	)
	c,err:= proto.Marshal(a)
	fmt.Println("Marshal 1 ",string(c),"err ",err)

	c1:=proto.MarshalTextString(a)
	fmt.Println("MarshalMessageSetJSON 1 \n",string(c1) ,"err ",err)

	fmt.Println(err)
	d:=&packet.StringMessage{}
	proto.Unmarshal(c,d)
	proto.UnmarshalText(string(c),d)
	fmt.Println("2 ",d)
}
