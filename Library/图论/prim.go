package main

func prim(n int, g [][]int) int {
	var (
		inf = 0x3f3f3f3f3f3f3f3f
		dis = make([]int, n)
		vis = make([]bool, n)
	)

	dis[0] = 0
	for i := 1; i < n; i++ {
		dis[i] = inf
	}

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

		if u == -1 { // 图不连通，直接返回
			return inf
		}

		vis[u] = true
		res += dis[u]
		for v := 0; v < n; v++ {
			if !vis[v] && dis[v] > g[u][v] {
				dis[v] = g[u][v]
			}
		}
	}
	return res
}
