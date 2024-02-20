package main

import "sort"

// 题目链接：https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/
//
// 题目大意：找出 nums 中和至少为 k 的 最短非空子数组
//
// 由于 -10^5 <= nums[i] <= 10^5，因此整个从左到右的前缀和不存在单调性
// 但是，由于我们求解的是 [最短非空子数组]，假设存在 i, j 的前缀和相等且 i < j，我们容易知道 j 肯定优于 i
// 因此 i 没有存在的必要性，可以根据此维护一个单调栈；每次遇到 i，我们在单调栈中进行二分查找确定最优点即可
// 最终复杂度为 O(nlogn)

func shortestSubarray(nums []int, k int) int {
	type pair struct{ s, p int }
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(nums)
	que := make([]pair, 0, n+1)
	que = append(que, pair{0, -1})
	psum, ret := 0, n+1
	for i := 0; i < n; i++ {
		psum += nums[i]

		// 二分查找确定边界点
		idx := sort.Search(len(que), func(j int) bool {
			return que[j].s > psum-k
		}) - 1
		if idx >= 0 {
			ret = min(ret, i-que[idx].p)
		}

		// 维护单调栈的单调性
		for len(que) > 0 && que[len(que)-1].s >= psum {
			que = que[:len(que)-1]
		}
		que = append(que, pair{psum, i})
	}

	if ret > n {
		return -1
	}
	return ret
}

// 更进一步，在 1 解法的基础上；我们维护了一个单调递增栈；假设当前遍历到下标 i，在单调栈中存在几个元素满足子数组和 >= k
// 由于后续的 i+1 在处理这几个元素时，其结果必然不会优于 i，因此这些结果我们可以在遍历到 i 时，直接剔除掉。
// 这样下来整个循环遍历过程中，我们维护单调队列的均摊复杂度变为 O(1)；算法最终复杂度变为 O(n)。

func shortestSubarray2(nums []int, k int) int {
	type pair struct{ s, p int }
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(nums)
	que := make([]pair, 0, n+1)
	que = append(que, pair{0, -1})
	psum, ret := 0, n+1
	for i := 0; i < n; i++ {
		psum += nums[i]
		// 剔除掉 i 为右端点时满足的元素；i+x 时结果不会更优
		for len(que) > 0 && psum-que[0].s >= k {
			ret = min(ret, i-que[0].p)
			que = que[1:]
		}

		// 维护单调栈的单调性
		for len(que) > 0 && que[len(que)-1].s >= psum {
			que = que[:len(que)-1]
		}
		que = append(que, pair{psum, i})
	}

	if ret > n {
		return -1
	}
	return ret
}
