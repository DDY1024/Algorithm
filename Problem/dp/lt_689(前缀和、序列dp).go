package main

// 题目链接: https://leetcode-cn.com/problems/maximum-sum-of-3-non-overlapping-subarrays/
// 解题思路: 一种典型的序列 dp，但是本题需要求解具体的方案，在计算最优值的同时，保存状态转移的方案
// dp[i][j]: 前 i 个数分成 j 段能够获得的最大和
// path[i][j]: 存储状态转移，选择第 j 个区间的右端点下标

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	inf := 0x3f3f3f3f
	dp := make([][]int, n+1)
	path := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 4)
		path[i] = make([]int, 4)
		for j := 0; j <= 3; j++ {
			dp[i][j] = -inf
			path[i][j] = -1
		}
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}

	var calc = func(i, j int) int {
		ret := preSum[j]
		if i > 0 {
			ret -= preSum[i-1]
		}
		return ret
	}

	for i := k; i <= n; i++ {
		for j := 1; j <= 3; j++ {
			dp[i][j] = dp[i-1][j]
			path[i][j] = path[i-1][j]
			if dp[i-k][j-1] >= 0 && dp[i-k][j-1]+calc(i-k+1, i) > dp[i][j] {
				dp[i][j] = dp[i-k][j-1] + calc(i-k+1, i)
				path[i][j] = i // 存储区间右端点
			}
		}
	}

	ret := make([]int, 3)
	idx, num := n, 3
	for num > 0 {
		ret[num-1] = path[idx][num] - k // 下标索引从 0 开始，需要 -1 操作
		idx, num = path[idx][num]-k, num-1
	}
	return ret
}
