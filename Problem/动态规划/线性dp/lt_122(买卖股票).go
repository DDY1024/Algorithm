package main

// 题目链接：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150

// dp[i][0]: 第 0 ~ i 天手上不持有股票的最大收益
// dp[i][1]: 第 0 ~ i 天手上持有一只股票的最大收益

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = maxInt(dp[i-1][0], dp[i-1][1]+prices[i]) // 当天卖出
		dp[i][1] = maxInt(dp[i-1][1], dp[i-1][0]-prices[i]) // 当天买入
	}
	return dp[n-1][0]
}
