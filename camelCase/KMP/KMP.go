package KMP

import "fmt"

func Do() (int, int) {
	var A = []int32{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 2, 1, 2, 1, 1}
	var B = []int32{1, 2, 1, 2, 1, 1}

	computeCustomLPS(B)

	fmt.Println(GetALLIndex(A, B, LPS(B))) //全部的匹配的节点
	return GetIndex(A, B, LPS(B))          // 第一个
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

func computeCustomLPS(arr []int32) []int {
	n := len(arr)
	lps := make([]int, n)

	for i := 0; i < n; i++ {
		maxLen := 0
		// 检查所有可能的前缀后缀
		for length := 1; length <= i; length++ {
			prefix := arr[:length]
			suffix := arr[i-length+1 : i+1]

			// 检查前缀和后缀是否相等
			equal := true
			for j := 0; j < length; j++ {
				if prefix[j] != suffix[j] {
					equal = false
					break
				}
			}
			if equal && length > maxLen {
				maxLen = length
			}
		}
		lps[i] = maxLen
	}
	return lps
}
