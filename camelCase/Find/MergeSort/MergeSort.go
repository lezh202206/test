package MergeSort

import "fmt"

func MergeSort() {
	var arr = []int32{16, 24, 66, 45, 21, 24, 33}
	Sort(arr, 0, len(arr)-1)
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
	这里的坐标都是给临时数组的
	临时数组是由原数组 复制过来的
	区别在于 始终从 0 开始 这是重点 ！！！！

	tempLeftEnd    = mid - leftStart
	举例
	leftStart = 3
	rightEnd  = 7
	mid = (3 + 7) / 2 = 5
	原数组下标 → tempArr 下标
	3 → 0
	4 → 1
	5 → 2
	6 → 3
	7 → 4

	*/
	var (
		mid            = (rightEnd + leftStart) / 2
		tempLeftStart  int               // 临时数组永远都是从 0 开始
		tempLeftEnd    = mid - leftStart // 因为 tempArr 的下标从 0 开始，而 mid 是原数组的下标，必须减去 leftStart 才能映射到 tempArr 的正常索引
		tempRightStart = tempLeftEnd + 1
		tempRightEnd   = len(tempArr) - 1
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
