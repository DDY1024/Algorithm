package main

// 题目链接: https://leetcode-cn.com/problems/knight-probability-in-chessboard/
// 解题思路：显而易见的一道概率 dp 的题目

func knightProbability(n int, k int, row int, column int) float64 {
	dp := make([][][]float64, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]float64, n)
		}
	}
	dp[0][row][column] = 1.0

	dx := []int{-2, -1, 1, 2, 2, 1, -1, -2}
	dy := []int{-1, -2, -2, -1, 1, 2, 2, 1}

	for i := 0; i < k; i++ {
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				for idx := 0; idx < 8; idx++ {
					rr, cc := r+dx[idx], c+dy[idx]
					if rr >= 0 && rr < n && cc >= 0 && cc < n {
						dp[i+1][rr][cc] += dp[i][r][c] * 0.125
					}
				}
			}
		}
	}

	ret := float64(0.0)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			ret += dp[k][r][c]
		}
	}
	return ret
}
