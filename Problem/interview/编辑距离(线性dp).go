package main

// https://leetcode.cn/problems/edit-distance/?favorite=2cktkvj
//
// 编辑距离：经典动态规划问题

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

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	// dp(i, 0) = i  --> 删除
	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}

	// dp(0, i) = i  --> 插入
	for i := 1; i <= m; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 删除、替换、插入三种操作
				dp[i][j] = minInt(dp[i-1][j], minInt(dp[i-1][j-1], dp[i][j-1])) + 1
			}
		}
	}
	return dp[n][m]
}
