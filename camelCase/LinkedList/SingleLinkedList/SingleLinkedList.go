package SingleLinkedList

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

func (l *LinkNodeList) IncSize() {
	l.size++
}

func (l *LinkNodeList) DecSize() {
	l.size--
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

func (l *LinkNodeList) NewNode(v int) *Node {
	return &Node{data: v, next: nil}
}

func (l *LinkNodeList) isEmpty() bool {
	return l.head == nil
}

// validaIndex 验证数组越界
func (l *LinkNodeList) validaIndex(index int) bool {
	return index < 0 || index > l.size
}

// FindNodeLink 找到目标左链
func (l *LinkNodeList) FindNodeLink(index int) *Node {
	cur := l.head
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}
	return cur
}

// ================== 新增 ================== //

// AppendEnd 尾部追加一个
func (l *LinkNodeList) AppendEnd(v int) {
	if l.isEmpty() {
		l.head = l.NewNode(v)
	} else {
		//  A -> B -> C
		//  D ->
		//  遍历至最后一个链 把新链填入到 最后一个链的 next 中
		// A -> B -> C -> D
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.next = l.NewNode(v)
	}
	l.IncSize()
}

// AppendTop 头部追加一个
func (l *LinkNodeList) AppendTop(v int) {
	if l.isEmpty() {
		l.head = l.NewNode(v)
	} else {
		//  B -> C -> D
		//  A -> (next)
		//  A -> B -> C -> D
		temp := l.NewNode(v)
		temp.next = l.head
		l.head = temp
	}
	l.IncSize()
}

// AddLocationNode 指定位置加一个
func (l *LinkNodeList) AddLocationNode(index int, v int) {
	if l.validaIndex(index) {
		fmt.Println("数组越界")
		return
	}
	if index == 0 {
		l.AppendTop(v)
	} else if index == l.size {
		l.AppendEnd(v)
	} else {
		// 找到最终插入 位置
		cur := l.FindNodeLink(index) // 左链

		newNode := l.NewNode(v)
		// 如果先填入左链 会导致右链全部丢失
		newNode.next = cur.next // 先把右链 填入到 新链的 next 中
		cur.next = newNode      // 再把新链 填入到 左链的 next 中
		l.IncSize()
	}
}

// ================== 删除 ================== //

// DelNodeTop 删除第一个链
func (l *LinkNodeList) DelNodeTop() {
	if l.isEmpty() {
		return
	}
	if l.size == 1 {
		l.head = nil
	} else {
		l.head = l.head.next
	}
	l.DecSize()
}

// DelNodeEnd 删除最后一个链
func (l *LinkNodeList) DelNodeEnd() {
	if l.isEmpty() {
		return
	}
	cur := l.head
	for i := 0; i <= l.size-2; i++ {
		// A -> B -> C
		// A -> B
		if l.size-2 == i || l.size == 0 { // 找到倒数第二个链 置为空  或者 总链就剩下最后一个
			cur.next = nil
			continue
		}
		cur = cur.next
	}
	l.DecSize()
}

// DelLocationNode 指定位置删一个
func (l *LinkNodeList) DelLocationNode(index int) {
	if l.validaIndex(index) {
		fmt.Println("数组越界")
		return
	}
	if index == 0 {
		l.DelNodeTop()
	} else if index == l.size {
		l.DelNodeEnd()
	} else {

		cur := l.FindNodeLink(index) // 左链
		cur.next = cur.next.next
		l.DecSize()
	}
}

func Do() {
	nodeList := &LinkNodeList{}
	nodeList.AppendEnd(3)
	nodeList.AppendEnd(4)
	nodeList.AppendTop(2)
	nodeList.AddLocationNode(3, 5)
	nodeList.AddLocationNode(0, 0)
	nodeList.AddLocationNode(1, 1)
	nodeList.DelNodeTop()
	nodeList.DelNodeEnd()
	nodeList.DelLocationNode(1)

	nodeList.Println()
}
