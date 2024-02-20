package main

import (
	"math"
	"sort"
)

// 1. 非严格上升子序列
func getLIS(arr []int) int {
	n := len(arr)
	dp := make([]int, 0)
	dp = append(dp, math.MinInt)

	for i := 0; i < n; i++ {
		if dp[len(dp)-1] <= arr[i] { // <=
			dp = append(dp, arr[i])
			continue
		}

		idx := sort.Search(len(dp), func(j int) bool {
			return dp[j] > arr[i] // >
		})
		dp[idx] = arr[i]
	}
	return len(dp) - 1
}

// 2. 严格上升子序列
func getLIS2(arr []int) int {
	n := len(arr)
	dp := make([]int, 0)
	dp = append(dp, math.MinInt)

	for i := 0; i < n; i++ {
		if dp[len(dp)-1] < arr[i] { // <
			dp = append(dp, arr[i])
			continue
		}

		idx := sort.Search(len(dp), func(j int) bool {
			return dp[j] >= arr[i] // >=
		})
		dp[idx] = arr[i]
	}
	return len(dp) - 1
}
