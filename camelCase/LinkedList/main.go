package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// LinkNodeList 单链表
type LinkNodeList struct {
	head *Node
	size int
}

func main() {
	nodeList := &LinkNodeList{}
	nodeList.Append(1)
	nodeList.Append(2)
	nodeList.Append(2)
	nodeList.Append(3)
	nodeList.DelAtPosition(1)
	nodeList.Append(4)
	nodeList.Append(5)
	nodeList.Append(7)
	nodeList.InsertAtPosition(5, 6)
	nodeList.InsertAtPosition(7, 8)
	nodeList.InsertAtPosition(7, 8)
	nodeList.DelAtPosition(7)
	nodeList.Println()
}

func (l *LinkNodeList) Println() {
	cur := l.head
	num := 1
	for cur != nil { // 有链就会一直循环 直到最后一个链
		if num == l.size {
			fmt.Printf("%d", cur.data)
		} else {
			fmt.Printf("%d -> ", cur.data)
		}
		cur = cur.next // 打印下一个链
		num++
	}
}

// Append 每次在最后追加一个新链
func (l *LinkNodeList) Append(v int) {
	node := &Node{data: v, next: nil}
	// 第一次 没有数据 类似初始化了
	if l.size == 0 {
		l.head = node
	} else {
		cur := l.head
		// 如果这个链下面已经存在链的话就添加到下一个链中 以免被覆盖 应该是追加
		// 这里要循环 找到最后一个
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = node // 先把新链 填到末尾的 next 中
	}
	l.size++
}

// InsertAtPosition 不能插入队首 不能超过链路总长
func (l *LinkNodeList) InsertAtPosition(index, v int) {
	if index < 0 || index > l.size {
		fmt.Printf("数组越界")
		return
	}
	newNode := &Node{data: v, next: nil}
	// 找到目标位置的链
	// 左边的 next 要链给新的链 没有就是 链首
	// 新链的 next 要填入右边的链 没有就是 null
	cur := l.head
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}

	// 如果先把左链的 next 填充新链 就找不到右链了 ！！！！
	newNode.next = cur.next // 要先把右链先填充到 新链的 next 中
	cur.next = newNode      // 把新链 放到 左链的 next 中
	l.size++
}

// DelAtPosition 删除指定节点
func (l *LinkNodeList) DelAtPosition(index int) {
	if index < 0 || index > l.size-1 {
		fmt.Printf("数组越界")
		return
	}
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	temp := cur.next.next // 将下级节点给临时变量 也就是右边的链
	cur.next = temp       // 把右链直接给左链的 next
	l.size--              // 数量减 1
}
