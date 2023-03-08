package main

const (
	inf  = 0x3f3f3f3f
	maxn = 110
)

var (
	adj  [maxn][maxn]int // 邻接矩阵
	path [maxn][maxn]int // i -> j 最短路径中 j 的前驱节点
	dist [maxn][maxn]int // i -> j 最短路径长度
)

func solve(n int) {

	// 1. 初始化
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dist[i][j] = adj[i][j]
			if adj[i][j] >= inf {
				path[i][j] = -1
			} else {
				path[i][j] = i
			}
		}
	}

	// 2. i -> j 依次经过 1,...,n 顶点时的最短路径
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					path[i][j] = path[k][j]
				}
			}
		}
	}

	// 逆序输出 path 路径上的节点
}

// 1. 最小环问题（无向图）
// https://blog.csdn.net/AAMahone/article/details/90347573
//		第 K-1 次迭代：表示经过 0, ..., k-1 中间点的最短路径
//		第 K 次迭代：表示经过 0, ..., k-1, k 中间点的最短路径
//
// 求解最小环，通过枚举最小环上的最大顶点 u 及其邻接点 i,j, 再加上 i -> j 经过 0,...,u-1 顶点的最短路径

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCycle(n int) {
	ret := 0x3f3f3f3f
	for k := 0; k < n; k++ {
		// 无向图中 (i,j) 和 (j,i) 对称相同
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				ret = minInt(ret, adj[k][i]+adj[k][j]+dist[i][j])
			}
		}

		// 正常 floyd 迭代过程
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = minInt(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
}

// 2. 求解图的传递闭包
/*
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				ok[i][j] |= ok[i][k]&ok[k][j]
			}
		}
	}
*/
