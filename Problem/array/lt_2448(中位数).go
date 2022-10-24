package array

import "sort"

// 题目链接：https://leetcode.cn/problems/minimum-cost-to-make-array-equal/
// 解题报告：https://leetcode.cn/problems/minimum-cost-to-make-array-equal/
//
// 题目总结
// 		1. 相比于 462 题，新增了一个限制条件即每个数 nums[i] 每次操作代价由 1 变成 cost[i]
//      2. 1 我们可以等价为存在 cost[i] 个 nums[i]；且每个数的操作代价仍为 1
//      3. 经过 2 的等价转化，问题变成了在 sum_cost 个数中操作使得每个数相等的最小操作次数；中位数贪心选择仍然成立

type Pair struct {
	x, c int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minCost(nums, cost []int) int64 {
	arr := make([]Pair, 0, len(nums))
	sumCost := 0
	for i, c := range cost {
		arr = append(arr, Pair{nums[i], c})
		sumCost += c
	}

	// 按照 nums[i] 从小到大排序
	sort.Slice(arr, func(i, j int) bool { return arr[i].x < arr[j].x })

	// 选择第 (sumCost+1)/2 个数作为中位数
	pos, mid, choose, ret := 0, (sumCost+1)/2, 0, int64(0)
	for _, p := range arr {
		pos += p.c
		if pos >= mid { // 把所有数变成 p.x
			choose = p.x
			break
		}
	}

	for _, p := range arr {
		ret += int64(abs(p.x-choose)) * int64(p.c)
	}
	return ret
}
