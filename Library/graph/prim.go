package main

// 参考资料：https://blog.csdn.net/thexue/article/details/121916959

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// prim 算法（加点法）
func prim(adj [][]int, n int) int {
	dis := make([]int, n)
	vis := make([]bool, n)
	for i := 0; i < n; i++ {
		dis[i] = 0x3f3f3f3f
	}
	dis[0] = 0

	res := 0
	for i := 0; i < n; i++ {
		u := -1
		for v := 0; v < n; v++ {
			if vis[v] {
				continue
			}
			if u == -1 || dis[u] > dis[v] {
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
