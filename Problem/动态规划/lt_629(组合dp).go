package main

// 题目链接：https://leetcode-cn.com/problems/k-inverse-pairs-array/
// 解题思路: 组合 dp
// dp[i][j]: 前 i 个数存在逆序对数为 j 的情况下的组合方案数
// 对于第 i 个数我们可以选择放第1大的数、第2大的数、...、第 i 大的数，选择完后子问题便是
// dp[i][j] = dp[i-1][j] + dp[i-1][j-1] + ... + dp[i-1][j-(i-1)]
// 由状态转移方程可以看到，我们是可以利用前缀和进行优化的
// 边界条件: dp[i][0] = 1 (0 <= i <= n), dp[1][0] = 1, dp[1][1] = 0

func kInversePairs(n int, k int) int {
	dp := make([][]int, n+1)
	mod := int(1e9 + 7)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	for i := 2; i <= n; i++ {
		tmp := 1
		for j := 1; j <= k; j++ {
			dp[i][j] = (dp[i-1][j] + tmp) % mod
			tmp = (tmp + dp[i-1][j]) % mod
			if j >= i-1 {
				tmp -= dp[i-1][j-i+1]
				tmp = (tmp + mod) % mod
			}
		}
	}
	return dp[n][k]
}
