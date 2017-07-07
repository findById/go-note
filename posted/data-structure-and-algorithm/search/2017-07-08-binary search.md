---
template:       article
title:          Binary Search
date:           2017-07-08 09:00:00 +0800
keywords:       Binary Search
description:    Binary Search
---

# 二分查找
#### 二分查找又称折半查找，它是一种效率较高的查找方法
#### 必须采用顺序存储结构
#### 必须按关键字大小有序排列
### 说明：元素必须是有序的，如果是无序的则要先进行排序操作。
### 基本思想：也称为是折半查找，属于有序查找算法。用给定值k先与中间结点的关键字比较，中间结点把线形表分成两个子表，若相等则查找成功；若不相等，再根据k与该中间结点关键字的比较结果确定下一步查找哪个子表，这样递归进行，直到查找到或查找结束发现表中没有这样的结点。
### 复杂度分析：最坏情况下，关键词比较次数为log2(n+1)，且期望时间复杂度为O(log2n)；
##### 注：折半查找的前提条件是需要有序表顺序存储，对于静态查找表，一次排序后不再变化，折半查找能得到不错的效率。但对于需要频繁执行插入或删除操作的数据集来说，维护有序的排序会带来不小的工作量，那就不建议使用。——《大话数据结构》
 
# Source
```go
func binarySearch(array []int /*有序数组*/ , key int /*查找元素*/) int {
	low := 0
	high := len(array)
	for low <= high {
		middle := (low + high) / 2
		if key == array[middle] {
			return middle
		} else if key < array[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}
```
### 递归方式实现
```go
func binarySearchRecursion(array []int, key int, beginIndex int, endIndex int) int {
	midIndex := (beginIndex + endIndex) / 2
	if key < array[beginIndex] || key > array[endIndex] || beginIndex > endIndex {
		return -1
	}
	if key < array[midIndex] {
		return binarySearchRecursion(array, key, beginIndex, midIndex-1)
	} else if key > array[midIndex] {
		return binarySearchRecursion(array, key, midIndex+1, endIndex)
	} else {
		return midIndex
	}
}
```

# Test
```go
func Test_binarySearch(t *testing.T) {
	array := []int{1, 3, 5, 7, 8, 9}
	key := 5
	fmt.Println(binarySearch(array, key))
	fmt.Println(binarySearchRecursion(array, key, 0, len(array)-1))
}
```