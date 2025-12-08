package LinearTable

import "fmt"

var ARR = []int32{2, 3, 6, 8, 11, 14, 16, 20, 25, 27, 30, 33}

func BinarySearch() int {
	const target int32 = 11
	var (
		low  = 0
		mid  = 0
		high = len(ARR) - 1
	)

	for low <= high {
		mid = (low + high) / 2
		if ARR[mid] == target {
			fmt.Println("下标为:", mid)
			return mid
		}
		if ARR[low] == target {
			fmt.Println("下标为:", low)
			return low
		}
		if ARR[high] == target {
			fmt.Println("下标为:", high)
			return high
		}
		if ARR[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("查找失败～～～")
	return -1
}
