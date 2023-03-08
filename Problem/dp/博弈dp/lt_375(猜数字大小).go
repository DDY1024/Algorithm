package main

// 题目链接：https://leetcode.cn/problems/guess-number-higher-or-lower-ii/

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

// dp[i][j]: [i,j] 区间猜对数字需要的最少代价
// dp[i][j] = min{max{dp[i][k-1],dp[k+1][j]}+k}
// 上述这类 最小+最大 或 最大+最小 的状态转移思路
func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int, n+2)
		dp[i][i] = 0
	}

	for l := 2; l <= n; l++ {
		for i := 1; i+l-1 <= n; i++ {
			j := i + l - 1
			dp[i][j] = 0x3f3f3f3f
			for k := i; k <= j; k++ {
				dp[i][j] = minInt(dp[i][j], k+maxInt(dp[i][k-1], dp[k+1][j]))
			}
		}
	}
	return dp[1][n]
}
