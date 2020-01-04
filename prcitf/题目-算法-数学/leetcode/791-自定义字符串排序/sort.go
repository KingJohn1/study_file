// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _91_自定义字符串排序

import "sort"

//S = "cba"
//T = "abcd"
//输出: "cbad"

type MyString []byte

func (m MyString) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MyString) getString() string {
	return string(m)
}

func (m MyString) Len() int {
	return len(m)
}

func (m MyString) Less(i, j int) bool {
	return seq[int(m[i]-'a')] < seq[int(m[j]-'a')]
}

var seq []int

func customSortString(S string, T string) string {
	seq = make([]int, 26)
	for i := 0; i < len(S); i++ {
		seq[int(S[i]-'a')] = i + 1
	}

	var s = make(MyString, len(T))
	for i := 0; i < len(T); i++ {
		s[i] = T[i]
	}

	sort.Sort(s)
	return s.getString()
}
