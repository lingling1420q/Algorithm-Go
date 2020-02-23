package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 100} //升序序列
	key := 100
	index := SearchInsert(slice, key)
	if index == -1 {
		fmt.Printf("%v 不存在元素%v\n", slice, key)
	} else {
		fmt.Printf("%v 位于%v下标为%v的位置。\n", key, slice, index)
	}
}

func SearchInsert(slice []int, key int) int {
	left, right := 0, len(slice)-1
	for left <= right {
		if slice[right] == slice[left] {
			if slice[left] == key {
				return left
			}
			return -1
		}
		mid := left + (right-left)*(key-slice[left])/(slice[right]-slice[left]) // 计算mid关键步骤
		if mid < left || mid > right {
			return -1
		}
		if slice[mid] == key {
			return mid
		}
		if slice[mid] < key {
			return mid + 1
		}
		if slice[mid] > key {
			return mid - 1
		}
	}
	return -1
}

/*
插值查找是对二分查找的优化，是有序序列的查找算法。

二分查找：mid = (left + right)/2 = left/2 + right/2 = left - left/2 + right/2

　　　　　　   = left + (right - left)/2 = left + (right - left) * (1/2)

二分查找选取中间位置，插值查找则通过查找值判定大概位于序列的哪个位置比例。

插入查找：选择下标 = left + (right - left) * (key - arr[left])/(arr[right] - arr[left])
*/
