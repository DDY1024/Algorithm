package main

import "sort"

// 参考资料: https://zhuanlan.zhihu.com/p/121032448

// 最长非严格上升子序列 a1 <= a2 <= ... ，O(N*logN) 解法
func getLIS(arr []int) int {
	n := len(arr)
	inf := 0x3f3f3f3f
	dp := make([]int, 0, n+1)
	dp = append(dp, -inf) // dp[0] = -inf

	for i := 0; i < n; i++ {
		if dp[len(dp)-1] <= arr[i] {
			dp = append(dp, arr[i])
		} else {
			idx := sort.Search(len(dp), func(idx int) bool {
				return dp[idx] > arr[i]
			})
			dp[idx] = arr[i] // >
		}
	}
	return len(dp) - 1
}

// 最长严格上升子序列 a1 < a2 < ...
func getLIS2(arr []int) int {
	n := len(arr)
	inf := 0x3f3f3f3f
	dp := make([]int, 0, n+1)
	dp = append(dp, -inf)

	for i := 0; i < n; i++ {
		if dp[len(dp)-1] < arr[i] {
			dp = append(dp, arr[i])
		} else {
			idx := sort.Search(len(dp), func(idx int) bool {
				return dp[idx] >= arr[i] // >=
			})
			dp[idx] = arr[i]
		}
	}
	return len(dp) - 1
}
