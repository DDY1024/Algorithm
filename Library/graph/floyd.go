package main

const (
	inf = 0x3f3f3f3f3f3f3f3f
)

// dist[i][j]：顶点 i --> j 的最短路径
// path[i][j]：顶点 i --> j 最短路径上 j 的前继节点
func floyd(n int, g [][]int) {
	dist := make([][]int, n)
	path := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		path[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dist[i][j] = g[i][j]
			if g[i][j] >= inf {
				path[i][j] = -1
			} else {
				path[i][j] = i
			}
		}
	}

	// i --> j 依次经过 0, 1, ..., n-1 顶点时的最短路径
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					path[i][j] = path[k][j]
				}
			}
		}
	}

	// Tips：最短路径可借助 path[i][j] 进行逆序输出
}

// 无向图最小环
// 		https://blog.csdn.net/AAMahone/article/details/90347573
//
//	枚举最小环上的最大顶点 u 及其邻接点 i(i<u), j(j<u)，再加上 i --> j 经过 0, ..., u-1 顶点的最短路径
//
//	因此可以看出，最小环的计算，可以利用 floyd 算法递推计算的过程进行求解

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCycle(n int, g [][]int) int {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = g[i][j]
		}
	}

	ret := inf
	for k := 0; k < n; k++ {
		// 更新最小环的值
		for i := 0; i < k; i++ {
			for j := i + 1; j < k; j++ {
				ret = minInt(ret, g[k][i]+g[k][j]+dist[i][j])
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = minInt(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	return ret
}

// 无向图的传递闭包
func closure(n int, g [][]int) {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				g[i][j] |= g[i][k] & g[k][j]
			}
		}
	}
}
