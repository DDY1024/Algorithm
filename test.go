package main

import "sort"

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

type Node struct {
	idx  int // 数组索引下标
	maxh int // 允许的最大高度
}

func maxBuilding(n int, restrictions [][]int) int {
	m := len(restrictions)
	nds := make([]Node, 0, m+1)
	nds = append(nds, Node{1, 0})
	for i := 0; i < m; i++ {
		nds = append(nds, Node{restrictions[i][0], restrictions[i][1]})
	}
	sort.Slice(nds, func(i, j int) bool {
		return nds[i].idx < nds[j].idx
	})

	limit, minH := make([]int, m+1), 0
	for i := 0; i <= m; i++ {
		limit[i] = nds[i].maxh
	}

	// dp[i]: 1 ~ i 满足高度限制条件下能够达到的最大高度，即
	// minInt(nds[0].maxh + nds[i].idx-nds[0].idx, nds[1].maxh + nds[i].idx-nds[1].idx, ...)
	// 容易想到下面的递推方程: dp[i] = minInt(dp[i-1] + (nds[i].idx - nds[i-1].idx, nds[i].maxh)
	for i := 1; i <= m; i++ {
		d := nds[i].idx - nds[i-1].idx
		limit[i] = minInt(limit[i], minH+d)
		minH = minInt(minH+d, nds[i].maxh)
	}

	// 同上进行反向递推
	minH = nds[m].maxh
	for i := m - 1; i >= 0; i-- {
		d := nds[i+1].idx - nds[i].idx
		limit[i] = minInt(limit[i], minH+d)
		minH = minInt(minH+d, nds[i].maxh)
	}

	ans := 0
	// 最高点可能出现在两个卡点 (i, i+1) 之间，至于 best(i, i+1) 需要满足如下条件:
	// best(i, i+1) - limit(i) + best(i, i+1) - limit(i+1) <= nds[i+1].idx - nds[i].idx
	// 2 * best(i, i+1) <= limit(i) + limit(i+1) - nds[i+1].idx - nds[i].idx
	for i := 0; i <= m; i++ {
		ans = maxInt(ans, limit[i])
		if i+1 <= m {
			ans = maxInt(ans, (limit[i]+limit[i+1]+nds[i+1].idx-nds[i].idx)/2)
		}
	}
	// 最后一个卡点不要忘记计算，直接累加
	ans = maxInt(ans, limit[m]+n-nds[m].idx)
	return ans
}

func main() {

}
