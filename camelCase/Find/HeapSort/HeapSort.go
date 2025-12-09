package HeapSort

import "fmt"

func HeapSort() {
	var arr = []int32{53, 17, 78, 9, 45, 65, 87, 32}
	//var arr = []int32{7, 1, 2, 6, 5, 3, 8, 9, 10}

	for i := len(arr) - 1; i >= 0; i-- { // 递减是因为 每次下沉后最后一个是有序的
		buildMaxHeap(arr, i+1) //构建 大堆顶 也就是把最大的值放到堆顶
		swapElement(arr, 0, i) // 放到元素末尾 后续遍历不会影响他
	}
	fmt.Println(arr)
}

func buildMaxHeap(arr []int32, len int) {
	for lastNodeIndex := len/2 - 1; lastNodeIndex >= 0; lastNodeIndex-- { // 多少个分支结点就循环几次
		heapifyDown(arr, lastNodeIndex, len) // 堆 下沉 替换
	}
}

func heapifyDown(arr []int32, lastNodeIndex, len int) {
	var (
		leftChildIndex  = 2*lastNodeIndex + 1 // 二叉树中的左孩子 看图
		rightChildIndex = leftChildIndex + 1  // 二叉树中的右孩子 看图
	)
	// 左孩子不存在 说明 也没有右孩子 (从索引下标角度)
	if leftChildIndex >= len {
		return
	}

	largerChild := leftChildIndex
	if rightChildIndex < len && arr[rightChildIndex] > arr[leftChildIndex] { // 首先确定存在右孩子 再确定右孩子是不是比左孩子大 如果是就把右孩子上移
		largerChild = rightChildIndex
	}

	// 如果左右孩子都比当前节点小 则无需调整
	if arr[lastNodeIndex] >= arr[largerChild] {
		return
	}

	swapElement(arr, lastNodeIndex, largerChild) // 交换最大节点
	heapifyDown(arr, largerChild, len)           // 再看最大节点的这边是否受影响 有影响就往下处理
}

func swapElement[T comparable](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
