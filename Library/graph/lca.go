package main

import (
	"fmt"
	"math"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Edge struct {
	v, next int
}

var edges []Edge
var head []int
var enum int
var n, m int

func initGraph() {
	head = make([]int, n)
	for i := 0; i < n; i++ {
		head[i] = -1
	}
	enum = 0
	edges = make([]Edge, m)
}

func addEdge(u, v int) {
	edges[enum].v = v
	edges[enum].next = head[u]
	head[u] = enum
	enum++

	edges[enum].v = u
	edges[enum].next = head[v]
	head[v] = enum
	enum++
}

//  1. 根节点至当前节点的路径权值和
//     进入点赋【正权值】，出去点赋【负权】值；求解前缀和即为根节点到某一节点的路径和
//
//  2. in[u] 和 out[u] 之间的区间代表整棵子树
//     可以求解子树节点的权值和，也可以求解子树节点的个数等等
var in []int
var out []int
var vertex []int
var dfn int

// 第一种标记方法
func dfs1(u, p int) {
	dfn++
	in[u] = dfn // 进入点
	vertex[dfn] = u

	for i := head[u]; i != -1; i = edges[i].next {
		v := edges[i].v
		if v == p {
			continue
		}
		dfs1(v, u)
	}

	dfn++
	out[u] = dfn // 出去点
	vertex[dfn] = u
}

// 树的欧拉序参考：https://www.jianshu.com/p/55037ae618ce
var first []int // 记录该顶点在欧拉序中第一次出现的位置
var rmq []int   // 记录每个顶点的深度

// 第二种标记方法
func dfs2(u, p, d int) {

	dfn++
	vertex[dfn] = u // 进入点
	rmq[dfn] = d
	first[u] = dfn

	for i := head[u]; i != -1; i = edges[i].next {
		v := edges[i].v
		if v == p {
			continue
		}

		dfs2(v, u, d+1)

		// 欧拉序：出去点（回溯点）
		dfn++
		vertex[dfn] = u
		rmq[dfn] = d
	}
	return
}

// 欧拉序 + rmq 在线求解 lca
var dp [][]int

func initRMQ() {
	dp = make([][]int, 2*n)
	for i := 0; i < 2*n; i++ {
		dp[i] = make([]int, 20)
	}

	for i := 1; i < 2*n; i++ {
		dp[i][0] = i
	}

	for l := 1; 1<<uint(l) < 2*n; l++ {
		for i := 1; i+(1<<uint(l)) < 2*n; i++ {
			id1 := dp[i][l-1]
			id2 := dp[i+(1<<uint(l-1))][l-1]
			if rmq[id1] <= rmq[id2] {
				dp[i][l] = id1
			} else {
				dp[i][l] = id2
			}
		}
	}
}

// first[u] 和 first[v] 之间深度最小的顶点即为 lca(u,v)
func queryRMQ(u, v int) int {
	l := minInt(first[u], first[v])
	r := maxInt(first[u], first[v])
	st := int(math.Log2(float64(r - l + 1)))
	id1 := dp[l][st]
	id2 := dp[r-(1<<uint(st))+1][st]
	if rmq[id1] <= rmq[id2] {
		return vertex[id1]
	}
	return vertex[id2]
}

// 2. tarjan 算法求解 lca
var parent []int

func initSet() {
	parent = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
}

func find(u int) int {
	if parent[u] < 0 {
		return u
	}
	parent[u] = find(parent[u])
	return parent[u]
}

func union(u, v int) {
	r1, r2 := find(u), find(v)
	if r1 != r2 {
		parent[r1] = r2
	}
}

type Query struct {
	v, idx int
}

var vis []bool
var ancestor []int
var query [][]Query

func tarjan(u int) {
	ancestor[u] = u
	vis[u] = true

	for i := head[u]; i != -1; i = edges[i].next {
		v := edges[i].v
		if vis[v] {
			continue
		}

		tarjan(v)
		union(u, v)
		ancestor[find(u)] = u
	}

	// 离线查询处理
	for i := range query[u] {
		v := query[u][i].v
		if vis[v] {
			fmt.Println("LCA: ", u, v, ancestor[find(v)])
		}
	}
}

// var fa [][]int   // fa[i][j] 节点 i 的第 2^j 个祖先节点
// var vdepth []int // 节点 i 的深度

// func BFS(root, n int) {
// 	fa = make([][]int, n)
// 	vdepth = make([]int, n)
// 	vis = make([]bool, n)
// 	for i := 0; i < n; i++ {
// 		fa[i] = make([]int, 10)
// 	}
// 	que := make([]int, 0, n)
// 	que = append(que, root)
// 	fa[root][0] = root
// 	vdepth[root] = 0
// 	vis[root] = true
// 	for len(que) > 0 {
// 		u := que[0]
// 		que = que[1:]
// 		for i := 1; i < 10; i++ {
// 			fa[u][i] = fa[fa[u][i-1]][i-1]
// 		}
// 		for i := head[u]; i != -1; i = edges[i].next {
// 			v := edges[i].v
// 			if vis[v] {
// 				continue
// 			}
// 			vis[v] = true // 此处其实可以直接用父节点做判断
// 			vdepth[v] = vdepth[u] + 1
// 			fa[v][0] = u
// 			que = append(que, v)
// 		}
// 	}
// 	return
// }

// // 采用倍增法求解 lca
// // 自底向上一步步逼近
// func BLCA(u, v int) int {
// 	if vdepth[u] > vdepth[v] {
// 		vdepth[u], vdepth[v] = vdepth[v], vdepth[u]
// 	}
// 	for dh, i := vdepth[v]-vdepth[u], 0; dh > 0; dh, i = dh/2, i+1 {
// 		if dh&1 > 0 {
// 			v = fa[v][i]
// 		}
// 	}
// 	if u == v {
// 		return u
// 	}
// 	// fa[u][i] 到顶了，则 fa[u][i+1] 也是到顶的
// 	for i := 9; i >= 0; i-- { // 此处求解的方法比较巧妙
// 		if fa[u][i] == fa[v][i] {
// 			continue
// 		}
// 		u, v = fa[u][i], fa[v][i]
// 	}
// 	return fa[u][0]
// }
