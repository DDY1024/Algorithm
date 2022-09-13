package main

// 参考资料：https://zhuanlan.zhihu.com/p/93647900

var parent []int

func initSet(n int) {
	parent = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
}

func findSet(x int) int {
	if parent[x] < 0 {
		return x
	}
	parent[x] = findSet(parent[x]) // 路径压缩
	return parent[x]
}

func unionSet(x, y int) {
	rx, ry := findSet(x), findSet(y)
	if rx == ry {
		return
	}

	if parent[rx] < parent[ry] {
		parent[rx] += parent[ry]
		parent[ry] = rx
	} else {
		parent[ry] += parent[rx]
		parent[rx] = ry
	}
}
