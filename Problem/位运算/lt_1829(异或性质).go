package main

// 题目连接：https://leetcode.cn/problems/maximum-xor-for-each-query/description/
//
// 解题思路
// 		1. 简单异或性质

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	xor, mask := 0, (1<<maximumBit)-1
	for i := 0; i < n; i++ {
		xor ^= nums[i]
	}

	ans := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		ans = append(ans, xor^mask)
		xor ^= nums[i]
	}
	return ans
}
