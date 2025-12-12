package CountSort

import (
	"fmt"
)

// CountSort 计数排序
func CountSort() {
	var arr = []int32{2, 4, 3, 0, 2, 3, 0, 3}

	Sort(arr)
	fmt.Println(arr)
}

func Sort(arr []int32) {
	var countArr = make([]int32, len(arr))
	for _, v := range arr { // 元素出现几次
		countArr[v]++
	}
	for i := 1; i < len(countArr); i++ {
		countArr[i] += countArr[i-1]
	}

	var resultArr = make([]int32, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		countArr[arr[i]] = countArr[arr[i]] - 1
		resultArr[countArr[arr[i]]] = arr[i]
	}
	fmt.Println(resultArr)
}
