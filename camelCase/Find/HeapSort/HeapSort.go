package HeapSort

import "fmt"

func HeapSort() {
	var arr = []int32{53, 17, 78, 9, 45, 65, 87, 32}
	//var arr = []int32{7, 1, 2, 6, 5, 3, 8, 9, 10}

	/**
	堆排序
	*/
	for i := len(arr) - 1; i >= 0; i-- {
		buildMaxHeap(arr, i)
		swapElement(arr, 0, i)
	}
	fmt.Println(arr)
}

func buildMaxHeap(arr []int32, len int) {
	for lastNodeIndex := len/2 - 1; lastNodeIndex >= 0; lastNodeIndex-- { // 多少个分支结点就循环几次
		heapifyDown(arr, lastNodeIndex, len)
	}
}

func heapifyDown(arr []int32, lastNodeIndex, len int) {
	var (
		leftChildIndex  = 2*lastNodeIndex + 1
		rightChildIndex = leftChildIndex + 1
	)

	if rightChildIndex < len {
		if arr[leftChildIndex] < arr[rightChildIndex] && arr[lastNodeIndex] < arr[rightChildIndex] {
			swapElement(arr, lastNodeIndex, rightChildIndex)
		}
		heapifyDown(arr, rightChildIndex, len)
		return
	}

	if leftChildIndex < len {
		if leftChildIndex <= len && arr[lastNodeIndex] < arr[leftChildIndex] {
			swapElement(arr, lastNodeIndex, leftChildIndex)
		}
		heapifyDown(arr, rightChildIndex, len)
		return
	}
}

func swapElement[T comparable](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
