// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _73_最接近原点的K个点

import (
	"container/heap"
)

// 解法一：大根堆
type myHeap []int

func (m myHeap) Len() int {
	return len(m)
}

func (m myHeap) Less(i, j int) bool {
	//fmt.Println(i,j)
	return getDistance(myPoints[m[i]]) > getDistance(myPoints[m[j]])
}

func (m myHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *myHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *myHeap) Pop() (v interface{}) {
	// *c 是值，不是变量，所以不能短变量声明
	//var a []int=[]int{1,1}
	//c:=&a
	//*c,b:=[]int{1,2},2
	*m, v = []int(*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

var myPoints [][]int

func kClosest(points [][]int, K int) (kClosestP [][]int) {
	myPoints = points
	heapK := &myHeap{}
	heap.Init(heapK)
	for i := 0; i < len(points); i++ {
		if i < K {
			heap.Push(heapK, i)
			//fmt.Println(i,heapK,points[(*heapK)[0]])
			continue
		}
		if getDistance(points[i]) < getDistance(points[(*heapK)[0]]) {
			heap.Pop(heapK)
			heap.Push(heapK, i)
		}
		//fmt.Println(i,heapK,points[(*heapK)[0]])
	}
	//fmt.Println(heapK)
	kClosestP = make([][]int, heapK.Len())
	for i, v := range *heapK {
		kClosestP[i] = points[v]
	}
	return
}

func getDistance(point []int) (distance int) {
	//fmt.Println(point[0]*point[0] + point[1]*point[1])
	return point[0]*point[0] + point[1]*point[1]
}

// 分治法：题解看到的解法
//执行用时 : 116 ms , 在所有 golang 提交中击败了 95.35% 的用户
//内存消耗 : 7.3 MB , 在所有 golang 提交中击败了 100.00% 的用户
func kClosest1(points [][]int, K int) (kClosestP [][]int) {

	if len(points) <= K {
		return points
	}

	baseValue := getDistance(points[0])
	lIndex := 0
	rIndex := len(points) - 1
	for {
		for ;lIndex<rIndex;rIndex--{
			if getDistance(points[rIndex]) < baseValue {
				break
			}
		}
		// 比下面实现简洁
		for ;lIndex<rIndex;lIndex++{
			if getDistance(points[lIndex]) > baseValue {
				break
			}
		}
		//for {
		//	if lIndex == rIndex {
		//		break
		//	}
		//	if getDistance(points[lIndex]) > baseValue {
		//		break
		//	}
		//	lIndex++
		//
		//}
		if lIndex == rIndex {
			break
		}
		points[lIndex], points[rIndex] = points[rIndex], points[lIndex]
	}
	points[0], points[lIndex] = points[lIndex], points[0]

	if lIndex+1 == K {
		return points[:K]
	}
	if lIndex+1 > K {
		return kClosest1(points[:lIndex], K)
	}
	kClosest1(points[lIndex+1:], K-lIndex-1)
	return points[:K]
}
