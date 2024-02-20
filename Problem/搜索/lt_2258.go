package main

// https://leetcode.cn/problems/escape-the-spreading-fire/description/?envType=daily-question&envId=2023-11-09
// 1. 二分 + bfs 判定是否能抵达
// 2. 单源 bfs 搜索 vs 多源 bfs 搜索

func maximumMinutes(grid [][]int) int {
	n, m, inf := len(grid), len(grid[0]), 0x3f3f3f3f3f3f3f3f
	dx, dy := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}
	fireT := make([][]int, n)
	for i := 0; i < n; i++ {
		fireT[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fireT[i][j] = inf
		}
	}

	q := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				fireT[i][j] = 0
				q = append(q, i*m+j)
			}
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		x, y := u/m, u%m
		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < n && yy >= 0 && yy < m && grid[xx][yy] != 2 && fireT[xx][yy] == inf {
				fireT[xx][yy] = fireT[x][y] + 1
				q = append(q, xx*m+yy)
			}
		}
	}
	// fmt.Println("Test:", fireT)

	var check = func(st int) bool {
		vis := make([][]bool, n)
		dis := make([][]int, n)
		for i := 0; i < n; i++ {
			vis[i] = make([]bool, m)
			dis[i] = make([]int, m)
		}
		vis[0][0] = true
		dis[0][0] = st
		q := make([]int, 0)
		q = append(q, 0)
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			x, y := u/m, u%m
			for i := 0; i < 4; i++ {
				xx, yy := x+dx[i], y+dy[i]
				if xx >= 0 && xx < n && yy >= 0 && yy < m && grid[xx][yy] == 0 && !vis[xx][yy] {
					if xx == n-1 && yy == m-1 { // 终点判断
						return fireT[xx][yy] >= dis[x][y]+1 // >=
					}
					if fireT[xx][yy] > dis[x][y]+1 { // >
						dis[xx][yy] = dis[x][y] + 1
						vis[xx][yy] = true
						q = append(q, xx*m+yy)
					}
				}
			}
		}
		return false
	}

	// n*m 为上限值，理论上不需要
	l, r, ret := 0, n*m, -1
	for l <= r {
		mid := l + (r-l)/2
		if check(mid) {
			ret = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if ret >= n*m {
		return int(1e9)
	}
	return ret
}
