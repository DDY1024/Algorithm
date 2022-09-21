package main

// https://leetcode.cn/problems/longest-consecutive-sequence/
// 通过记忆化搜索的方法来实现 O(n) 复杂度

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 一种巧妙的 O(n) 复杂度解法：记忆化搜索
func longestConsecutive(nums []int) int {
	n, ans := len(nums), 0

	// 标记某个数是否存在
	mark := make(map[int]bool)
	for i := 0; i < n; i++ {
		mark[nums[i]] = true
	}

	// 由于 -10^9 <= nums[i] <= 10^9 数据范围比较大，因此采用记忆化搜索的方式进行求解
	dp := make(map[int]int, n)
	var calc func(x int) int
	calc = func(x int) int {
		if !mark[x] {
			return 0
		}
		if r, ok := dp[x]; ok {
			return r
		}
		dp[x] = calc(x+1) + 1
		return dp[x]
	}
	for i := 0; i < n; i++ {
		ans = maxInt(ans, calc(nums[i]))
	}
	return ans
}
