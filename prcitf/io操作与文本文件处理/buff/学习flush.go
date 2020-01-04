// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var buf *bufio.ReadWriter = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	buf.WriteString("hello world")

	fmt.Println("begin sleep")
	time.Sleep(100 * time.Millisecond)
	buf.Flush() //写入并不是真的写入指定地方，而是先写入缓存区，flush 强制缓冲区写入指定位置。（这里是标准输出）
}
