package graph

import "fmt"

// https://oi-wiki.org/graph/tree-diameter/
//
// 树的直径：树上任意两点之间最长的简单路径
//
// 重要结论
//		1. 一棵树可以存在【多条】直径。并且长度相等
//		2. 如果树的【边权】均为【正整数】，则所有直径的【中点】重合
//
// 两次搜索算法【边权非负】
// 		a. 从任意一点 u 出发，按照 dfs 或 bfs 搜索离 u 最远的顶点 v
//      b. 从顶点 v 出发，按照 dfs 或 bfs 搜索离 v 最远的顶点 w
//      c. 顶点 v 到 w 的简单路径长度，即为树的直径
//      d. 顶点 v 到 w 路径的中点（一个或两个）即为树的中心
// 注意，如果树中存在【负权边】，则【两次搜索】算法的结论不成立
//
//

func solve() {
	var (
		// n               int
		g               [][]int // 邻接表
		d               []int
		x, source, sink int
	)

	var dfs func(u, p int)
	dfs = func(u, p int) {
		if d[u] > d[x] {
			x = u
		}
		for i := 0; i < len(g[u]); i++ {
			v := g[u][i]
			if v == p {
				continue
			}
			d[v] = d[u] + 1
			dfs(v, u)
		}
	}

	x = 0
	dfs(x, -1) // 1. 第一次搜索
	source = x
	d[x] = 0
	dfs(x, -1) // 2. 第二次搜索
	sink = x
	// 树直径：source --> sink
	fmt.Println(source, sink)
}

//  2. 树形 DP
//     a. 树上最长路径的两个端点，必然是【叶子节点】或【根节点】  --> 【边权非负】
//     b. 针对【负权边】场景，仍然能够正确求解
//
//     状态定义
//     d1(i): 以 i 为根节点的子树，向下延伸到达的【最远】节点
//     d2(i): 以 i 为根节点的子树，向下延伸到达的【次远】节点
//
//     状态转移
//     d1(u) = max{d1(v)} + 1
//     d2(u) = second_max{d1(v)} + 1
//
//     树的直径 = max{ d1(u) + d2(u) }
//
//     如何求解树的直径上的节点?
//     a. 找到节点 u
//     b. 找到 d1(u) 的最大子节点
//     c. 找到 d2(u) 的最大子节点
//     d. 根据 2 和 3 中子节点的【向下最长路径】的节点构成
func solve2() int {
	var (
		n   int
		g   [][]int // 邻接表
		d1  = make([]int, n)
		d2  = make([]int, n)
		ret int
	)

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dfs func(u, p int)
	dfs = func(u, p int) {
		d1[u], d2[u] = 0, 0
		for i := 0; i < len(g[u]); i++ {
			v := g[u][i]
			if v == p {
				continue
			}
			dfs(v, u)
			d := d1[v] + 1 // 边权值均为 1
			if d > d1[u] {
				d2[u], d1[u] = d1[u], d
			} else if d > d2[u] {
				d2[u] = d
			}
		}
		ret = maxInt(ret, d1[u]+d2[u]) // 遍历树节点的过程中枚举完全部节点
	}
	dfs(0, -1)
	return ret
}

// https://leetcode-cn.com/problems/minimum-height-trees/
// 结论：以【树的直径】中【中间节点】为【根节点】生成的树高度【最小】

func findMinHeightTrees(n int, edges [][]int) []int {
	m := len(edges)
	adj := make([][]int, n)
	for i := 0; i < m; i++ {
		u, v := edges[i][0], edges[i][1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	d1 := make([]int, n)
	d2 := make([]int, n)
	c1 := make([]int, n) // 最远路径的孩子节点
	c2 := make([]int, n) // 次远路径的孩子节点
	maxD := 0

	var dfs func(u, p int)
	dfs = func(u, p int) {
		d1[u], d2[u], c1[u], c2[u] = 0, 0, -1, -1
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			dfs(v, u)
			tmp := d1[v] + 1
			if tmp > d1[u] {
				d2[u] = d1[u]
				c2[u] = c1[u]
				d1[u] = tmp
				c1[u] = v
			} else if tmp > d2[u] {
				d2[u] = tmp
				c2[u] = v
			}
		}
		maxD = maxInt(maxD, d1[u]+d2[u]+1)
	}
	dfs(0, -1)

	// 找出树直径上的节点
	path := make([]int, maxD)
	for i := 0; i < n; i++ {
		if d1[i]+d2[i]+1 == maxD {
			path[d1[i]] = i
			// 沿着最远路径进行寻找
			u, idx := c1[i], d1[i]-1
			for u != -1 {
				path[idx] = u
				idx--
				u = c1[u]
			}

			// 沿着次远路径进行寻找
			u, idx = c2[i], d1[i]+1
			for u != -1 {
				path[idx] = u
				idx++
				u = c1[u]
			}
			break
		}
	}

	// 寻找树的直径的中间节点
	if maxD&1 > 0 {
		return []int{path[maxD>>1]}
	}

	return []int{path[(maxD>>1)-1], path[maxD>>1]}
}
