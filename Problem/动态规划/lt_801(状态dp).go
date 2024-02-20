package main

// 题目链接：https://leetcode.cn/problems/minimum-swaps-to-make-sequences-increasing/
//
// 解题思路：状态 dp
//
// 为了保证可以实现操作，则必然满足下述两个条件之一
// 1. nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1]
// 2. nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1]
//
// 我们在设计状态时，便可以考虑 nums1[i] 和 nums2[i] 是否发生交换，即
// dp[i][0]: 表示 nums1[i] 和 nums2[i] 不交换情况下的最少操作次数
// dp[i][1]: 表示 nums1[i] 和 nums2[i] 交换情况下的最少操作次数
//
// 根据 nums1[i]、nums1[i-1]、nums2[i]、nums2[i-1] 的大小关系进行状态转移
//

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)

	// dp[i][0]: 表示 nums1[i] 和 nums2[i] 不交换
	// dp[i][1]: 表示 nums1[i] 和 nums2[i] 交换
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0], dp[0][1] = 0, 1
	for i := 1; i < n; i++ {
		dp[i][0], dp[i][1] = n, n // 初始化为无穷大，最多交换 n 次

		// 关系 1
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			dp[i][0] = minInt(dp[i][0], dp[i-1][0])
			dp[i][1] = minInt(dp[i][1], dp[i-1][1]+1)
		}

		// 关系 2
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			dp[i][0] = minInt(dp[i][0], dp[i-1][1])
			dp[i][1] = minInt(dp[i][1], dp[i-1][0]+1)
		}

		// 相邻的 nums1[i], nums1[i-1], nums2[i], nums2[i-1] 满足的大小关系
		// 1. 只满足关系 1：i 交换，则 i-1 必然交换（相同操作）
		// 2. 只满足关系 2: i 交换，则 i-1 不交换（相反操作）
		// 3. 同时满足关系 1 和 关系 2：i 交换，则 i-1 可交换，也可不交换
	}

	return minInt(dp[n-1][0], dp[n-1][1])
}
