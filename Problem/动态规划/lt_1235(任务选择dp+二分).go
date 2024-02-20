package main

import "sort"

// 题目链接: https://leetcode.cn/problems/maximum-profit-in-job-scheduling/
// 题目大意: 如何安排任务，获得最大收益？
//
// 解题报告：https://leetcode.cn/problems/maximum-profit-in-job-scheduling/
// 解题思路：
// 		1. 按照任务结束时间排序
//      2. dp[i] 表示前 i 个任务能够获取的最大收益
//			由于任务按照结束时间进行了排序，因此可以二分查找任务 k 使得其结束时间最接近 start[i]
// 			dp[i] = max{ dp[i-1], dp[k] + profit[i] }，第 i 个任务做还是不做?

func jobScheduling(startTime, endTime, profit []int) int {
	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(startTime)
	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}

	// 直接排序二维数组
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1] // 按照任务结束时间从小到大排序
	})

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 此处是 > 判断
		k := sort.Search(i, func(j int) bool { return jobs[j][1] > jobs[i-1][0] }) // 寻找最小的任务 k，其结束时间大于当前任务的开始时间
		dp[i] = maxInt(dp[i-1], dp[k]+jobs[i-1][2])
	}

	return dp[n]
}
