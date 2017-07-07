---
template:       article
title:          Quick Sort
date:           2017-07-08 09:00:00 +0800
keywords:       Quick Sort
description:    Quick Sort
---

```go
func quickSort(array []int, start, end int) {
	if start < end {
		i, j := start, end
		key := array[(start+end)/2]
		for i <= j {
			for array[i] < key {
				i++
			}
			for array[j] > key {
				j--
			}
			if i <= j {
				array[i], array[j] = array[j], array[i]
				i++
				j--
			}
		}
		if start < j {
			quickSort(array, start, j)
		}
		if end > i {
			quickSort(array, i, end)
		}
	}
}
```

```go
func Test_quickSort(t *testing.T) {
	array := []int{7, 3, 1, 5, 9}
	quickSort(array, 0, len(array)-1)
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d ", array[i])
	}
}
```