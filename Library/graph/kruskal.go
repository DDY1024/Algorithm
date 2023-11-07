package main

import "sort"

// 边排序 + 并查集

type Edge struct {
	u, v, w int
}

func initSet(n int, parent []int) {
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
}

func findSet(x int, parent []int) int {
	if parent[x] < 0 {
		return x
	}

	parent[x] = findSet(parent[x], parent)
	return parent[x]
}

func unionSet(x, y int, parent []int) {
	r1, r2 := findSet(x, parent), findSet(y, parent)
	if r1 == r2 {
		return
	}

	if parent[r1] < parent[r2] {
		parent[r1] += parent[r2]
		parent[r2] = r1
	} else {
		parent[r2] += parent[r1]
		parent[r1] = r2
	}
}

func kruskal(n int, edges []Edge) int {
	// 边排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}

	res, enum := 0, 0
	for i := 0; i < len(edges) && enum < n-1; i++ {
		u, v, w := edges[i].u, edges[i].v, edges[i].w
		if findSet(u, parent) != findSet(v, parent) {
			unionSet(u, v, parent)
			res += w
			enum++
		}
	}
	if enum < n-1 {
		return -1 // 图不连通
	}
	return res
}
