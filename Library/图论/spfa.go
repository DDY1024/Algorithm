package main

// 参考洛谷的一篇经典文章：https://zhuanlan.zhihu.com/p/58727559

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
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

// spfa 算法求解最短路径
func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	dis := make([]int, n*m)
	for i := 1; i < n*m; i++ {
		dis[i] = 0x3f3f3f3f
	}
	dx, dy := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}
	que := make([]int, 0, 1000000)
	inq := make([]bool, n*m)
	que = append(que, 0)
	inq[0] = true
	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		inq[u] = false
		x, y := u/m, u%m
		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < n && yy >= 0 && yy < m {
				v := xx*m + yy
				w := maxInt(dis[u], absInt(heights[x][y]-heights[xx][yy]))
				if dis[v] > w {
					dis[v] = w
					if !inq[v] {
						inq[v] = true
						que = append(que, v)
					}
				}
			}
		}
	}
	return dis[n*m-1]
}

// spfa 算法不仅仅可以是一维的也可以是多维的
