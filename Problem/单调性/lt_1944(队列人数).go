package main

// 题目链接：https://leetcode.cn/problems/number-of-visible-people-in-a-queue/description/?envType=daily-question&envId=2024-01-05
//
//
// 1. 单调栈解决
// 		如果 i < j && height[i] > height[j]，则前面的人最多看到 i，必然看不到 j；所有这样的 j 均不会出现在前面人能看到的人当中，存在单调性

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	ans := make([]int, n)
	stk := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		cnt := 0
		for len(stk) > 0 && heights[stk[len(stk)-1]] < heights[i] {
			stk = stk[:len(stk)-1]
			cnt++
		}
		if len(stk) > 0 {
			cnt++
		}
		stk = append(stk, i)
		ans[i] = cnt
	}
	return ans
}
