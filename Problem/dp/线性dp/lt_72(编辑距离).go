package main

// 题目链接：https://leetcode.cn/problems/edit-distance/?favorite=2cktkvj

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

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 && j == 0 {
				continue
			}

			// 插入
			if i == 0 {
				dp[i][j] = dp[i][j-1] + 1
				continue
			}

			// 删除
			if j == 0 {
				dp[i][j] = dp[i-1][j] + 1
				continue
			}

			dp[i][j] = 0x3f3f3f3f
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minInt(dp[i][j], dp[i][j-1]+1)   // 插入
				dp[i][j] = minInt(dp[i][j], dp[i-1][j-1]+1) // 替换
				dp[i][j] = minInt(dp[i][j], dp[i-1][j]+1)   // 删除
			}
		}
	}
	return dp[n][m]
}
