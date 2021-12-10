package main

import (
	"fmt"
)

const (
	Black = -1
	Gray  = 0
	White = 1
)

const (
	maxVertexNum = 110
	maxEdgeNum   = maxVertexNum * maxVertexNum
)

type Edge struct {
	u, v, next int
}

// 存储图相关信息
var head [maxVertexNum]int
var edges [maxEdgeNum]Edge
var edgeNum int

// InitGraph 初始化构建
func InitGraph() {
	edgeNum = 0
	for i := 0; i < maxVertexNum; i++ {
		head[i] = -1
	}
}

// AddEdge 有向图 或 无向图
func AddEdge(u, v int) {
	// 正向边
	edges[edgeNum] = Edge{u: u, v: v, next: head[u]}
	head[u] = edgeNum
	edgeNum++

	// 反向边
	edges[edgeNum] = Edge{u: v, v: u, next: head[v]}
	head[v] = edgeNum
	edgeNum++
}

// IsExistOddCycle
func IsExistOddCycle(n int) bool {

	color := make([]int, n+1)
	var dfs func(u, c int) bool
	dfs = func(u, c int) bool {
		color[u] = c
		for i := head[u]; i != -1; i = edges[i].next {
			if color[edges[i].v] == Gray {
				return dfs(edges[i].v, -c)
			}
			if color[u] == color[edges[i].v] {
				return true
			}
		}
		return false
	}

	var isExist bool
	for i := 1; i <= n; i++ {
		if color[i] == Gray {
			isExist = isExist || dfs(i, Black)
		}
		if isExist {
			break
		}
	}
	return isExist
}

func main() {
	// 1. 奇环
	InitGraph()
	AddEdge(1, 2)
	AddEdge(2, 3)
	AddEdge(3, 1)
	fmt.Println(IsExistOddCycle(3))

	// 2. 偶环
	InitGraph()
	AddEdge(1, 2)
	AddEdge(2, 1)
	fmt.Println(IsExistOddCycle(2))

	// 3. 无环
	InitGraph()
	AddEdge(1, 2)
	AddEdge(2, 3)
	fmt.Println(IsExistOddCycle(3))

	// 4. 奇环 + 偶环
	InitGraph()
	AddEdge(1, 2)
	AddEdge(2, 3)
	AddEdge(3, 1)
	AddEdge(4, 5)
	AddEdge(5, 6)
	fmt.Println(IsExistOddCycle(6))
}
