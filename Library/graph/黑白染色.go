package main

// 1. 利用黑白染色法可以判定图中是否存在奇环

const (
	maxV = 110
	maxE = 10010
)

// 链式前向星图存储结构
type Edge struct {
	u, v, next int
}

var (
	head  []int
	edges []Edge
	eNum  int
)

func initGraph(n, m int) {
	head = make([]int, n)
	for i := 0; i < n; i++ {
		head[i] = -1
	}
	eNum = 0
	edges = make([]Edge, m)
}

func addEdge(u, v int) {
	edges[eNum] = Edge{u, v, head[u]}
	head[u] = eNum
	eNum++

	edges[eNum] = Edge{v, u, head[v]}
	head[v] = eNum
	eNum++
}

// 判定是否存在奇环
// 0: 未染色
// 1: 黑色
// 2: 白色
// 3: mask
func check(n int) bool {
	mask := 3
	color := make([]int, n+1)

	var dfs func(u, c int) bool
	dfs = func(u, c int) bool {
		color[u] = c
		for i := head[u]; i != -1; i = edges[i].next {
			v := edges[i].v
			if color[v] == 0 { // 未染色
				return dfs(v, c^mask)
			}
			if color[u] == color[v] {
				return true
			}
		}
		return false
	}

	for i := 1; i <= n; i++ {
		// 多个连通分支
		if color[i] == 0 && dfs(i, 1) {
			return true
		}
	}
	return false
}
