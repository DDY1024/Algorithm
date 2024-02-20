package main

// 题目链接: https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/
// 重点理解状态转移方程
// dp[i][j]: 猜测区间 [i,j] 范围内的数字需要的最少花费
// dp[i][j] = min{max{dp[i][k-1]+dp[k+1][j]}+k}，其中 i <= k <= j
func getMoneyAmount(n int) int {

	var minInt = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int, n+2)
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
