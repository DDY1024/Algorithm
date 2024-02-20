package main

// https://leetcode.cn/problems/making-a-large-island/
//
// 解题思路：简单的利用并查集进行统计计数的题目

func initSet(parent []int, n int) {
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
}

func findSet(parent []int, x int) int {
	if parent[x] < 0 {
		return x
	}
	parent[x] = findSet(parent, parent[x])
	return parent[x]
}

func unionSet(parent []int, x, y int) {
	rx, ry := findSet(parent, x), findSet(parent, y)
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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestIsland(grid [][]int) int {
	dx, dy := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}
	n, m := len(grid), len(grid[0])
	parent := make([]int, n*m)
	initSet(parent, n*m)

	zeroCnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				zeroCnt++
				continue
			}

			for k := 0; k < 4; k++ {
				ii, jj := i+dx[k], j+dy[k]
				if ii >= 0 && ii < n && jj >= 0 && jj < m && grid[ii][jj] == 1 {
					unionSet(parent, i*m+j, ii*m+jj)
				}
			}
		}
	}

	if zeroCnt == 0 {
		return n * m
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				mark, cc := make(map[int]bool, 4), 1
				for k := 0; k < 4; k++ {
					ii, jj := i+dx[k], j+dy[k]
					if ii >= 0 && ii < n && jj >= 0 && jj < m && grid[ii][jj] == 1 {
						pp := findSet(parent, ii*m+jj)
						if !mark[pp] {
							cc += -parent[pp]
							mark[pp] = true
						}
					}
				}
				ans = maxInt(ans, cc)
			}
		}
	}
	return ans
}
