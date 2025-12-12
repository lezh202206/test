package main

import (
	"test/camelCase/Find/CountSort"
)

func main() {
	//SingleLinkedList.Do() // 单链表

	//DoubleLinkedList.Do() // 双链表

	//StackList.Do()     // 栈

	//CircularQueue.Do() // 循环队列

	//SlidingWindow.Do() // 滑动窗口

	//KMP.Do()    // 子串第一次出现的位置

	//Tbtree.Do() // 线索二叉树

	//Graph.Prim() // 图 Prim算法 最小生成数

	//LinearTable.BinarySearch() // 线性表 (顺序表)二分法查询

	//InsertSort.InsertSort() // 插入排序 // 两个指针依次对比

	//BubbleSort.BubbleSort() // 冒泡排序 // 两个循环对比 依次对比

	//FastSort.FastSort() // 快速排序 通过基准确定 位置 再分为左右分区

	//HeapSort.HeapSort() // 堆排序 二叉树的视角 大堆顶

	//MergeSort.MergeSort() // 归并排序 切割为多个数组 再合并

	//RadixSort.RadixSort() // 基数排序 // 适用于多维度排序

	// O(n+k) 计数排序 n 长度 k 取值范围 相同的算一个
	CountSort.CountSort() // 计数排序 空间换时间 如果每个元素取值差异都较大 不太适用
}
