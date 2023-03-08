package main

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//  1. 最大 1 子段和
//     dp[i] 表示以 arr[i] 为结尾的最大连续子段和
func maxSubSum(arr []int) int {
	n := len(arr)
	dp := make([]int, n)

	dp[0] = arr[0]
	ret := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = maxInt(arr[i], dp[i-1]+arr[i])
		ret = maxInt(ret, dp[i])
	}
	return ret
}

// 2. 最大 M 子段和
// 		dp[i][j] 表示前 i 个数分成 j 段并以 arr[i] 结尾的最大值
//	状态转移方程
// 		dp[i][j] = dp[i-1][j] + arr[i] --> 表示 arr[i] 与 arr[i-1] 为同一段
// 		dp[i][j] = max{ dp[t][j-1] + arr[i] } --> 表示 arr[i] 单独成为一段

func maxMSubSum(arr []int, m int) int {
	n := len(arr)
	inf := 0x3f3f3f3f

	// 滚动数组
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= m; i++ {
		dp[i&1][i-1] = -inf
		maxP := dp[(i-1)&1][i-1]
		for j := i; j <= n; j++ {
			dp[i&1][j] = dp[i&1][j-1] + arr[j-1]           // 1. 共享一段
			dp[i&1][j] = maxInt(dp[i&1][j], maxP+arr[j-1]) // 2. 单独一段
			maxP = maxInt(maxP, dp[(i-1)&1][j])            // 优化实现 O(1) 状态转移
		}
	}

	ans := -inf
	for i := m; i <= n; i++ {
		ans = maxInt(ans, dp[m&1][i])
	}
	return ans
}
