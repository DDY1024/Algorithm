package main

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp[i]: 以 arr[i] 为结尾的最大连续子段和
func maxSubSum(arr []int) int {
	n := len(arr)
	dp := make([]int, n)
	dp[0] = arr[0]
	ans := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = maxInt(arr[i], dp[i-1]+arr[i])
		ans = maxInt(ans, dp[i])
	}
	return ans
}

// 最大 M 子段和
// 从结果推导：最后一段肯定是以某个 arr[i] 为结尾的
// 状态表示：dp[i][j] 表示前 i 个数分成 j 段，以 arr[i] 为结尾的最大值
//
// 状态转移
// dp[i][j] = dp[i-1][j] + arr[i]  --> arr[i] 与 arr[i-1] 同为一个子段（共享一段）
// dp[i][j] = max{dp[t][j-1] + arr[i]}, [j-1,i) --> arr[i] 为新的一个子段（独自一段）

func maxMSubSum(arr []int, m int) int {
	n := len(arr)
	inf := 0x3f3f3f3f

	// 滚动数组 --> 减少空间占用
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= m; i++ {
		dp[i&1][i-1] = -inf
		maxPre := dp[(i-1)&1][i-1]
		for j := i; j <= n; j++ {
			dp[i&1][j] = dp[i&1][j-1] + arr[j-1]
			dp[i&1][j] = maxInt(dp[i&1][j], maxPre+arr[j-1])
			maxPre = maxInt(maxPre, dp[(i-1)&1][j]) // maxPre 在求解中维护，O(1) 复杂度进行状态转移
		}
	}

	ans := -inf
	// 最终结果 max{dp[i][m]}, i >= m
	for i := m; i <= n; i++ {
		ans = maxInt(ans, dp[m&1][i])
	}
	return ans
}
