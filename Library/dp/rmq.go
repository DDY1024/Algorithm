package main

import (
	"math"
)

// RMQ 算法 O(n * logn) 复杂度求解区间最值

var (
	dp  [][]int
	arr []int
)

func initRMQ() {
	n := len(arr)
	m := log2(n)
	dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = arr[i]
	}
	for l := 1; l <= m; l++ {
		for i := 0; i+(1<<l)-1 < n; i++ {
			dp[i][l] = maxInt(dp[i][l-1], dp[i+(1<<(l-1))][l-1])
		}
	}
}

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

func log2(x int) int {
	// 2^y <= x 的最大 y 取值
	return int(math.Floor(math.Log2(float64(x))))
}

// 满足区间加法性
func query(l, r int) int {
	sz := log2(r - l + 1)
	return maxInt(dp[l][sz], dp[r-(1<<sz)+1][sz])
}
