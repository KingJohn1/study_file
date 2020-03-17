// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _17_太平洋大西洋水流问题

var iMax,jMax int

// 思想：从大西洋边缘向高处遍历，所有的点都是可以到达大西洋，从太平洋出发同理；如果某个点大西洋太平洋都可以到达，则为目标点之一
func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	if len(matrix[0]) == 0 {
		return nil
	}
	iMax=len(matrix)
	jMax=len(matrix[0])
	alreadyRead := make([][]int, len(matrix))
	for i := 0; i < len(alreadyRead); i++ {
		alreadyRead[i] = make([]int, len(matrix[i]))
	}
	// 第一列
	for i := 0; i < len(matrix); i++ {
		dfs(i, 0, matrix, alreadyRead)
	}
	// 第一行
	for j := 0; j < len(matrix[0]); j++ {
		dfs(0, j, matrix, alreadyRead)
	}

	alreadyReadAnother := make([][]int, len(matrix))
	for i := 0; i < len(alreadyReadAnother); i++ {
		alreadyReadAnother[i] = make([]int, len(matrix[i]))
	}
	//最后一列
	for i := 0; i < len(matrix); i++ {
		dfs(i, len(matrix[0])-1, matrix, alreadyReadAnother)
	}
	// 最后一行
	for j := 0; j < len(matrix[0]); j++ {
		dfs(len(matrix)-1, j, matrix, alreadyReadAnother)
	}
	result := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if alreadyRead[i][j] == alreadyReadAnother[i][j] && alreadyReadAnother[i][j] == 1 {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func dfs(i, j int, matrix, alreadyRead [][]int) {

	iValid := func(i int) bool {
		if i < 0 || i >= len(matrix) {
			return false
		}
		return true
	}
	jValid := func(j int) bool {
		if j < 0 || j >= len(matrix[0]) {
			return false
		}
		return true
	}
	alreadyRead[i][j] = 1

	if iValid(i-1) && alreadyRead[i-1][j] == 0 && matrix[i-1][j] >= matrix[i][j] {
		dfs(i-1, j, matrix, alreadyRead)
	}
	if iValid(i+1) && alreadyRead[i+1][j] == 0 && matrix[i+1][j] >= matrix[i][j] {
		dfs(i+1, j, matrix, alreadyRead)
	}

	if jValid(j-1) && alreadyRead[i][j-1] == 0 && matrix[i][j-1] >= matrix[i][j] {
		dfs(i, j-1, matrix, alreadyRead)
	}
	if jValid(j+1) && alreadyRead[i][j+1] == 0 && matrix[i][j+1] >= matrix[i][j] {
		dfs(i, j+1, matrix, alreadyRead)
	}
}


func iValid (i int) bool {
	if i < 0 || i >= iMax {
		return false
	}
	return true
}
func jValid (j int) bool {
	if j < 0 || j >= jMax {
		return false
	}
	return true
}