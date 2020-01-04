// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _15_数组中的第K个最大元素

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"testing"
	"testing/quick"
)

type Input struct {
	nums []int
	k    int
}

func TestCorrect(t *testing.T) {
	err := quick.Check(func(nums []int) bool {
		if len(nums) == 0 {
			return true
		}
		setInt32 := new(big.Int).SetInt64(int64(len(nums)))
		n, err := rand.Int(rand.Reader, setInt32)
		if err != nil {
			fmt.Println(err)
			return false
		}
		k := int(n.Int64())
		fmt.Println(k,findKthLargest(nums,k),findKthLargestByQuickSort(nums, k),findKthLargestBySort(nums,k))
		fmt.Println(nums)
		return findKthLargest(nums, k) == findKthLargestByQuickSort(nums, k)

	}, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

var randomData []Input

func genarateRandomInput(num int) {

	randomData = make([]Input, num)
	for i := 0; i < num; i++ {

	}
}

var data = []struct {
	nums   []int
	k      int
	output int
}{
	{
		nums:   []int{3, 2, 1, 5, 6, 4},
		k:      2,
		output: 5,
	},
	{
		nums:   []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		k:      4,
		output: 4,
	},
}

func TestFindKthLargest(t *testing.T) {
	for i := range data {
		actual := findKthLargest(data[i].nums, data[i].k)
		if actual != data[i].output {
			t.Errorf("index:%v\tnums:%v\tk:%v\texpect:%v\tactual:%v\n", i, data[i].nums, data[i].k, data[i].output,
				actual)
		}
	}
}

func TestFindKthLargestByQuickSort(t *testing.T) {
	for i := range data {
		actual := findKthLargestByQuickSort(data[i].nums, data[i].k)
		if actual != data[i].output {
			t.Errorf("index:%v\tnums:%v\tk:%v\texpect:%v\tactual:%v\n", i, data[i].nums, data[i].k, data[i].output,
				actual)
		}
	}
}

func TestCompare(t *testing.T) {
	fmt.Println(411977409308886866>(1<<61-1),411977409308886866,1<<31,math.MaxInt64)
}