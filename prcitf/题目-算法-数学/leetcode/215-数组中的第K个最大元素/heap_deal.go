// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _15_数组中的第K个最大元素

type IntHeap struct {
	data  []int
	lenth int
}

func findKthLargest(nums []int, k int) int {
	heap := InitHeap(nums, k)
	for i := k; i < len(nums); i++ {
		heap.TryReplace(nums[i])
	}
	return heap.data[0]
}

// 实现堆
func InitHeap(nums []int, lenth int) IntHeap {

	heap := IntHeap{
		data:  nums,
		lenth: lenth,
	}
	for i := lenth/2 - 1; i >= 0; i-- {
		heap.down(i)
	}
	//fmt.Println(heap)
	return heap
}

// i<j 则是小的在数组前面，实现小根堆。
func (h *IntHeap) Less(i, j int) bool {
	return h.data[i] < h.data[j]
}

func (h *IntHeap) LessNum(i, j int) bool {
	return i < j
}

func (h *IntHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

//
//func down(h Interface, i0, n int) bool {
//	i := i0
//	for {
//		j1 := 2*i + 1
//		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
//			break
//		}
//		j := j1 // left child
//		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
//			j = j2 // = 2*i + 2  // right child
//		}
//		if !h.Less(j, i) {
//			break
//		}
//		h.Swap(i, j)
//		i = j
//	}
//	return i > i0
//}
func (h *IntHeap) down(parent int) {

	if parent >= h.lenth {
		return
	}

	for {
		minSon := 2*parent + 1
		if minSon >= h.lenth {
			return
		}
		if j1 := 2*parent + 2; j1 < h.lenth {
			if h.Less(j1, minSon) {
				minSon = j1
			}
		}
		if !h.Less(minSon, parent) {
			break
		}
		h.Swap(minSon, parent)
		parent = minSon
	}
}

//func up(h Interface, j int) {
//	for {
//		i := (j - 1) / 2 // parent
//		if i == j || !h.Less(j, i) {
//			break
//		}
//		h.Swap(i, j)
//		j = i
//	}
//}
func (h *IntHeap) up1(son int) {
	if son >= h.lenth {
		return
	}

	for {
		if son == 0 {
			break
		}
		parent := (son - 1) / 2

		if parent == son || !h.Less(son, parent) {
			break
		}

		if h.Less(son, parent) {
			h.Swap(son, parent)
			son = parent
		} else {
			break
		}
	}

}

// 这样写更简洁，逻辑也相等
func (h *IntHeap) up(son int) {
	if son >= h.lenth {
		return
	}

	for {
		parent := (son - 1) / 2
		if parent == son || !h.Less(son, parent) {
			break
		}
		h.Swap(son, parent)
		son = parent

	}
}

func (h *IntHeap) Push(num int) {
	h.lenth++
	if h.lenth > len(h.data) {
		h.data = append(h.data, num)
	}
	h.data[h.lenth-1] = num
	h.up(h.lenth - 1)
}

func (h *IntHeap) Pop() int {
	if h.lenth <= 0 {
		return -1
	}
	h.Swap(0, h.lenth-1)
	h.lenth--
	return h.data[h.lenth]
}

func (h *IntHeap) TryReplace(num int) bool {
	//defer fmt.Println("replace nums:",h.data)
	if h.lenth <= 0 {
		return false
	}
	if !h.LessNum(num, h.data[0]) {
		h.data[0] = num
		h.down(0)
		return true
	}
	return false
}
