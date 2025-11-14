package StackList

import "fmt"

type stackNodeList struct {
	data    []int
	maxSize int
}

func (s *stackNodeList) validSize() bool {
	return s.Size() > s.maxSize
}

// Push 入栈
func (s *stackNodeList) Push(v int) {
	if s.validSize() {
		return
	}
	s.data = append(s.data, v)
}

// Pop 出栈
func (s *stackNodeList) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, true
}

// Top 获取栈顶元素
func (s *stackNodeList) Top() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	return s.data[len(s.data)-1], true
}

// Empty 判断是否为空
func (s *stackNodeList) Empty() bool {
	return len(s.data) == 0
}

// Size 返回栈大小
func (s *stackNodeList) Size() int {
	return len(s.data)
}

func Do() {
	stackList := &stackNodeList{maxSize: 4}

	stackList.Push(1)
	stackList.Push(2)
	stackList.Push(3)

	fmt.Println(stackList.Top()) // 3 true
	fmt.Println(stackList.Pop()) // 3 true
	fmt.Println(stackList.Pop()) // 2 true

	fmt.Println(stackList.Empty()) // false
	fmt.Println(stackList.Size())  // 1
}
