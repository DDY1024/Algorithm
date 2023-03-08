package main

// 题目链接：https://leetcode.cn/problems/stone-game/description/

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp[i][j] 表示 [i,j] 区间内的石子先手比后手最多多取多少
func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}

	inf := 0x3f3f3f3f
	// 此题状态定义 A 比 B 多多少，更方便求解
	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			dp[i][j] = -inf
			dp[i][j] = maxInt(dp[i][j], piles[i]-dp[i+1][j])
			dp[i][j] = maxInt(dp[i][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] > 0
}
