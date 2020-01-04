// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package example

import (
	"bufio"
	"fmt"
	"os"
)

func ExampleWriter() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush() // Don't forget to flush!
	// Output: Hello, world!
}