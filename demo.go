package main

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

func largestArea(grid []string) int {
	n, m := len(grid), len(grid[0])
	tgrid := make([][]int, n)
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		tgrid[i] = make([]int, m)
		vis[i] = make([]bool, m)
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tgrid[i][j] = int(grid[i][j] - '0')
		}
	}

	var check = func(x, y int) bool {
		if x == 0 || x == n-1 || y == 0 || y == m-1 {
			return false
		}
		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < n && yy >= 0 && yy < m && tgrid[xx][yy] == 0 {
				return false
			}
		}
		return true
	}

	var dfs func(x, y, z int) (int, bool)
	dfs = func(x, y, z int) (int, bool) {
		vis[x][y] = true
		total, flag := 1, true
		if !check(x, y) {
			flag = false
		}
		for i := 0; i < 4; i++ {
			xx := x + dx[i]
			yy := y + dy[i]
			if xx >= 0 && xx < n && yy >= 0 && yy < m && tgrid[xx][yy] == z && !vis[xx][yy] {
				num, ok := dfs(xx, yy, z)
				total += num
				flag = flag && ok
			}
		}
		return total, flag
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if tgrid[i][j] > 0 {
				if !vis[i][j] {
					num, flag := dfs(i, j, tgrid[i][j])
					if flag {
						ans = maxInt(ans, num)
					}
				}
			}
		}
	}
	return ans
}

func main() {
	largestArea([]string{"111", "222", "333"})
}
