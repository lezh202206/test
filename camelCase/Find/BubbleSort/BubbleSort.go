package BubbleSort

import "fmt"

func BubbleSort() {
	var arr = []int32{49, 38, 65, 97, 76, 13, 27, 49}
	Sort(arr) // 简单版
	fmt.Println(arr)
}

func Sort(arr []int32) {
	for i := 0; i < len(arr)-1; i++ { // 这里决定几趟对比
		flag := true
		for j := 0; j < len(arr)-1; j++ { // 这里决定对比几次
			if arr[j] > arr[j+1] {
				swap(&arr, j, j+1)
				flag = false
			}
		}
		if flag { // flag 始终为 true 的话 证明后面排序都是对的
			return
		}
	}
}

func swap(arr *[]int32, i, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}
