package CountSort

import "fmt"

// CountSort 计数排序
func CountSort() {
	var arr = []int32{49, 38, 65, 97, 76, 13, 27, 49}

	Sort(arr)
	fmt.Println(arr)
}

func Sort(arr []int32) {

	/**
	求最大的个数
	*/
	var countArr = make([]int32, len(arr))
	for i, v := range arr {
		countArr[i] = v
	}
}
