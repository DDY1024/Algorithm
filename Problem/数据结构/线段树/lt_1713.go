package main

import "fmt"

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 利用线段树优化 dp 状态转移方程
// dp[i] = max{1, dp[1]...dp[i-1]}
// 求解 target 在 arr 中出现的最长子序列，利用动态规划求解，但是在状态转移方程中我们明显可以看到需要动态维护求解区间最值
// 这时我们便想到利用线段树来优化状态转移方程
// https://leetcode.com/contest/weekly-contest-222/problems/minimum-operations-to-make-a-subsequence/
//
func push_up(idx int, maxv []int) {
	maxv[idx] = maxInt(maxv[idx<<1], maxv[(idx<<1)|1])
}

func build(l, r, idx int, maxv []int) {
	if l == r {
		maxv[idx] = 0
		return
	}
	mid := (l + r) >> 1
	build(l, mid, idx<<1, maxv)
	build(mid+1, r, (idx<<1)|1, maxv)
	push_up(idx, maxv)
}

func query(l, r, s, t, idx int, maxv []int) int {
	if l <= s && t <= r {
		return maxv[idx]
	}
	mid, result := (s+t)>>1, 0
	if l <= mid {
		result = maxInt(result, query(l, r, s, mid, idx<<1, maxv))
	}
	if r > mid {
		result = maxInt(result, query(l, r, mid+1, t, (idx<<1)|1, maxv))
	}
	return result
}

func update(pos, c, s, t, idx int, maxv []int) {
	if s == t {
		maxv[idx] = c
		return
	}
	mid := (s + t) >> 1
	if pos <= mid {
		update(pos, c, s, mid, idx<<1, maxv)
	} else {
		update(pos, c, mid+1, t, (idx<<1)|1, maxv)
	}
	push_up(idx, maxv)
}

func minOperations(target []int, arr []int) int {
	n, m := len(target), len(arr)
	mark := make(map[int]int)
	for i := 0; i < n; i++ {
		mark[target[i]] = i + 1 // 由于 target 数组元素是唯一的，因此我们直接进行 mark 标记即可
	}
	maxv, ans := make([]int, (n+5)<<2), 0
	dp := make([]int, n+1)
	build(1, n, 1, maxv)
	for i := 0; i < m; i++ {
		if pos, ok := mark[arr[i]]; ok {
			if pos > 1 {
				dp[pos] = maxInt(1, query(1, pos-1, 1, n, 1, maxv)+1)
			} else {
				dp[pos] = maxInt(dp[pos], 1)
			}
			update(pos, dp[pos], 1, n, 1, maxv)
			ans = maxInt(ans, dp[pos])
		}
	}
	fmt.Println(dp)
	return n - ans
}
