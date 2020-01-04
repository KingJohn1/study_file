// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _06_写字符串需要的行数

func numberOfLines(widths []int, S string) []int {
	if len(S)==0{
		return []int{0,0}
	}
	var line=1
	var lineWidth=0
	for i:=0;i<len(S);i++{
		lineWidth+=widths[S[i]-'a']
		if lineWidth>100{
			line++
			lineWidth=widths[int(S[i]-'a')]
		}
	}
	return []int{line,lineWidth}
}