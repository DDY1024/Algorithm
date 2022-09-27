package main

import "sort"

// 题目链接: https://leetcode.cn/problems/number-of-good-paths/
//
// 解题思路: https://leetcode.cn/problems/number-of-good-paths/solution/bing-cha-ji-by-endlesscheng-tbz8/
//
//
// 1. 当涉及到最大/最小的约束时，往往要按照节点值的顺序思考
// 2. 应该从大到小，还是从小到大呢？

func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	if n == 1 {
		return 1
	}

	adj := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		u, v := edges[i][0], edges[i][1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 逆向并查集，按照节点值从小到大进行加边，合并连通分量，并进行统计
	// 两个连通分量一旦合并，则两个连通分量内所有节点值最大的顶点之间的路径一定是好路径
	// 且该类统计方案是不存在重复的

	var initSet = func(parent []int, n int) {
		for i := 0; i < n; i++ {
			parent[i] = -1
		}
	}

	var findSet func(parent []int, x int) int
	findSet = func(parent []int, x int) int {
		if parent[x] < 0 {
			return x
		}

		parent[x] = findSet(parent, parent[x])
		return parent[x]
	}

	// 初始化并查集
	parent := make([]int, n)
	initSet(parent, n)

	// 顶点按照节点值从小到大进行排序，不断地进行合并
	ids := make([]int, n)
	sarr := make([]int, n) // 标识，每个连通块内，节点值最大的顶点个数
	for i := 0; i < n; i++ {
		ids[i] = i
		sarr[i] = 1
	}
	sort.Slice(ids, func(i, j int) bool {
		return vals[ids[i]] < vals[ids[j]]
	})

	ret := n // 单节点直接统计
	for _, u := range ids {
		fu := findSet(parent, u)
		for _, v := range adj[u] {
			fv := findSet(parent, v)
			if fu == fv || vals[v] > vals[u] { // 同连通块内 或 节点值更大，不处理
				continue
			}

			// 连通块内节点最大值相同，则需要累计最终方案数，并更新连通块内最大值节点的个数
			if vals[fu] == vals[fv] {
				ret += sarr[fu] * sarr[fv] // 两个连通块内最大值节点间的路径肯定是好路径
				sarr[fu] += sarr[fv]       // 连通块内最大值节点个数求和
			}
			parent[fv] = fu
		}
	}
	return ret
}
