package main

// 题目链接：https://leetcode.cn/problems/maximum-width-ramp/
// 解题思路
//		同 1124 题目一致的解题思路，求解 A[i] <= A[j] 的最长区间

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxWidthRamp(nums []int) int {
	n := len(nums)
	stk := []int{0}

	// 同样维护一个单调递减的队列
	for i := 1; i < n; i++ {
		if nums[i] < nums[stk[len(stk)-1]] {
			stk = append(stk, i)
		}
	}

	ans := 0
	// 逆序遍历，计算最长区间
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && nums[stk[len(stk)-1]] <= nums[i] {
			ans = maxInt(ans, i-stk[len(stk)-1])
			stk = stk[:len(stk)-1]
		}
	}
	return ans
}
