package MergeSort

import "fmt"

func MergeSort() {
	//var arr = []int32{49, 38, 65, 97, 76, 13, 27, 49}
	var arr = []int32{16, 24, 37, 45, 21, 24, 33}
	Sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func Sort(arr []int32, leftStart, rightEnd int) {
	if leftStart < rightEnd {
		Merge(arr, leftStart, rightEnd)
		Sort(arr, leftStart, len(arr)/2)  // 左排序
		Sort(arr, len(arr)/2+1, len(arr)) // 右排序
	}
}

func Merge(sourceArray []int32, leftStart, rightEnd int) {
	var (
		tempArray = copyBack(sourceArray, leftStart, rightEnd)
		mid       = len(sourceArray) / 2 // 中枢 划分左右数组 中枢包含在左边数组中
		left      = leftStart            // 左边的起始位置
		right     = mid + 1              // 右边数组起始位置 中枢下标+1
		k         int                    // 每次改sourceArray 都会+1 最终 会等于 len
	)

	for k = left; left <= mid && right < len(sourceArray); k++ {
		if tempArray[left] <= tempArray[right] {
			sourceArray[k] = tempArray[left]
			left++
		} else {
			sourceArray[k] = tempArray[right]
			right++
		}
	}

	for left < mid+1 { // mid 是左边数组的边界 +1是因为下班为 0 开始的
		sourceArray[k] = tempArray[left]
		k++
		left++
	}

	for right < len(sourceArray) { // len(sourceArray) 是右边数组的边界
		sourceArray[k] = tempArray[right]
		k++
		right++
	}
}

func copyBack(sourceArray []int32, leftPointer, rightPointer int) []int32 {
	// 一个数组拆为分为左右两边
	var (
		tempArray = make([]int32, len(sourceArray))
	)
	for index, v := range sourceArray[leftPointer : rightPointer+1] {
		tempArray[index] = v
	}
	return tempArray
}
