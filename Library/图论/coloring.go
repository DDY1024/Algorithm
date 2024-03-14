package main

// 黑白染色判断奇环

type Edge struct {
	u, v, next int
}

var (
	head  []int
	edges []Edge
	enum  int
)

func initGraph(n, m int) {
	head = make([]int, n)
	for i := 0; i < n; i++ {
		head[i] = -1
	}
	enum = 0
	edges = make([]Edge, m)
}

func addEdge(u, v int) {
	edges[enum] = Edge{u, v, head[u]}
	head[u] = enum
	enum++

	edges[enum] = Edge{v, u, head[v]}
	head[v] = enum
	enum++
}

// 0: 未染色, 1:黑色, 2:白色, 3: mask
func check(n int) bool {
	var (
		mask  = 3
		color = make([]int, n)
	)

	var dfs func(u, c int) bool
	dfs = func(u, c int) bool {
		color[u] = c
		for i := head[u]; i != -1; i = edges[i].next {
			v := edges[i].v
			if color[v] == 0 {
				return dfs(v, c^mask) // 1^3 = 2, 2^3 = 1
			}
			if color[u] == color[v] { // 存在奇环
				return true
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		// 多连通分支依次判断
		if color[i] == 0 && dfs(i, 1) {
			return true
		}
	}
	return false
}
