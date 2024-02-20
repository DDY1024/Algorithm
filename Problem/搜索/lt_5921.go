package main

// 题目链接：https://leetcode-cn.com/problems/maximum-path-quality-of-a-graph/
// 本题解题思路需要从数据范围中挖掘
// 1. 10 <= timej, maxTime <= 100 意味着最终结果最多经过 10 条边
// 2. 每个顶点最多有四条出边
// 3. 加上一定的剪枝策略，我们完全可以在较短时间内进行求解(最短路径剪枝，事实上完全没有必要)
type Edge struct {
	v, w int
}

func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n, m := len(values), len(edges)
	adj := make([][]*Edge, n)
	for i := 0; i < m; i++ {
		u, v, w := edges[i][0], edges[i][1], edges[i][2]
		adj[u] = append(adj[u], &Edge{v, w})
		adj[v] = append(adj[v], &Edge{u, w})
	}

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := 0
	path := make(map[int]int, 20)
	path[0]++
	var dfs func(u, cost int)
	dfs = func(u, cost int) {
		if u == 0 {
			tmp := 0
			for u, c := range path {
				if c > 0 {
					tmp += values[u]
				}
			}
			ans = maxInt(ans, tmp)
		}

		for _, e := range adj[u] {
			if cost+e.w > maxTime {
				continue
			}
			path[e.v]++
			dfs(e.v, cost+e.w)
			path[e.v]--
			if path[e.v] == 0 {
				delete(path, e.v)
			}
		}
	}
	dfs(0, 0)
	return ans
}
