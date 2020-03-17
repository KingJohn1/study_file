// Copyright (c) liuhao. 2012-2050. All rights reserved.
package tt

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi_scan函数处理(str string) int {

	s := strings.TrimLeft(str, " ")
	if len(s) == 0 {
		return 0
	}

	var b int32

	var isNegative bool

	if s[0] == '-' {
		isNegative = true
	} else if (s[0] >= '0' && s[0] <= '9') || s[0] == '+' {
		isNegative = false
	} else {
		return 0
	}

	_, err := fmt.Fscanf(strings.NewReader(s), "%d", &b)
	if err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), "overflow")||strings.Contains(err.Error(),"out of range") {
			if isNegative {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return int(b)
}

func myAtoi(str string) int {
	s := strings.TrimLeft(str, " ")
	var d int64

	if len(s) == 0 {
		return 0
	}
	var isNegative bool
	if s[0] == '-' {
		isNegative = true
	} else if (s[0] >= '0' && s[0] <= '9') || s[0] == '+' {
		isNegative = false
	} else {
		return 0
	}

	i := 0
	if s[0] == '+' || s[0] == '-' {
		i = 1
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			if isNegative{
				return int(-d)
			}
			return int(d)
		}

		d *= 10
		d += int64(s[i] - '0')
		if isNegative && -d < math.MinInt32 {

			return math.MinInt32
		}
		if !isNegative && d > math.MaxInt32 {

			return math.MaxInt32
		}
	}

	if isNegative{
		return int(-d)
	}
	return int(d)
}
