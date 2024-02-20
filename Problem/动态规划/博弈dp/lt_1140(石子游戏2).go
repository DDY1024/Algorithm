package main

import "math"

// 题目链接：https://leetcode.cn/problems/stone-game-ii/description/

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 思路一：先手比后手多拿多少
// dp(i,m): 表示从 i 开始，可以获取 [1,2m] 块石头时，先手比后手可以多获取的石头数量
// dp(0,1) 表示先手一开始可以比后手最多多获取多少石头
// (dp(0,1)+sum)/2 即为先手获取的石头数
func stoneGameII(piles []int) int {
	n := len(piles)
	pSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pSum[i] = pSum[i-1] + piles[i-1]
	}

	type pair struct{ i, m int }
	dp := make(map[pair]int)
	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i+2*m > n { // 可以一次性拿完
			return pSum[n] - pSum[i-1]
		}

		if v, ok := dp[pair{i, m}]; ok {
			return v
		}

		ret := math.MinInt
		for j := 1; j <= 2*m; j++ {
			// 先手决策中获取的石头数 - 后手面对残局先手多获取的石头数 = 先手面对当前局面能够多获取的石头数
			ret = maxInt(ret, pSum[i+j-1]-pSum[i-1]-dfs(i+j, maxInt(j, m)))
		}

		dp[pair{i, m}] = ret
		return ret
	}

	return (dfs(1, 1) + pSum[n]) / 2
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 思路二: 先手最多可以拿多少
// 面对任何一个局面，因为 i 确定以后，所有【石头总价值便确定了】；先手拿的多意味着后手拿的少;
// 利用后缀和，我们其实可以直接求解 dp(i.m) 能够获取的最大价值和
// dp(i,m): 表示从 i 开始，可以获取 [1,2m] 块石头时，先手能够获取的最大价值和
func stoneGameII2(piles []int) int {
	n := len(piles)
	suffix := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = piles[i]
		if i+1 < n {
			suffix[i] += suffix[i+1]
		}
	}

	type pair struct{ i, m int }
	dp := make(map[pair]int)

	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i+2*m >= n { // 可以一次性拿完
			return suffix[i]
		}

		if v, ok := dp[pair{i, m}]; ok {
			return v
		}

		ret := math.MinInt
		for j := 1; j <= 2*m; j++ {
			// 直接总和 - 残局先手获取的最大价值和中取最优
			ret = maxInt(ret, suffix[i]-dfs(i+j, maxInt(j, m)))
		}

		dp[pair{i, m}] = ret
		return ret
	}

	return dfs(0, 1)
}

// 思路三
// 记忆化搜索转递推: https://leetcode.cn/problems/stone-game-ii/solutions/2125753/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-jjax/
// 2 + 4 + 8 + ... + m = 2*m-2
// 2*m - 2 <= i, m <= (i+2)/2
func stoneGameII3(piles []int) int {
	n := len(piles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
	}

	for i, suffix := n-1, 0; i >= 0; i-- {
		suffix += piles[i]
		for j := 1; j <= (i+1)/2+1; j++ { // 上边界见上述推导
			if i+2*j >= n {
				dp[i][j] = suffix
				continue
			}
			for k := 1; k <= 2*j; k++ {
				dp[i][j] = maxInt(dp[i][j], suffix-dp[i+k][maxInt(j, k)])
			}
		}
	}
	return dp[0][1]
}
