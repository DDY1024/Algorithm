package main

import "sort"

// https://leetcode.cn/problems/minimum-initial-energy-to-finish-tasks/

// 自定义排序（贪心思路）
//
// 假设前 x 个任务已经按照最优序列完成，剩余最后两个任务；这两个任务应该怎样安排，才能使得需要的初始值最小
// 不难想到取  max{ cost+b1, cost+a1+b2}、max{cost+b2, cost+a2+b1} 中的最小者
// 这便定义出了我们贪心的排序比较规则

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		// 贪心排序比较规则
		return max(tasks[i][1], tasks[i][0]+tasks[j][1]) <= max(tasks[j][1], tasks[j][0]+tasks[i][1])
	})

	ans, cost := 0, 0
	for i := 0; i < len(tasks); i++ {
		ans = max(ans, cost+tasks[i][1])
		cost += tasks[i][0]
	}
	ans = max(ans, cost)
	return ans
}
