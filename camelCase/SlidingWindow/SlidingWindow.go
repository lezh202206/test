package SlidingWindow

import "fmt"

func Do() {
	var A = []int32{1, 2, 1, 2, 1, 2, 1, 1, 1, 2, 3}
	var B = []int32{1, 2}

	//fmt.Println(SlidingWindow(A, B))
	fmt.Println(KMP(A, B))
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

// KMP 查找子序列 B 在 A 中的起止下标（左闭右开）
// 找不到返回 (0,0) 和 false
//func KMP(A, B []int32) (int, int, bool) {
//	n, m := len(A), len(B)
//	if m == 0 || n < m {
//		return 0, 0, false
//	}
//
//	// 1. 构建 LPS（最长前缀后缀表）
//	// 对比值     是否相等   左右指针是否后挪                        当前左右指针值       改后左右指针值       next 数组    变动后next 数组
//	// 1 2(左右)  不相等      右指针往后挪                          left：0 right：1     left：0 right：2      [0]        [0]
//	// 1 1        相等      左右指针往后挪                         left：0 right：2     left：1 right：3      [0]        [0,0,2]
//	// 2 2        相等      左右指针往后挪                         left：1 right：3     left：2 right：4      [0,0,2]    [0,0,1,2]
//	// 1 4       不相等      左指针取 next 下标 left 中-1个        left：2 right：4     left：2 right：5      [0,0,1,2]  [0,0,1,2,0]
//	lps := make([]int, m)
//	for right, left := 1, 0; right < m; {
//		if B[right] == B[left] {
//			left++
//			lps[right] = left
//			right++
//		} else { // n n-1 前后不一致
//			if left != 0 {
//				left = lps[left-1]
//				//left = 0
//			} else {
//				right++ // 前后不一致 右指针 往后挪
//			}
//		}
//	}
//
//	// 2. 匹配主串
//	for i, j := 0, 0; i < n; {
//		if A[i] == B[j] {
//			i++
//			j++
//		}
//
//		// 匹配完成
//		if j == m {
//			start := i - j
//			return start, start + m, true
//		}
//
//		if i < n && A[i] != B[j] {
//			if j != 0 { // 说明前面没到 最长前缀的头部
//				j = lps[j-1]
//			} else {
//				i++ // 到这说明没重复的了 所以 A 数组要往后找了 所以 i++
//			}
//		}
//	}
//
//	return 0, 0, false
//}

func KMP(A, B []int32) (int, int) {
	fmt.Println(GetALLIndex(A, B, LPS(B)))
	return GetIndex(A, B, LPS(B))
}

func LPS(data []int32) []int {
	if len(data) == 0 {
		return []int{}
	}

	// LPS 原理为: 寻找最长相同的 前后缀 长度
	// 也就是说 举个例子
	// [1,2,1,2,3,1] LPS 数组为 [0,0,1,2,0,1]
	// 对应 index 为 0 时 (1)默认为 0
	// 对应 index 为 1 时 (1,2) 为 0
	// 对应 index 为 2 时 对比 (1,2,1) => ([1],2,[1]) 前后缀长度为 1
	// 对应 index 为 3 时 对比 (1,2,1,2) => ([1,2],[1,2]) 前后缀长度为 2
	// 对应 index 为 4 时 对比 (1,2,1,2,3) => (1,2,1,2,3) 前后缀长度为 0 (因为前后缀的组合没有相同的 也就是必须包含 3 除非 前面有 1，2，3 此时 ([1,2,3],[1,2,3])) 长度为 3)
	// 对应 index 为 5 时 对比 (1,2,1,2,3,1) => ([1],2,1,2,3,[1]) 前后缀长度为 1
	var lps = make([]int, len(data))
	for left, right := 0, 1; right < len(data); {
		if data[left] == data[right] {
			left++
			lps[right] = left
			right++
		} else {
			if left != 0 {
				left = lps[left-1]
			} else {
				right++
			}
		}
	}

	return lps
}

func GetIndex(A, B []int32, LPS []int) (int, int) {
	for AIndex, BIndex := 0, 0; AIndex < len(A); {
		if A[AIndex] == B[BIndex] {
			AIndex++
			BIndex++
		}

		if BIndex == len(B) {
			start := AIndex - BIndex
			return start, start + len(B)
		}

		if AIndex < len(A) && A[AIndex] != B[BIndex] {
			if BIndex != 0 {
				// 精髓在这
				// LPS 存的可以跳过的位数 或 可以理解为前面最长前缀的 index
				// 重复对比 直到 BIndex没有和 A[AIndex] 当前匹配上的为止
				// 没有就跳过 执行下一个
				// 好处就是 B 数组重复的越多 被跳过的越多
				BIndex = LPS[BIndex-1]
			} else {
				AIndex++
			}
		}
	}
	return 0, 0
}

type allIndex struct {
	left, right int
}

func GetALLIndex(A, B []int32, LPS []int) []allIndex {
	var _allIndex = make([]allIndex, 0)
	for AIndex, BIndex := 0, 0; AIndex < len(A); {
		if A[AIndex] == B[BIndex] {
			AIndex++
			BIndex++
		}

		if BIndex == len(B) {
			start := AIndex - BIndex
			_allIndex = append(_allIndex, allIndex{start, start + len(B) - 1}) // 把结果都加进去
			BIndex = 0                                                         // 循环匹配
			continue
		}

		if AIndex < len(A) && A[AIndex] != B[BIndex] {
			if BIndex != 0 {
				BIndex = LPS[BIndex-1]
			} else {
				AIndex++
			}
		}
	}
	return _allIndex
}
