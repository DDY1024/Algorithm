package main

// 参考资料：https://blog.csdn.net/thexue/article/details/121916959

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 1. prim 算法（加点法）
func prim(adj [][]int, n int) int {
	dis := make([]int, n)
	vis := make([]bool, n)
	for i := 0; i < n; i++ {
		dis[i] = 0x3f3f3f3f
	}
	dis[0] = 0

	res := 0
	for i := 0; i < n; i++ { // 迭代 n 次
		u := -1
		for v := 0; v < n; v++ {
			if !vis[v] && (u == -1 || dis[u] > dis[v]) {
				u = v
			}
		}
		vis[u] = true
		res += dis[u]
		for v := 0; v < n; v++ {
			if !vis[v] {
				dis[v] = minInt(dis[v], adj[u][v])
			}
		}
	}
	return res
}

// 2. kruskal 算法（加边法）
// 	a. 所有边按照权值从小到大排序
//  b. 从小到大选择边，利用并查集判断两个顶点是否已经处于同一集合中（并查集应用）
//  c. 选择出 n-1 条边，构成最小生成树
