package DoubleLinkedList

import "fmt"

// Node 表示双向链表的节点
type Node struct {
	data int
	prev *Node
	next *Node
}

// DoublyLinkedList 表示双向链表结构
type DoublyLinkedList struct {
	head *Node // 头节点
	tail *Node // 尾节点
	size int   // 链表长度
}

// Append 在链表尾部添加节点
func (l *DoublyLinkedList) Append(val int) {
	newNode := &Node{data: val}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
	l.size++
}

// Prepend 在链表头部添加节点
func (l *DoublyLinkedList) Prepend(val int) {
	newNode := &Node{data: val}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
	l.size++
}

// InsertAt 在指定位置插入节点 (0-based)
func (l *DoublyLinkedList) InsertAt(index int, val int) {
	if index < 0 || index > l.size {
		panic("index out of range")
	}

	if index == 0 {
		l.Prepend(val)
		return
	}
	if index == l.size {
		l.Append(val)
		return
	}

	newNode := &Node{data: val}
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	prevNode := cur.prev
	prevNode.next = newNode
	newNode.prev = prevNode
	newNode.next = cur
	cur.prev = newNode
	l.size++
}

// Remove 删除指定位置节点
func (l *DoublyLinkedList) Remove(index int) {
	if index < 0 || index >= l.size {
		panic("index out of range")
	}

	if index == 0 {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil // 删除后为空表
		}
		l.size--
		return
	}

	if index == l.size-1 {
		l.tail = l.tail.prev
		l.tail.next = nil
		l.size--
		return
	}

	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	l.size--
}

// Find 查找指定值的索引
func (l *DoublyLinkedList) Find(val int) int {
	cur := l.head
	i := 0
	for cur != nil {
		if cur.data == val {
			return i
		}
		cur = cur.next
		i++
	}
	return -1
}

// PrintForward 从头到尾打印
func (l *DoublyLinkedList) PrintForward() {
	cur := l.head
	for cur != nil {
		fmt.Printf("%d -> ", cur.data)
		cur = cur.next
	}
	fmt.Println("nil")
}

// PrintBackward 从尾到头打印
func (l *DoublyLinkedList) PrintBackward() {
	cur := l.tail
	for cur != nil {
		fmt.Printf("%d -> ", cur.data)
		cur = cur.prev
	}
	fmt.Println("nil")
}

func Do() {
	list := &DoublyLinkedList{}

	list.Append(1)
	list.Append(3)
	list.Append(5)
	list.PrintForward()  // 1 -> 3 -> 5 -> nil
	list.PrintBackward() // 5 -> 3 -> 1 -> nil

	list.InsertAt(1, 2)
	list.InsertAt(3, 4)
	list.PrintForward()  // 1 -> 2 -> 3 -> 4 -> 5 -> nil
	list.PrintBackward() // 5 -> 4 -> 3 -> 2 -> 1 -> nil

	list.Remove(2)
	list.PrintForward() // 1 -> 2 -> 4 -> 5 -> nil

	fmt.Println("Find(4):", list.Find(4)) // 2
	fmt.Println("Size:", list.size)
}
