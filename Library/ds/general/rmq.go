package main

import (
	"math"
)

// 1. 区间倍增 DP
// 2. 静态查询，一次性预处理，不允许动态修改
// 3. 区间满足 "加法" 性质，均可以采用 rmq 方法来处理
// 4. 预处理时间复杂度 O(N * logN)

var ls = func(x int) int { return 1 << uint(x) }

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calcExp(n int) int {
	// 此处如果为了确保精度，可以直接循环计算出最大的 exp，满足 2^exp <= n
	return int(math.Floor(math.Log2(float64(n)))) // math.Floor
}

func getMax(l, r int, dp [][]int) int {
	exp := calcExp(r - l + 1)
	return maxInt(dp[l][exp], dp[r-ls(exp)+1][exp])
}

func getMin(l, r int, dp [][]int) int {
	exp := calcExp(r - l + 1)
	return minInt(dp[l][exp], dp[r-ls(exp)+1][exp])
}

// dp[i][j] 表示以 i 开始区间长度为 1<<j 的最值
// dp[i][j] = OP{ dp[i][j-1], dp[i+ls(j-1)][j-1] }，其中 OP 表示任何区间可叠加类的操作
func rmqInit(n int, arr []int) [][]int {
	exp := calcExp(n)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, exp+1)
		dp[i][0] = arr[i]
	}

	for l := 1; l <= exp; l++ {
		for i := 0; i+ls(l)-1 < n; i++ {
			dp[i][l] = maxInt(dp[i][l-1], dp[i+ls(l-1)][l-1])
			// dp[i][l] = minInt(dp[i][l-1], dp[i+ls(l-1)][l-1])
		}
	}
	return dp
}
