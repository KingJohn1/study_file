// Copyright (c) liuhao. 2012-2050. All rights reserved.
package pp

import (
	"testing"
)

func Test_IsMatch(t *testing.T) {
	var data = []struct {
		s, p   string
		expect bool
	}{
		{
			"aa", "a*", true,
		},
		{"ab", ".*", true,},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"", "", true},
		{"cd", "cd*cd", false},
		{"", "c*a*b*", true},
		{"", "c*a*b*d", false},
		{"e", "c*a*b*e", true},
		{"e", "c*a*b*ee", false},
		{"e", "c*a*b*.e", false},
		{"e", "c*a*b*..", false},
		{"e", "c*a*b*.", true},
		{"a", ".*..a*", false},
	}

	for i := range data {
		if isMatch(data[i].s, data[i].p) != data[i].expect {
			t.Errorf("%v input s:%v,p:%v,get:%v,expect%v", i, data[i].s, data[i].p, isMatch(data[i].s, data[i].p), data[i].expect)
		}
	}
}
