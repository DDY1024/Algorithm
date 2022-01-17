package main

// 题目链接: https://leetcode-cn.com/problems/count-vowels-permutation/
// 题目大意:
// 求解满足条件的元音字母序列的个数，运用简单的计数 dp 进行求解

func countVowelPermutation(n int) int {
	mod := int(1e9 + 7)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 5)
	}

	// 维护转化关系
	g := [][]int{
		{4, 2, 1},
		{0, 2},
		{3, 1},
		{2},
		{3, 2},
	}

	for i := 0; i < 5; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 0; j < 5; j++ {
			for _, k := range g[j] {
				dp[i][j] += dp[i-1][k]
				dp[i][j] %= mod
			}
		}
	}

	ret := 0
	for i := 0; i < 5; i++ {
		ret += dp[n][i]
		ret %= mod
	}
	return ret
}
