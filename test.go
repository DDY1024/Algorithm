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

func minSkips(dist []int, speed int, hoursBefore int) int {
	n, totalDist := len(dist), 0
	for i := 0; i < n; i++ {
		totalDist += dist[i]
	}
	if totalDist > hoursBefore*speed {
		return -1
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n) // 最多只需要 n-1 跳过，最后一段路结束后不需要跳过
		for j := 0; j < n; j++ {
			dp[i][j] = 0x3f3f3f3f3f3f3f3f
		}
	}
	dp[0][0] = 0

	// x/y 整数向上取整，注意 x = 0 时，结果为 0
	var calc = func(x, y int) int {
		if x == 0 {
			return 0
		}
		return ((x-1)/y + 1) * y
	}

	for i := 1; i <= n; i++ {
		dp[i][0] = calc(dp[i-1][0], speed) + dist[i-1] // 注意: dp[1][0] = 0 + dist[0]
		for j := 1; j < i; j++ {
			dp[i][j] = minInt(
				calc(dp[i-1][j], speed)+dist[i-1],
				dp[i-1][j-1]+dist[i-1],
			)
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if dp[n][i] <= hoursBefore*speed {
			ans = i
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(math.Ceil(8.0 + 1.0/3 + 1.0/3 + 1.0/3))
}
