package Tbtree

import "fmt"

// Node 定义
type Node struct {
	Val     int
	Left    *Node
	Right   *Node
	LThread bool // true 表示 Left 指向前驱（线索），false 表示指向左子树
	RThread bool // true 表示 Right 指向后继（线索），false 表示指向右子树
}

// NewNode 快捷构造
func NewNode(v int) *Node {
	return &Node{Val: v}
}

// 线索化 可以 参照同级目录的图示
func depthTree(root *Node) {
	var per *Node
	var dfs = func(n *Node) {}

	dfs = func(n *Node) {
		if n == nil {
			return
		}
		// 递归到最左的节点
		if !n.LThread {
			dfs(n.Left)
		}

		// 要加判断
		// 因为上面的递归结束 就会执行到这 左右孩子都可能存在 所以不能线索化
		// 左指针线索化
		if n.Left == nil {
			n.LThread = true
			n.Left = per // 上一个节点就是他的 前驱 最左边的没有前驱 所以是 nil
		}

		if per != nil && per.Right == nil {
			per.RThread = true
			per.Right = n
		}
		per = n //当前节点已经完成 将 per 指向下一个节点也就是当前节点

		// 递归到最右的节点
		if !n.RThread {
			dfs(n.Right)
		}
	}
	dfs(root)
}

// buildSampleTree 返回一个示例二叉树，结构：
//
//	    4
//	   / \
//	  2   6
//	 / \ / \
//	1  3 5  7
func buildSampleTree() *Node {
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)
	n5 := NewNode(5)
	n6 := NewNode(6)
	n7 := NewNode(7)

	n4.Left, n4.Right = n2, n6
	n2.Left, n2.Right = n1, n3
	n6.Left, n6.Right = n5, n7
	return n4
}

// leftmost 返回以 node 为根的子树中的最左（中序开始点）
func leftmost(node *Node) *Node {
	if node == nil {
		return nil
	}
	p := node
	// 当 Left 是子树（LThread == false）则继续向左
	for p != nil && !p.LThread && p.Left != nil {
		p = p.Left
	}
	return p
}

// InOrderTraversal 使用线索进行中序遍历（无栈、无递归）
func InOrderTraversal(root *Node, visit func(v int)) {
	// 从根的最左节点开始
	cur := leftmost(root)
	for cur != nil {
		visit(cur.Val)
		// 如果右是线索，则直接跳到后继
		if cur.RThread {
			cur = cur.Right
		} else {
			// 否则进入右子树的最左节点
			cur = leftmost(cur.Right)
		}
	}
}

// 示例：打印中序
func Do() {
	root := buildSampleTree()
	// 未线索化前，如果你想，也可以验证普通中序遍历（这里略）
	depthTree(root)

	fmt.Print("中序遍历（线索）: ")
	InOrderTraversal(root, func(v int) {
		fmt.Printf("%d ", v)
	})
	fmt.Println()
}
