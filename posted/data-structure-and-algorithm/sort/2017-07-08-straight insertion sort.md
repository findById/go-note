---
template:       article
title:          Straight Insertion Sort
date:           2017-07-08 09:00:00 +0800
keywords:       Straight Insertion Sort
description:    Straight Insertion Sort
---

# Source
```go
func straightInsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			j := i - 1
			temp := array[i]
			for j >= 0 && array[j] > temp {
				array[j+1] = array[j]
				j--
			}
			array[j+1] = temp
		}
	}
}
```

# Test
```go
func Test_straightInsertionSort(t *testing.T) {
	array := []int{7, 3, 1, 5, 9}
	straightInsertionSort(array)
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
}
```