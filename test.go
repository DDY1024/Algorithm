package main

import (
	"fmt"
	"math"
)

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// func minSkips(dist []int, speed int, hoursBefore int) int {
// 	n := len(dist)
// 	totalDist := 0
// 	for i := 0; i < n; i++ {
// 		totalDist += dist[i]
// 	}
// 	if float64(totalDist)/float64(speed) > float64(hoursBefore) {
// 		return -1
// 	}

// 	// dp := make([][]float64, n+1)
// 	// for i := 0; i <= n; i++ {
// 	// 	dp[i] = make([]float64, n)
// 	// 	for j := 0; j < n; i++ {
// 	// 		dp[i][j] = 0x3f3f3f3f
// 	// 	}
// 	// }
// 	// dp[0][0] = 0.0
// 	// for i := 1; i <= n; i++ {
// 	// 	dp[i][0] = dp[i-1][0] + math.Ceil(float64(dist[i])/float64(speed))
// 	// 	for j := 1; j < n; j++ {
// 	// 		// dp[i][j] = minFloat(dp[i-1][])
// 	// 	}
// 	// }

// }

func main() {
	fmt.Println(math.Ceil(8.0 + 1.0/3 + 1.0/3 + 1.0/3))
}
