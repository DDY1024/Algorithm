package main

// 题目链接：https://leetcode.cn/problems/paint-house-ii/
//
// 通过维护上一轮中计算出的最小值、次小值；在下一轮计算中 O(1) 复杂度进行转移
//
// 从原始状态转移方程想一想，为什么只需要维护这两个值？

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostII(costs [][]int) int {
	n, m := len(costs), len(costs[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	first, second := -1, -1
	for i := 0; i < m; i++ {
		dp[0][i] = costs[0][i]
		if first == -1 || dp[0][first] > dp[0][i] {
			first, second = i, first
		} else if second == -1 || dp[0][second] > dp[0][i] {
			second = i
		}
	}

	for i := 1; i < n; i++ {
		newFirst, newSecond := -1, -1
		for j := 0; j < m; j++ {
			if first != j {
				dp[i][j] = dp[i-1][first] + costs[i][j]
			} else {
				dp[i][j] = dp[i-1][second] + costs[i][j]
			}

			// 维护该轮计算中最小值、次小值下标；用于下轮递推优化
			if newFirst == -1 || dp[i][newFirst] > dp[i][j] {
				newFirst, newSecond = j, newFirst
			} else if newSecond == -1 || dp[i][newSecond] > dp[i][j] {
				newSecond = j
			}
		}
		first, second = newFirst, newSecond
	}

	ret := 0x3f3f3f3f
	for i := 0; i < m; i++ {
		ret = minInt(ret, dp[n-1][i])
	}
	return ret
}
