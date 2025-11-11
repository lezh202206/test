package main

import "fmt"

func move(arr []int32, pos int, val int32) []int32 {
	if pos > len(arr) || pos < 0 {
		fmt.Println("Invalid position")
		return nil
	}
	arr = append(arr, 0)
	// 从队尾开始复制 (一个一个挪)
	// 挪到目标位置不用循环了 也就是说 i 值只需要等于 总容量 - 目标位置 = 差值 （例如 总容量 6 目标 3 只需要循环 3次）
	// 如果差值 小于等于 0 说明 到队尾了不需要挪直接赋值就好
	for i := len(arr); i > pos; i-- {
		arr[i-1] = arr[i-2]
	}
	arr[pos] = val
	return arr
}

func moveDel(arr []int32, pos int) []int32 {
	if pos < 0 || pos >= len(arr) {
		fmt.Println("数组越界")
		return nil
	}
	// 起点为被删除的 索引
	// 依次替换
	// 到达最后一个索引后 跳出循环
	// 最后一个值赋值为空
	for i := pos; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = 0
	return arr
}

// 顺序表
func main() {
	var A = make([]int32, 10)
	A = []int32{1, 2, 3, 5, 4}
	move(A, 3, 3)
	moveDel(A, 5)
}
