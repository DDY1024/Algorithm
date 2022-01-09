package main

import "sort"

// 题目链接: https://leetcode-cn.com/problems/earliest-possible-day-of-full-bloom/
// 题目大意: 怎样安排种花的顺序使得最终所有花开的时间最小
//
// 贪心问题 --> 排序算法解决
// 假设对于两棵植物分别是 p1, g1, p2, g2
// 植物 1 优先于植物 2 种植的条件是
// max{p1+g1, p1+p2+g2} <= max{p2+g2, p2+p1+g1}
// 所以我们完全可以通过排序算法决定出每棵植物的种植顺序
// 思路参考: https://leetcode-cn.com/problems/earliest-possible-day-of-full-bloom/solution/tan-xin-ji-qi-zheng-ming-by-endlesscheng-hfwe/

type Node struct {
	p, g int
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	nds := make([]Node, 0, n)
	for i := 0; i < n; i++ {
		nds = append(nds, Node{plantTime[i], growTime[i]})
	}

	sort.Slice(nds, func(i, j int) bool {
		return maxInt(nds[i].p+nds[i].g, nds[i].p+nds[j].p+nds[j].g) <= maxInt(nds[j].p+nds[j].g, nds[j].p+nds[i].p+nds[i].g)
	})

	ret, cost := 0, 0
	for i := 0; i < n; i++ {
		cost += nds[i].p
		ret = maxInt(ret, cost+nds[i].g)
	}
	return ret
}
