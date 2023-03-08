package main

import (
	"container/heap"
)

// 参考
//      https://zhuanlan.zhihu.com/p/96621396
//      https://oi-wiki.org/graph/shortest-path/
//
// 贪心思路
//      每次选择距离源点 S 距离最近的顶点进行扩展（一维）
// 扩展：https://blog.csdn.net/u013081425/article/details/26020401
//      每次选择的点同时考虑了距离和花费两个条件
//
// dijstra 算法适用于【边权非负】情况下的最短路径求解

func dijstra(n, src int, g [][]int) []int {
	inf := 0x3f3f3f3f
	dist := make([]int, n)
	vis := make([]bool, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}
	dist[src] = 0

	for i := 0; i < n; i++ {
		u := -1
		for j := 0; j < n; j++ {
			if vis[j] {
				continue
			}

			if u == -1 || dist[u] > dist[j] {
				u = j
			}
		}

		if u == -1 { // 图非连通，算法提前退出
			break
		}

		vis[u] = true
		for j := 0; j < n; j++ {
			if vis[j] {
				continue
			}

			if dist[j] > dist[u]+g[u][j] {
				dist[j] = dist[u] + g[u][j]
			}
		}
	}
	return dist
}

// 2. 堆实现的 dijkstra 算法
//      复杂度 O(m * logm)

type Edge struct {
	v, w int
}

type Node struct {
	u, d int
}

type MinPQ []*Node

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Less(i, j int) bool {
	return pq[i].d < pq[j].d
}

func (pq MinPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinPQ) Push(x interface{}) {
	Node := x.(*Node)
	*pq = append(*pq, Node)
}

func (pq *MinPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	Node := old[n-1]
	*pq = old[0 : n-1]
	return Node
}

func (pq *MinPQ) Top() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	return (*pq)[0]
}

func dijkstraHeap(n, src int, g [][]Edge) []int {
	inf := 0x3f3f3f3f
	vis := make([]bool, n)
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = inf
	}
	dis[src] = 0

	pq := make(MinPQ, 0)
	heap.Push(&pq, &Node{src, 0})
	for pq.Len() > 0 {
		nd := heap.Pop(&pq).(*Node)
		if vis[nd.u] {
			continue
		}

		vis[nd.u] = true
		for _, e := range g[nd.u] {
			v, w := e.v, e.w
			if dis[v] > dis[nd.u]+w {
				dis[v] = dis[nd.u] + w
				heap.Push(&pq, &Node{v, dis[v]})
			}
		}
	}
	return dis
}
