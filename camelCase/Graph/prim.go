package Graph

import (
	"container/heap"
	"fmt"
)

// ---------- 图结构 ----------
type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	Adj [][]Edge
}

func NewGraph(n int) *Graph {
	return &Graph{Adj: make([][]Edge, n)}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.Adj[u] = append(g.Adj[u], Edge{To: v, Weight: w})
	g.Adj[v] = append(g.Adj[v], Edge{To: u, Weight: w})
}

// ---------- 小顶堆 ----------
type Item struct {
	Node   int
	Weight int
}

type MinHeap []Item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// ---------- Prim Algorithm ----------
func PrimMST(g *Graph) (mstWeight int, mstEdges [][2]int) {
	n := len(g.Adj)
	visited := make([]bool, n)

	h := &MinHeap{}
	heap.Init(h)

	// 从 0 节点开始
	visited[0] = true
	for _, e := range g.Adj[0] {
		heap.Push(h, Item{Node: e.To, Weight: e.Weight})
	}

	for h.Len() > 0 {
		item := heap.Pop(h).(Item)
		v := item.Node

		if visited[v] {
			continue
		}
		visited[v] = true
		mstWeight += item.Weight

		// 记录最小生成树的边（Prim 需要额外记录）
		// 这里需要找到连接到 v 的那条边
		for _, e := range g.Adj[v] {
			if visited[e.To] && e.Weight == item.Weight {
				mstEdges = append(mstEdges, [2]int{v, e.To})
				break
			}
		}

		// 将新的边加入堆
		for _, e := range g.Adj[v] {
			if !visited[e.To] {
				heap.Push(h, Item{Node: e.To, Weight: e.Weight})
			}
		}
	}

	return
}

// ---------- 示例 ----------
func Prim() {
	// 无向图(双向图) 参考图片
	g := NewGraph(5)
	g.AddEdge(0, 1, 2) // 节点 0 -> 节点 1, 权重为 2  无向图也就是说 节点 1 -> 节点 0, 权重也是 2
	g.AddEdge(0, 3, 6) // 节点 0 -> 节点 3, 权重为 6  无向图也就是说 节点 1 -> 节点 0, 权重也是 6
	g.AddEdge(1, 2, 3) // 以此类推
	g.AddEdge(1, 3, 8)
	g.AddEdge(1, 4, 5)
	g.AddEdge(2, 4, 7)
	g.AddEdge(3, 4, 9)

	totalWeight, mstEdges := PrimMST(g)

	fmt.Println("最小生成树权重 =", totalWeight)
	fmt.Println("最小生成树边：")
	for _, e := range mstEdges {
		fmt.Println(e[0], "<->", e[1])
	}
}
