package MergeSort

import "fmt"

func MergeSort() {
	var arr = []int32{16, 24, 37, 45, 21, 24, 33}
	Sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func Sort(arr []int32, leftStart, rightEnd int) {
	if leftStart >= rightEnd { // 重合了 只有一个元素 不需要循环 left > right 是异常情况无需处理
		return
	}

	mid := (leftStart + rightEnd) / 2
	Merge(arr, leftStart, rightEnd)
	Sort(arr, leftStart, mid)  // 左边数组一直递归下去
	Sort(arr, mid+1, rightEnd) // 右边数组一直递归下去

}

func Merge(sourceArray []int32, leftStart, rightEnd int) {
	mid := (leftStart + rightEnd) / 2

	// copy 当前区间的数据
	tempArray := copyBack(sourceArray, leftStart, rightEnd)

	// temp 内部的对应下标要从 0 开始计算
	left := 0                    // 左数组下标起点
	right := mid - leftStart + 1 // 右数组下标起点
	k := leftStart               // k 表示添加到原数组的下标

	leftEnd := mid - leftStart           // 左数组下标终点
	rightEndTemp := rightEnd - leftStart // 右数组下标终点

	// 合并左右
	for left <= leftEnd && right <= rightEndTemp { // 只要两个左右指针都没到各自的终点就一直循环
		if tempArray[left] <= tempArray[right] {
			sourceArray[k] = tempArray[left]
			left++
		} else {
			sourceArray[k] = tempArray[right]
			right++
		}
		k++
	}

	// 剩余左边
	for left <= leftEnd { // 右边的数组已经全部赋值完 把左边剩下的还原到原数组
		sourceArray[k] = tempArray[left]
		left++
		k++
	}

	// 剩余右边
	for right <= rightEndTemp { // 左边的数组已经全部赋值完 把右边剩下的还原到原数组
		sourceArray[k] = tempArray[right]
		right++
		k++
	}
}

func copyBack(sourceArray []int32, leftPointer, rightPointer int) []int32 {
	tempArray := make([]int32, rightPointer-leftPointer+1)
	for i := leftPointer; i <= rightPointer; i++ { // 只需要复制 对应容量的数据 不需要全部
		tempArray[i-leftPointer] = sourceArray[i]
	}
	return tempArray
}
