package main

// 参考文章：https://zhuanlan.zhihu.com/p/58727559
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 1. 二分答案 + 连通性判断(dfs、bfs)
// 2. 直接求解(spfa)

func minimumEffortPath(heights [][]int) int {
	var (
		n    = len(heights)
		m    = len(heights[0])
		inf  = 0x3f3f3f3f3f3f3f3f
		dx   = []int{-1, 1, 0, 0}
		dy   = []int{0, 0, -1, 1}
		dist = make([]int, n*m)
	)

	dist[0] = 0
	for i := 1; i < n*m; i++ {
		dist[i] = inf
	}

	que := make([]int, 0, n*m)
	inq := make([]bool, n*m)
	que = append(que, 0)
	inq[0] = true
	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		inq[u] = false // 出队

		x, y := u/m, u%m
		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < n && yy >= 0 && yy < m {
				v := xx*m + yy
				w := maxInt(dist[u], absInt(heights[x][y]-heights[xx][yy]))
				if dist[v] > w { // 扩展
					dist[v] = w
					if !inq[v] {
						inq[v] = true
						que = append(que, v) // 入队
						// 判断负权环：统计节点入队次数 > n
					}
				}
			}
		}
	}
	return dist[n*m-1]
}
