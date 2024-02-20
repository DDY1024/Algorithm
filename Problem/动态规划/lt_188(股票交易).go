package main

// 题目链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
// 解题思路
// 		1. 常规 dp
//      2. 第 i 天可选择的操作 不操作、买入、卖出
//      3. dp[i][j]：前 i 天发生 j 次股票交易的情况下，能够获得的最大收益（j 次操作可以是买入、卖出）

func maxProfit(k int, prices []int) int {
	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(prices)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k*2+1)
		for j := 1; j <= 2*k; j++ { // 最多 2k 次操作，包括买入、卖出
			dp[i][j] = -0x3f3f3f3f
		}
	}

	// dp[i][0]
	for i := 1; i <= n; i++ {
		for j := 1; j <= 2*k; j++ {
			dp[i][j] = dp[i-1][j]
			if j&1 > 0 {
				dp[i][j] = maxInt(dp[i][j], dp[i-1][j-1]-prices[i-1])
			} else {
				dp[i][j] = maxInt(dp[i][j], dp[i-1][j-1]+prices[i-1])
			}
		}
	}

	ret := 0
	for i := 0; i <= 2*k; i++ {
		ret = maxInt(ret, dp[n][i])
	}

	return ret
}
