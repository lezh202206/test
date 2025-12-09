package FastSort

import "fmt"

func FastSort() {
	var arr = []int32{49, 38, 65, 97, 76, 13, 27, 49}
	Sort(arr)
	fmt.Println(arr)
}
func Sort(arr []int32) {
	// 核心是 划分
	// 用一个 基准 划分两个分区
	// 用两个头尾指针对比 基准
	// 若左指针 大于基准 则需要换位置 然后 ++
	// 若右指针 小于基准 则需要换位置 然后 --
	QuictSort(arr, 0, len(arr)-1)
}

func QuictSort(nums []int32, left, right int) {
	if left < right { // 一直递归 直到左指针不再小于右指针
		var pivotIndex = partition(nums, left, right) // 先进行一轮排序 确定基准
		QuictSort(nums, left, pivotIndex-1)           // 左分区
		QuictSort(nums, pivotIndex+1, right)          // 右分区
	}
}

// 分区
func partition(nums []int32, left, right int) int {
	var pivot = nums[left]
	for left < right {
		for left < right && nums[right] >= pivot { // 找到小于基准的值
			right--
		}
		nums[left] = nums[right] // 调换位置

		for left < right && nums[left] <= pivot { // 找到大于基准的值
			left++
		}
		nums[right] = nums[left] // 调换位置
	}
	nums[left] = pivot

	return left
}
