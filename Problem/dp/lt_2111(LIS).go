package dp

import "sort"

// 题目链接: https://leetcode-cn.com/problems/minimum-operations-to-make-the-array-k-increasing/
// 解题思路:
// 1. k 递增性质推导出 arr[i] <= arr[i+k] <= ... 其中 0 <= i < k
// 2. 求解使得一个数组变为单调不减的最小改变操作数即 a1 <= a2 <= ... <= an，结论为 (n - 最长上升子序列问题)
// 因此，此题最终转化为 LIS 求解

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getLIS(arr []int) int {
	n := len(arr)
	dp, dpl := make([]int, 0, n+1), 1
	dp = append(dp, -0x3f3f3f3f) // 一个无穷小的值
	for i := 0; i < n; i++ {
		if dp[dpl-1] <= arr[i] {
			dp = append(dp, arr[i])
			dpl++
		} else {
			idx := sort.Search(dpl, func(idx int) bool {
				return dp[idx] > arr[i]
			})
			dp[idx] = minInt(dp[idx], arr[i])
		}
	}
	return dpl - 1
}

func kIncreasing(arr []int, k int) int {
	n, ans := len(arr), 0
	for i := 0; i < k; i++ {
		tmp := make([]int, 0, n/k+1)
		for j := i; j < n; j += k {
			tmp = append(tmp, arr[j])
		}
		ans += len(tmp) - getLIS(tmp)
	}
	return ans
}
