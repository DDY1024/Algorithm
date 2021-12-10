package main

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://leetcode.com/contest/biweekly-contest-42/problems/minimum-adjacent-swaps-for-k-consecutive-ones/
// 中位数原理：中位数位置距离其它节点的距离之和是最小的
// 我们维护大小为 k 的窗口计算中位数的交换距离之和即可，统计计数时用到树状数组
// 利用树状数组统计 前缀和 + 区间和
// 中位数法则 --> 树状数组进行统计

func lowBit(x int) int {
	return x & (-x)
}

func add(idx, c, n int, tree []int) {
	for i := idx; i <= n; i += lowBit(i) {
		tree[i] += c
	}
}

func get(idx int, tree []int) int {
	sum := 0
	for i := idx; i > 0; i -= lowBit(i) {
		sum += tree[i]
	}
	return sum
}

func minMoves(nums []int, k int) int {
	n := len(nums)
	que := make([]int, 0, n)
	tree := make([]int, n+1)
	ans := 0x3f3f3f3f3f3f3f3f // default int64
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			que = append(que, i+1)
			add(i+1, i+1, n, tree)
		}
		if len(que) > k {
			add(que[0], -que[0], n, tree)
			que = que[1:]
		}
		if len(que) == k {
			midIdx, l, r := que[k/2], k/2, k-1-k/2
			tmpCost := l*midIdx - get(midIdx-1, tree)
			tmpCost -= (l + 1) * l / 2
			tmpCost += get(n, tree) - get(midIdx, tree) - r*midIdx
			tmpCost -= (r + 1) * r / 2
			ans = minInt(ans, tmpCost)
		}
	}
	return ans
}
