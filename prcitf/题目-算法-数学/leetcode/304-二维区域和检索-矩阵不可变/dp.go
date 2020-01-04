// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _04_二维区域和检索_矩阵不可变

type NumMatrix struct {
	dp                 [][]int
	matrix             [][]int
	rowLenth, colLenth int
}

func Constructor(matrix [][]int) NumMatrix {
	n := NumMatrix{}
	n.matrix = matrix
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return n
	}
	n.rowLenth, n.colLenth = len(matrix), len(matrix[0])
	n.dp = make([][]int, len(matrix)+1)
	for i := 0; i < len(n.dp); i++ {
		n.dp[i] = make([]int, len(n.matrix[0])+1)
	}

	n.initDp()

	return n
}

func (n NumMatrix) initDp() {
	for i := 0; i < n.rowLenth; i++ {
		for j := 0; j < n.colLenth; j++ {
			n.dp[i+1][j+1] = n.matrix[i][j] + n.dp[i][j+1] + n.dp[i+1][j] - n.dp[i][j]
		}
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int)(result int) {

	defer func() {
		// 输入不合理导致越界
		if err:=recover();err!=nil{
			result=0
		}
	}()

	return this.dp[row2+1][col2+1]-this.dp[row1][col2+1]-this.dp[row2+1][col1]+this.dp[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
