package main

const (
	inf  = 0x3f3f3f3f
	maxn = 110
)

var (
	adj  [maxn][maxn]int // 邻接矩阵
	path [maxn][maxn]int // i -> j 最短路径中j的前驱节点
	dp   [maxn][maxn]int // i -> j 最短路径取值
)

// O(n^3)
func solve(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = adj[i][j]
			if adj[i][j] >= inf {
				path[i][j] = -1
			} else {
				path[i][j] = i
			}
		}
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if dp[i][j] > dp[i][k]+dp[k][j] {
					dp[i][j] = dp[i][k] + dp[k][j]
					path[i][j] = path[k][j]
				}
			}
		}
	}
}
