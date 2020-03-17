// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _40_不同的子序列2

import (
	"fmt"
	"math"
	"testing"
)

func TestName(t *testing.T) {
	var data = []struct {
		s      string
		expect int
	}{
		{
			s:      "aaa",
			expect: 3,
		},
		{
			s:      "abc",
			expect: 7,
		},
		{
			s:      "aba",
			expect: 6,
		}, {
			s: "blljuffdyfrkqtwfyfztpdiyktrhftgtabxxoibcclbjvirnqyynkyaqlxgyybkgyzvcahmytjdqqtctirnxfjpktxmjkojlvvrr",
			expect: 589192369,
		},
	}
	for i, v := range data {
		if distinctSubseqII(v.s) != v.expect {
			t.Error(i, "error", "got", distinctSubseqII(v.s), "expect", v.expect)
		}
	}
}

func TestN1ame(t *testing.T) {
	fmt.Println(math.MaxInt32>1000000007,math.MaxInt32)
}