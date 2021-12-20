package main

import "sort"

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getLIS(arr []int) int {
	n := len(arr)
	dp, dpl := make([]int, 0, n+1), 1
	dp = append(dp, -0x3f3f3f3f) // 一个无穷小的值
	for i := 0; i < n; i++ {
		if dp[dpl-1] <= arr[i] {
			dp = append(dp, arr[i])
			dpl++
		} else {
			idx := sort.Search(dpl, func(idx int) bool {
				return dp[idx] > arr[i]
			})
			dp[idx] = minInt(dp[idx], arr[i])
		}
	}
	return len(dp) - 1
}
