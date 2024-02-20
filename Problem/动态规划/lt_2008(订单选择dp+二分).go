package main

import "sort"

// 题目链接：https://leetcode.cn/problems/maximum-earnings-from-taxi/
//
// 解题思路：同样采用 dp + 二分的解题思路

func maxTaxiEarnings(n int, rides [][]int) int64 {
	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	m := len(rides)
	dp := make([]int, m+1)

	sort.Slice(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1]
	})

	for i := 1; i <= m; i++ {
		idx := sort.Search(i, func(j int) bool {
			return rides[j][1] > rides[i-1][0]
		})
		dp[i] = maxInt(dp[i-1], dp[idx]+rides[i-1][1]-rides[i-1][0]+rides[i-1][2])
	}

	return int64(dp[m])
}
