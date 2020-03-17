// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _15_计算右侧小于当前元素的个数

// 网上的方法。从后往前依次插入2叉搜索树，并记录左右子节点的个数，从root开始找到要插入节点的位置，获取比当前元素小的节点个数。

// 这种实现还没搞明白
// bst 二叉树
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
	LeftSum int
	Dup int
}

func countSmaller1(nums []int) []int {
	n:=len(nums)
	if n==0{
		return nil
	}
	res:=make([]int,n)
	root:=&TreeNode{Val:nums[n-1]}
	for i:=n-1;i>=0;i--{
		res[i] = insertBST(root,nums[i])
	}
	return res
}

func insertBST(node *TreeNode,val int)int{
	sum:=0
	for node.Val!=val{
		if node.Val>val{
			if node.Left==nil{
				node.Left=&TreeNode{Val:val}
			}
			node.LeftSum++
			node = node.Left
		}else{
			sum+=node.LeftSum+node.Dup
			if node.Right ==nil{
				node.Right = &TreeNode{Val:val}
			}
			node = node.Right
		}
	}
	node.Dup++
	return sum+node.LeftSum
}