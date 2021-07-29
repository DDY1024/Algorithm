package main

import "fmt"

//
//
// 欠题
// TODO: 后缀数组
// https://leetcode-cn.com/problems/longest-common-subpath/
//

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n, m := len(passingFees), len(edges)
	dp := make([][]int, maxTime+1)
	for i := 0; i <= maxTime; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 0x3f3f3f3f
		}
	}
	dp[0][0] = passingFees[0]
	// 按照 cost time 递增的方向递推
	for i := 1; i <= maxTime; i++ {
		for j := 0; j < m; j++ {
			u, v, w := edges[j][0], edges[j][1], edges[j][2]
			if i >= w {
				dp[i][u] = minInt(dp[i][u], dp[i-w][v]+passingFees[u])
				dp[i][v] = minInt(dp[i][v], dp[i-w][u]+passingFees[v])
			}
		}
	}

	ans := 0x3f3f3f3f
	for i := 0; i <= maxTime; i++ {
		ans = minInt(ans, dp[i][n-1])
	}
	if ans >= 0x3f3f3f3f {
		ans = -1
	}
	return ans
}

func main() {
	fmt.Println("hello, world!")
}
