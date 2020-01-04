// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _55_输出二叉树

import (
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//执行用时 :
//0 ms , 在所有 golang 提交中击败了 100.00% 的用户
//内存消耗 :
//2.3 MB , 在所有 golang 提交中击败了 100.00% 的用户

func printTree(root *TreeNode) [][]string {
	var depth int = getTreeDepth(root)
	var bTree = make([][]string, depth)
	for i := 0; i < len(bTree); i++ {
		bTree[i] = make([]string, int(math.Pow(2, float64(depth)))-1)
	}
	if root!=nil{
		setBTree(bTree,depth,depth,root,int(math.Pow(2, float64(depth-1)))-1)
	}

	return bTree
}

func setBTree(bTree [][]string, depth int, currentDepth int, node *TreeNode, currentPos int) {

	bTree[depth-currentDepth][currentPos] = strconv.Itoa(node.Val)

	var spaceing = int(math.Pow(2, float64(currentDepth-2)))
	if node.Left != nil {
		setBTree(bTree, depth, currentDepth-1, node.Left, currentPos-spaceing)
	}
	if node.Right != nil {
		setBTree(bTree, depth, currentDepth-1, node.Right, currentPos+spaceing)
	}

}

// 从上往下调用 从下往上递归
func getTreeDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(getTreeDepth(node.Left), getTreeDepth(node.Right)) + 1
}

// 从上往下调用
func getTreesDepth1(node *TreeNode, parentDepth int, maxDepth *int) {
	if node == nil {
		return
	}
	currentDepth := parentDepth + 1
	if currentDepth > *maxDepth {
		*maxDepth = parentDepth + 1
	}
	getTreesDepth1(node.Left, currentDepth, maxDepth)
	getTreesDepth1(node.Right, currentDepth, maxDepth)
}

func max(depth int, depth2 int) int {
	if depth < depth2 {
		return depth2
	}
	return depth
}