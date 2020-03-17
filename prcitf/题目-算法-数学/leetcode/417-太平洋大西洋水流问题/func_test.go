// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _17_太平洋大西洋水流问题

import (
	"fmt"
	"testing"
)

//  太平洋 ~   ~   ~   ~   ~
//       ~  1   2   2   3  (5) *
//       ~  3   2   3  (4) (4) *
//       ~  2   4  (5)  3   1  *
//       ~ (6) (7)  1   4   5  *
//       ~ (5)  1   1   2   4  *

func TestName(t *testing.T) {
	fmt.Println(pacificAtlantic([][]int{
		{1,2,2,3,5},
		{3,2,3,4,4},
		{2,4,5,3,1},
		{6,7,1,4,5},
		{5,1,1,2,4},
	}))
}
