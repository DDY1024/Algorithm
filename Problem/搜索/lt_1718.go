package main

// 直接 dfs 搜索构造就 ok ??? 这也太 ok 了吧
// 首先，我们需要寻找可行解，在可行解的基础上，再去寻找最大解
// https://leetcode.com/submissions/detail/440935493/
func constructDistancedSequence(n int) []int {
	pos, m := make([]int, 2*n-1), 2*n-1
	mark := make([]bool, n+1)
	var dfs func(p int) bool
	dfs = func(p int) bool {
		if p >= m {
			return true
		}
		if pos[p] > 0 {
			return dfs(p + 1)
		}
		for i := n; i > 0; i-- {
			if mark[i] {
				continue
			}
			p2 := p + i
			if i == 1 {
				p2--
			}
			if p2 < m && pos[p2] == 0 {
				pos[p], pos[p2], mark[i] = i, i, true
				if dfs(p + 1) {
					return true
				}
				pos[p], pos[p2], mark[i] = 0, 0, false
			}
		}
		return false
	}
	dfs(0)
	return pos
}
