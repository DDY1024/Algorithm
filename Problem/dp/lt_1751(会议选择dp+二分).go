package main

import "sort"

// 题目链接：https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended-ii/
//
// 解题思路：同 1235，同样可以采用 dp + 二分的求解思路

func maxValue(events [][]int, k int) int {
	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(events)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			idx := sort.Search(i, func(pos int) bool { // i 和 i-1 均可，因为 i 必然满足
				return events[pos][1] >= events[i-1][0] // 此处 >= 判断
			})

			dp[i][j] = maxInt(dp[i-1][j], dp[idx][j-1]+events[i-1][2])
		}
	}

	return dp[n][k]
}
