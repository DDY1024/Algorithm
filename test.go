package main

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func rearrangeSticks(n int, k int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	mod := int(1e9 + 7)
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= minInt(i, k); j++ {
			dp[i][j] = dp[i-1][j] * (n - i) % mod
			if j-1 >= 0 {
				dp[i][j] += dp[i-1][j-1]
				dp[i][j] %= mod
			}
		}
	}
	return dp[n][k]
}
