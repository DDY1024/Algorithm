package other

// 参考自：https://oi-wiki.org/ds/dsu/
// 1. 普通并查集实现

func initSet(n int, parent []int) {
	for i := 0; i < n; i++ { // [0,n) 或 [1,n] 均可
		parent[i] = -1
	}
}

func findSet(x int, parent []int) int {
	if parent[x] < 0 { // parent[root] < 0
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

	// 按秩合并
	if parent[r1] < parent[r2] {
		parent[r1] += parent[r2]
		parent[r2] = r1
	} else {
		parent[r2] += parent[r1]
		parent[r1] = r2
	}
}

// 2. 带权并查集 or 种类并查集
