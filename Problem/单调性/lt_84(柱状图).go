package main

// 题目链接：https://leetcode.cn/problems/largest-rectangle-in-histogram/?envType=study-plan-v2&envId=top-100-liked
//
// 1. 利用单调栈预处理出【左边界】和【右边界】

func largestRectangleArea(heights []int) int {
	n := len(heights)
	lp, rp := make([]int, n), make([]int, n)

	var get = func(i int) int {
		if i < 0 || i >= n {
			return -1
		}
		return heights[i]
	}

	stk := []int{-1} // 方便处理
	for i := 0; i < n; i++ {
		for len(stk) > 0 && get(i) <= get(stk[len(stk)-1]) {
			stk = stk[:len(stk)-1]
		}

		lp[i] = stk[len(stk)-1] + 1
		stk = append(stk, i)
	}

	stk = []int{n} // 方便处理
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && get(i) <= get(stk[len(stk)-1]) {
			stk = stk[:len(stk)-1]
		}

		rp[i] = stk[len(stk)-1] - 1
		stk = append(stk, i)
	}

	ret := 0
	for i := 0; i < n; i++ {
		ret = max(ret, (rp[i]-lp[i]+1)*heights[i])
	}
	return ret
}
