---
template:       article
title:          Select Sort
date:           2017-07-08 09:00:00 +0800
keywords:       Select Sort
description:    Select Sort
---

# Source
```go
func selectSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		maxIndex := 0
		//寻找最大的一个数，保存索引值
		for j := 1; j < length-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[length-i-1], nums[maxIndex] = nums[maxIndex], nums[length-i-1]
	}
}
```

# Test
```go
func Test_selectSort(t *testing.T) {
	array := []int{7, 3, 1, 5, 9}
	selectSort(array)
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
}
```