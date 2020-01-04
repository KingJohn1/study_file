// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _15_计算右侧小于当前元素的个数

// 逆序对问题
//输入: [5,2,6,1]
//输出: [2,1,1,0]

type ReverseAttr struct {
	index int
	value int
	count int
}

var useRecursive = true

func countSmaller(nums []int) []int {
	if useRecursive {
		return recursiveProcess(nums)
	}
	return dontUseRecursiveProcess(nums)
}

func recursiveProcess(nums []int) []int {
	data := convertToSelfDateStruct(nums)
	data = getReverseOrderNum(data)
	return convertToOriginDateStruct(data)
}

func convertToOriginDateStruct(attrs []ReverseAttr) []int {
	r := make([]int, len(attrs))
	for _, v := range attrs {
		r[v.index] = v.count
	}
	return r
}

func getReverseOrderNum(data []ReverseAttr) (result []ReverseAttr) {
	//defer func() {
	//	fmt.Println(data, result)
	//}()
	if len(data) <= 1 {
		//fmt.Println("len data <=1")
		return data
	}
	midIndex := len(data) / 2
	numL, numR := getReverseOrderNum(data[0:midIndex]), getReverseOrderNum(data[midIndex:])
	if len(numL) == 0 || len(numR) == 0 {
		//	fmt.Println("len(numL) == 0 || len(numR) == 0")
		return append(numL, numR...)
	}
	// 递归结束后再申请内存
	result = make([]ReverseAttr, len(numL)+len(numR))
	lIndex := len(numL) - 1
	rIndex := len(numR) - 1
	index := len(numL) + len(numR) - 1
	cacuCount(lIndex, rIndex, index, numL, numR, result)

	return result
}

func cacuCount(lIndex, rIndex, index int, numL, numR, result []ReverseAttr) {
	for {
		if lIndex == -1 && rIndex == -1 {
			break
		}
		if lIndex == -1 {
			result[index] = numR[rIndex]
			rIndex--
			index--
			continue
		}

		if rIndex == -1 {
			result[index] = numL[lIndex]
			lIndex--
			index--
			continue
		}

		if numL[lIndex].value > numR[rIndex].value {
			numL[lIndex].count += rIndex + 1
			result[index] = numL[lIndex]
			lIndex--
		} else {
			result[index] = numR[rIndex]
			rIndex--
		}
		index--
	}
}

func convertToSelfDateStruct(nums []int) []ReverseAttr {
	r := make([]ReverseAttr, len(nums))
	for i := 0; i < len(nums); i++ {
		r[i].index = i
		r[i].value = nums[i]
	}
	return r
}

func dontUseRecursiveProcess(nums []int) []int {
	data := convertToSelfDateStruct(nums)
	data = dontUseRecusiveGetReverseOrderNum(data)
	return convertToOriginDateStruct(data)
}

func dontUseRecusiveGetReverseOrderNum(data []ReverseAttr) (result []ReverseAttr) {
	if len(data) <= 1 {
		return data
	}
	result = make([]ReverseAttr, len(data))
	stepLenth := 1
	for stepLenth < len(data) {
		for i := stepLenth - 1; i < len(data); i += stepLenth * 2 {
			lIndex := i
			rIndex := i + stepLenth
			if rIndex > len(data)-1 {
				rIndex = len(data) - 1
			}
			index := rIndex
			for {
				if lIndex == i-stepLenth && rIndex == i {
					break
				}
				if lIndex == i-stepLenth {
					result[index] = data[rIndex]
					rIndex--
					index--
					continue
				}

				if rIndex == i {
					result[index] = data[lIndex]
					lIndex--
					index--
					continue
				}

				if data[lIndex].value > data[rIndex].value {
					data[lIndex].count += rIndex - i
					result[index] = data[lIndex]
					lIndex--
				} else {
					result[index] = data[rIndex]
					rIndex--
				}
				index--
			}
			//fmt.Println(data, result)
		}
		stepLenth *= 2
	}
	return
}
