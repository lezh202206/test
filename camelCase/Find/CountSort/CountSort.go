package CountSort

import (
	"fmt"
)

// CountSort O(n+k) 计数排序 n 长度 k 取值范围 相同的算一个
func CountSort() {
	var arr = []int32{2, 4, 3, 0, 2, 3, 0, 3}
	fmt.Println(Sort(arr))
}

func Sort(arr []int32) []int32 {
	// 1. 找最大值
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	// 2. 计数数组
	countArr := make([]int32, maxVal+1)
	for _, v := range arr { // 有多少个相同的元素
		countArr[v]++
	}

	// 核心用前一个元素的值 +当前元素的值
	// 就是 该元素第一次出现的下标位置
	for i := 1; i < len(countArr); i++ {
		countArr[i] += countArr[i-1]
	}

	resultArr := make([]int32, len(arr))
	for i := len(arr) - 1; i >= 0; i-- { // 这里直接使用映射关系就行了 没消耗一次坐标 该元素位置要 -1
		v := arr[i]
		countArr[v]--
		resultArr[countArr[v]] = v
	}
	return resultArr
}
