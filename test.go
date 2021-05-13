package main

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numWays(steps int, arrLen int) int {
	n, m, mod := steps, minInt(steps+1, arrLen), int(1e9+7)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < minInt(i+1, m); j++ {
			dp[i][j] = dp[i-1][j]
			if j-1 >= 0 {
				dp[i][j] += dp[i-1][j-1]
			}
			if j+1 < m {
				dp[i][j] += dp[i-1][j+1]
			}
			dp[i][j] %= mod
		}
	}
	return dp[n][0]
}

func main() {

}
