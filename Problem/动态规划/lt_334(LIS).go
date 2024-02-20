package main

// 题目链接: https://leetcode-cn.com/problems/increasing-triplet-subsequence/
// LIS 经典应用，存在 O(nlogn) 解法
// 此题由于固定长度为 3，因此转移过程可以退化为 O(1)，总的时间复杂度为 O(n)
// LIS 中 O(nlogn) 求解算法

func increasingTriplet(nums []int) bool {
	dp := make([]int, 3)
	n := len(nums)
	dp[1], dp[2] = 0x3f3f3f3f3f3f3f3f, 0x3f3f3f3f3f3f3f3f
	for i := 1; i < n; i++ {
		if dp[2] < nums[i] {
			return true
		}
		if dp[1] < nums[i] && nums[i] < dp[2] {
			dp[2] = nums[i]
			continue
		}
		if nums[i] < dp[1] {
			dp[1] = nums[i]
		}
	}
	return false
}
