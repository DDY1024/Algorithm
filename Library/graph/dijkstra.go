package main

import (
	"container/heap"
)

// 贪心策略
// 		1. 每次选择距离【原点】距离最近的点进行扩展
//      2. 更新计算其邻接节点的距离
//
// 多维情况下的 dijstra 算法  https://blog.csdn.net/u013081425/article/details/26020401
// 		1. 如上所示在每次选择【扩展节点】时，综合考虑了【距离】和【花费】两个条件
//
// 注意事项
//      1. dijstra 仅适应于【边权非负】和【边权非正】情况下的最短路径问题
//

const (
	inf = 0x3f3f3f3f3f3f3f3f
)

// 1. 最短距离
// 2. 最长距离
func dijkstra(s, n int, g [][]int) []int {
	var (
		dis = make([]int, n)
		vis = make([]bool, n)
	)

	for i := 0; i < n; i++ {
		dis[i] = inf
	}
	dis[s] = 0

	// n 次迭代计算
	for i := 0; i < n; i++ {
		u := -1
		for j := 0; j < n; j++ {
			if vis[j] {
				continue
			}
			if u == -1 || dis[u] > dis[j] {
				u = j
			}
		}

		// 非连通图，停止计算
		if u == -1 {
			break
		}

		vis[u] = true
		for j := 0; j < n; j++ {
			if vis[j] {
				continue
			}
			if dis[j] > dis[u]+g[u][j] {
				dis[j] = dis[u] + g[u][j]
			}
		}
	}
	return dis
}

//  2. dijkstra + heap 实现
//     算法复杂度 O(m * logm)  --> 更适合稀疏图的场景
type Edge struct {
	v, w int
}

type Node struct {
	u, d int
}

type PQ []*Node

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].d < pq[j].d
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	Node := x.(*Node)
	*pq = append(*pq, Node)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	Node := old[n-1]
	*pq = old[0 : n-1]
	return Node
}

func (pq *PQ) Top() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	return (*pq)[0]
}

// 采用【邻接表】存储
func dijkstraHeap(s, n int, g [][]Edge) []int {
	var (
		dis = make([]int, n)
		vis = make([]bool, n)
	)

	for i := 0; i < n; i++ {
		dis[i] = inf
	}
	dis[s] = 0

	pq := make(PQ, 0)
	heap.Push(&pq, &Node{s, 0})
	for pq.Len() > 0 {
		nd := heap.Pop(&pq).(*Node)
		if vis[nd.u] { // 判重，避免重复计算
			continue
		}

		u := nd.u
		vis[u] = true
		for _, e := range g[u] {
			v, w := e.v, e.w
			if dis[v] > dis[u]+w {
				dis[v] = dis[u] + w
				heap.Push(&pq, &Node{v, dis[v]})
			}
		}
	}
	return dis
}

// 次短路径
// 		1. dist[u][0]：到达顶点 u 的最短路径
//      2. dist[u][1]：到达顶点 u 的次短路径
