// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _5_颜色分类


func sortColors(nums []int) {

	var i, j = 0, len(nums)-1

	for i <= j {

		if nums[i] == 1 {
			i++
			continue
		}

		if nums[i] == 0 {
			exist1 := false
			z := 1
			for ; i-z >= 0; z++ {
				if nums[i-z] == 1 {
					exist1 = true
				} else {
					break
				}
			}
			if exist1 {
				nums[i], nums[i-z+1] = nums[i-z+1], nums[i]
			}
			i++
			continue
		}

		for ; j > i; j-- {
			if nums[j] != 2 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}

		if i == j {
			break
		}
		j--

	}

}
