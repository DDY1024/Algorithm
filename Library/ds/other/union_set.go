package other

// 1. 普通并查集
func initSet(n int, parent []int) {
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
}

func findSet(x int, parent []int) int {
	if parent[x] < 0 { // root
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

// 2. 带权并查集、种类并查集
