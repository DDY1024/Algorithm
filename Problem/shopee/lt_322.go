package main

// 题目链接: https://leetcode-cn.com/problems/coin-change/
// 零钱兑换
// 1. 完全背包问题 --> 动态规划
// 状态转移方程: dp[i][j] = min{dp[i-1][j], dp[i][j-coins[i-1]]+1}

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			dp[i][j] = 0x3f3f3f3f
		}
		dp[i][0] = 0
	}

	var minInt = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j-coins[i-1] >= 0 {
				dp[i][j] = minInt(dp[i][j], dp[i][j-coins[i-1]]+1)
			}
		}
	}

	if dp[n][amount] >= 0x3f3f3f3f {
		return int(-1)
	}

	return dp[n][amount]
}
