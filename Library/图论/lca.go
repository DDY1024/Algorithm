package main

import (
	"fmt"
	"math"
)

var (
	n     int
	adj   [][]int // 邻接表
	in    []int   // 入编号
	out   []int   // 出编号
	vmark []int   // 顶点遍历序列
	dfn   int
)

// 第一种标记方法
//
// 1. 进编号
// 2. 出编号
func dfs1(u, p int) {
	dfn++
	in[u] = dfn
	vmark[dfn] = u

	for i := 0; i < len(adj[u]); i++ {
		v := adj[u][i]
		if v == p {
			continue
		}
		dfs1(v, u)
	}

	dfn++
	out[u] = dfn
	vmark[dfn] = u
}

// 树的欧拉序
var (
	first []int // 顶点在欧拉序中第一次出现的位置
	rmq   []int // 顶点深度
	dp    [][]int
)

// dfn 从 1 开始递增
// dfn 从 0 开始递增
func dfs2(u, p, d int) {
	dfn++
	first[u] = dfn
	vmark[dfn] = u
	rmq[dfn] = d

	for i := 0; i < len(adj[u]); i++ {
		v := adj[u][i]
		if v == p {
			continue
		}

		dfs2(v, u, d+1)

		dfn++
		vmark[dfn] = u
		rmq[dfn] = d
	}
}

func log2(n int) int {
	return int(math.Floor(math.Log2(float64(n))))
}

func initRMQ() {
	n, m := dfn+1, log2(dfn+1)
	dp = make([][]int, n)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = i
	}

	// 小区间  -->  大区间
	for l := 1; l <= m; l++ {
		for i := 0; i+(1<<l)-1 < n; i++ {
			id1 := dp[i][l-1]
			id2 := dp[i+(1<<(l-1))][l-1]
			if rmq[id1] <= rmq[id2] {
				dp[i][l] = id1
			} else {
				dp[i][l] = id2
			}
		}
	}
}

// LCA(u, v) = first[u] 和 first[v] 区间内【深度最小】的顶点
func queryRMQ(u, v int) int {
	l, r := first[u], first[v]
	if l > r {
		l, r = r, l
	}

	k := log2(r - l + 1)
	id1 := dp[l][k]
	id2 := dp[r-(1<<k)+1][k]
	if rmq[id1] <= rmq[id2] {
		return vmark[id1]
	}
	return vmark[id2]
}

// Tarjan LCA（离线算法）

type Query struct {
	v, idx int
}

var (
	parent   []int
	vis      []bool
	ancestor []int
	query    [][]Query
)

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

func tarjan(u int) {
	ancestor[u] = u
	vis[u] = true

	for i := 0; i < len(adj[u]); i++ {
		v := adj[u][i]
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

// 3. LCP 倍增算法
