---
template:       article
title:          Bubble Sort
date:           2017-07-08 09:00:00 +0800
keywords:       Bubble Sort
description:    Bubble Sort
---

```go
func bubbleSort(array []int) {
	temp := 0
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}
```

```go
func Test_bubbleSort(t *testing.T) {
	array := []int{7, 3, 1, 5, 9}
	bubbleSort(array)
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
}
```