package main

// 题目链接：https://leetcode.cn/problems/longest-consecutive-sequence/
//
// 解题思路
// 		记忆化搜索

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	n, ans := len(nums), 0

	mark := make(map[int]bool)
	for i := 0; i < n; i++ {
		mark[nums[i]] = true
	}

	dp := make(map[int]int, n)
	var calc func(x int) int
	calc = func(x int) int {
		if !mark[x] { // 不存在该元素，直接返回 0
			return 0
		}

		if r, ok := dp[x]; ok {
			return r
		}

		dp[x] = calc(x+1) + 1
		// dp[x] = calc(x-1) + 1
		return dp[x]
	}

	for i := 0; i < n; i++ {
		ans = maxInt(ans, calc(nums[i]))
	}

	return ans
}
