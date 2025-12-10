package MergeSort

import "fmt"

func MergeSort() {
	var arr = []int32{16, 24, 66, 45, 21, 24, 33}
	Sort(arr, 1, len(arr)-1)
	fmt.Println(arr)
}

func Sort(arr []int32, left, right int) {
	if left >= right {
		return
	}
	Merge(arr, left, right)
	mid := (right + left) / 2
	Sort(arr, left, mid)
	Sort(arr, mid+1, right)
}

func Merge(sourceArr []int32, leftStart, rightEnd int) {
	tempArr := copyBack(sourceArr, leftStart, rightEnd)
	/**
	定义 6 个下标 分别为
	左数组起始 leftStart
	左数组结束 通过中枢确定
	中枢下标 通过这个中枢确定左右两边数组 （中枢下标归为左边数组）
	右数组起始
	右数组结束 rightEnd
	原数组下标 正常来说是 0 这里应该保持和左数组起始下标保持同步
	*/
	var (
		mid           = (rightEnd + leftStart) / 2
		tempLeftStart int
		tempLeftEnd   = mid - leftStart

		tempRightStart = tempLeftEnd + 1
		tempRightEnd   = rightEnd - leftStart
		sourceArrIndex = leftStart
	)
	/**
	只要 左右数组下标不越界就已经循环
	也就是说 起始位置不能大于结束位置
	*/
	for tempLeftStart <= tempLeftEnd && tempRightStart <= tempRightEnd {
		if tempArr[tempLeftStart] <= tempArr[tempRightStart] {
			sourceArr[sourceArrIndex] = tempArr[tempLeftStart]
			tempLeftStart++
		} else {
			sourceArr[sourceArrIndex] = tempArr[tempRightStart]
			tempRightStart++
		}
		sourceArrIndex++
	}

	/**
	上面的循环可以 确保有一边数组是全部遍历结束了
	接下来只有把剩下那边的“尾部”补到原数组就好了

	那就是看指针 越没越界 就知道有没有处理完
	*/
	for tempLeftStart <= tempLeftEnd {
		sourceArr[sourceArrIndex] = tempArr[tempLeftStart]
		tempLeftStart++
		sourceArrIndex++
	}

	for tempRightStart <= tempRightEnd {
		sourceArr[sourceArrIndex] = tempArr[tempRightStart]
		tempRightStart++
		sourceArrIndex++
	}
}

func copyBack(sourceArr []int32, leftStart, rightEnd int) []int32 {
	/**
	rightEnd-leftStart 始终都是适量的容量不需要 copy 全部, +1 是因为下标从 0 开始的
	i-leftStart 下标都要从 0 开始 要不然地柜到最后会越界
	*/
	var tempArr = make([]int32, rightEnd-leftStart+1)
	for i := leftStart; i <= rightEnd; i++ {
		tempArr[i-leftStart] = sourceArr[i]
	}
	return tempArr
}
