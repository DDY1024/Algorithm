package main

// 题目链接: https://leetcode.cn/problems/predict-the-winner/

// 该题与【石子游戏】相同，采样相同的算法求解即可

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			dp[i][j] = -0x3f3f3f3f
			dp[i][j] = maxInt(dp[i][j], nums[i]-dp[i+1][j])
			dp[i][j] = maxInt(dp[i][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}
