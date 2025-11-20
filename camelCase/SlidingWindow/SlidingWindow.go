package SlidingWindow

import (
	"fmt"
)

func Do() {
	var A = []int32{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 2, 1, 2, 1, 1}
	var B = []int32{1, 2, 1, 2, 1, 1}
	// [0,1,2,3,0] // ai
	// [0,1,2,3,4] // 视频
	// [0,0,0,0,4 ] // 视频pro

	// 子串 第一次出现位置
	fmt.Println(SlidingWindow(A, B)) //滑动窗口 暴力破解

	//KMP2.Do(A, B)
}

func SlidingWindow(A, B []int32) (int, int) {
	var (
		left  = 0
		right = len(B)
	)

	for {
		if right > len(A) {
			fmt.Println("未找到")
			return 0, 0
		}
		num := 0
		for i, v := range A[left:right] {
			if v != B[i] {
				left += i + 1
				right = left + len(B)
				break
			}
			num++
			if num == len(B) {
				return left, right - 1
			}
		}

	}
}
