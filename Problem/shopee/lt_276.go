package main

// 题目链接: https://leetcode-cn.com/problems/paint-fence/

// 注意: 边界 0 的系数为 k
func numWays(n int, k int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = k
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] * (k - 1)
		if i-2 >= 0 {
			if i-2 == 0 {
				dp[i] += dp[0] * k
			} else {
				dp[i] += dp[i-2] * (k - 1)
			}
		}
	}
	return dp[n]
}

func main() {
	numWays(3, 2)
}

// 3, 2
// dp[1] = 2
//
