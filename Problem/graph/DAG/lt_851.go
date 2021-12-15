package main

// 题目链接: https://leetcode-cn.com/problems/loud-and-rich/
// 解题思路
// 1. DAG 上的动态规划
// 2. 记忆化搜索

func loudAndRich(richer [][]int, quiet []int) []int {
	n, m := len(quiet), len(richer)
	adj := make([][]int, n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, n)
		ans[i] = -1
	}

	for i := 0; i < m; i++ {
		u, v := richer[i][0], richer[i][1]
		adj[v] = append(adj[v], u)
	}

	// dfs: dag 上的动态规划，记忆化搜索
	var dfs func(u int)
	dfs = func(u int) {
		if ans[u] != -1 {
			return
		}
		ans[u] = u
		for _, v := range adj[u] {
			dfs(v)
			if quiet[ans[v]] < quiet[ans[u]] {
				ans[u] = ans[v]
			}
		}
	}

	// 非连通，需要全部结点都计算一遍
	for i := 0; i < n; i++ {
		dfs(i)
	}
	return ans
}
