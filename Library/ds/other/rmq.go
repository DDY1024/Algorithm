package other

import (
	"math"
)

// 1. 倍增思想（区间 dp）
// 2. 预处理，静态查询
// 3. 区间满足 "+" 性质，均可采用 RMQ 进行预处理
// 4. 预处理复杂度 O(n * logn)，查询时间复杂度 O(logn)

var ls = func(x int) int { return 1 << x }

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calcExp(n int) int {
	// 此处如果为了确保精度，可以直接循环计算出最大的 exp，满足 2^exp <= n
	return int(math.Floor(math.Log2(float64(n))))
}

func getMax(l, r int, dp [][]int) int {
	exp := calcExp(r - l + 1)
	return maxInt(dp[l][exp], dp[r-ls(exp)+1][exp])
}

func getMin(l, r int, dp [][]int) int {
	exp := calcExp(r - l + 1)
	return minInt(dp[l][exp], dp[r-ls(exp)+1][exp])
}

// dp[i][j]: 以 i 开始，区间长度为 1<<j 的最值
// dp[i][j] = OP{ dp[i][j-1], dp[i+ls(j-1)][j-1] }，其中 OP 表示任何区间可叠加类的操作
func initRMQ(n int, arr []int) [][]int {
	exp := calcExp(n)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, exp+1)
		dp[i][0] = arr[i] // 初始化边界条件
	}

	for l := 1; l <= exp; l++ {
		for i := 0; i+ls(l)-1 < n; i++ {
			dp[i][l] = maxInt(dp[i][l-1], dp[i+ls(l-1)][l-1])
		}
	}
	return dp
}
