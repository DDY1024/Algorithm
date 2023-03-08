package main

// 题目链接: https://leetcode.cn/problems/dice-roll-simulation/description/
//
// 解题思路: 组合计数 dp
//
// 方法一
// dp[i][j][k]: 表示前 i 个数且末尾出现连续 k 个 j 时的方案数
//		当 k > 1 时，dp[i][j][k] = dp[i-1][j][k-1]
//      当 k = 1 时，dp[i][j][1] = sum{dp[i-1][x][y]}，其中 x != j，1 <= y <= rollMax[x]，即上一轮末尾不是 j 的方案数之和

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dieSimulator(n int, rollMax []int) int {
	maxRoll := 0
	for i := 0; i < 6; i++ {
		maxRoll = maxInt(maxRoll, rollMax[i])
	}

	dp := make([][][]int, n)
	mod := int(1e9) + 7

	for i := 0; i < n; i++ {
		dp[i] = make([][]int, 6)
		for j := 0; j < 6; j++ {
			dp[i][j] = make([]int, maxRoll+1)
		}
	}

	for i := 0; i < 6; i++ {
		dp[0][i][1] = 1
		dp[0][i][0] = 1 // dp[i][j][0] 特殊存储 dp[i][j][1] ~ dp[i][j][rollMax[j]] 的和
	}

	for i := 1; i < n; i++ {
		sum := 0
		for j := 0; j < 6; j++ {
			sum += dp[i-1][j][0]
			sum %= mod
		}
		for j := 0; j < 6; j++ {
			for k := 2; k <= rollMax[j]; k++ {
				dp[i][j][k] = dp[i-1][j][k-1]

				dp[i][j][0] += dp[i][j][k]
				dp[i][j][0] %= mod
			}
			dp[i][j][1] = (sum - dp[i-1][j][0]) % mod

			dp[i][j][0] += dp[i][j][1]
			dp[i][j][0] %= mod
		}
	}

	ret := 0
	for i := 0; i < 6; i++ {
		ret += dp[n-1][i][0]
		ret %= mod
	}
	return (ret + mod) % mod
}

// 方法二（更方便）
// dp[i][j]: 表示前 i 个数且末尾是 j 的合法序列数
func dieSimulatorTwo(n int, rollMax []int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 7)
	}

	// dp[i][0]: 特殊位置表示 dp[i][1] ~ dp[i][6] 方案数之和
	// 为方便处理序列全被某一数字占满的情况时，方案数为 1，因此特殊初始化 dp[0][0] = 1
	dp[0][0] = 1

	mod := int(1e9) + 7
	for i := 1; i <= n; i++ {
		dp[i][0] = 0
		for j := 1; j <= 6; j++ {
			for k := 1; k <= rollMax[j-1] && i-k >= 0; k++ {
				dp[i][j] += dp[i-k][0] - dp[i-k][j] // dp[0][0] - dp[0][j] = 1
				dp[i][j] %= mod
			}
			dp[i][0] += dp[i][j]
			dp[i][0] %= mod
		}
	}
	return (dp[n][0] + mod) % mod
}
