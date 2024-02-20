package main

// 题目链接：https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/description/
// 解题报告: https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/description/
//
// 解题思路
// 		本题和 1124 题正好反过来，一个求解最长，一个求解最短
// 1. 同样考虑前缀和 preSum[i]、preSum[j]，i < j； 如果 preSum[j] <= preSum[i]；存在它们右边的一个 preSum[k] 满足区间和 >= k
//		的条件时，我们选择 preSum[j] 更优；因此 preSum[i] 不再考虑；这样我们遍历前缀和时，会形成一个单调递增的队列
// 2. 在遍历到 preSum[i] 时，同时考虑单调队列中已有的元素，如果队列头部元素 j，满足 preSum[i] - preSum[j] >= k；则对于 j 来说
// 		右端点 i 是最优的，因为后续的右端点不可能使得结果更优
// 3. 较 1124 题求解最长子数组不同
//		a. 1124 求解最长，队列中的元素是单调递减的；本题求解最短，队列中元素是单调递增的
//      b. 1124 求解最长，右端点的选取是从 n 到 1 的；本题求解最短，在遍历每个 i 时，需要将队列中的最小前缀和及时弹出确保长度最短

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	stk := []int{0}
	ans := n + 10

	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		// 在计算 preSum[i] 时，如果单调递增队列中存在元素使得子区间和 >= k，应该及时弹出队列，因为后续区间对
		// 队列中这些元素来说不会更优
		for len(stk) > 0 && preSum[i]-preSum[stk[0]] >= k {
			ans = minInt(ans, i-stk[0])
			stk = stk[1:]
		}

		// idx := sort.Search(len(stk), func(idx int) bool {
		//     return preSum[stk[idx]] > preSum[i] - k
		// })
		// if idx > 0 {
		//     ans = minInt(ans, i-stk[idx-1])
		// }

		// 由于求解最短区间，因此对于相同的 preSum[i1] 和 preSum[i2]，我们选择更大的 i2
		// 这样我们最终实际上维护的是一个单调递增的队列
		for len(stk) > 0 && preSum[stk[len(stk)-1]] >= preSum[i] {
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}

	if ans > n {
		return -1
	}
	return ans
}
