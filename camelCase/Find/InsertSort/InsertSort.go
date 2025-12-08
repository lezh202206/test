package InsertSort

import (
	"fmt"
)

// InsertSort 插入排序
func InsertSort() {
	var arr = []int32{49, 38, 65, 97, 76, 13, 27, 49}
	//SortV1(arr) // 简单版
	SortV2(arr) // 折半查找
	fmt.Println(arr)
}

// SortV1 原始版
func SortV1(arr []int32) []int32 {
	for i := 0; i < len(arr)-1; i++ {
		_temp := arr[i+1]
		if _temp > arr[i] {
			continue
		}
		for j := i; j >= 0; j-- {
			if arr[j] > _temp {
				arr[j+1] = arr[j]
				arr[j] = _temp
			}
		}
	}
	return arr
}

// SortV2  折半+ 插入排序
func SortV2(arr []int32) []int32 {
	for i := 0; i < len(arr)-1; i++ {
		_temp := arr[i+1]
		if _temp > arr[i] {
			continue
		}
		var (
			left  = i + 1
			right = 0
		)
		for right <= left { // 折半查找 找到最后 left 的位置
			mid := (left + right) / 2
			if arr[mid] > _temp {
				left = mid - 1
			} else {
				right = mid + 1
			}
		}

		for j := i; j >= left+1; j-- { // 需要向后移几个
			arr[j+1] = arr[j]
		}
		arr[left+1] = _temp
	}
	return arr
}
