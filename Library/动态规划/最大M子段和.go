package main

// dp[i]: 以 arr[i] 结尾的最大子段和
// ans = max(dp[i]) 0 <= i < n
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxSubSum(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	// ans, cur := arr[0], arr[0]
	// for i := 1; i < n; i++ {
	// 	if cur < 0 {
	// 		cur = 0
	// 	}
	// 	cur += arr[i]
	// 	ans = maxInt(ans, cur)
	// }

	var ans int
	dp := make([]int, n)
	dp[0] = arr[0]
	ans = dp[0]
	for i := 1; i < n; i++ {
		dp[i] = maxInt(arr[i], dp[i-1]+arr[i])
		ans = maxInt(ans, dp[i])
	}
	return ans
}

// 最大 M 子段和
// 最终结论：最后一段肯定是以某个 arr[i] 为结尾的，用于引导我们设计状态&&转移方程
// 状态表示：dp[i][j] 表示前 i 个数分成 j 段且最后一段包含 arr[i] 的最大值
// 状态转移方程
// dp[i][j] = dp[i-1][j] + arr[i]  ---> arr[i] 与 arr[i-1] 融为一个子段
// dp[i][j] = max(dp[t][j-1]+arr[i]), j-1 <= t < i ---> arr[i] 自成一段
// Tips: 状态转移方程优化 dp[t][j-1] 是可以在求解的过程中递推的；另外 j 阶段的状态基本全由 j-1 阶段状态推导而来，因此可以采用滚动数组进行优化。
func MaxMSubSum(arr []int, m int) int {
	n := len(arr)
	inf := 0x3f3f3f3f
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= m; i++ {
		dp[i&1][i-1] = -inf // 初始化赋值注意下
		maxPre := dp[(i-1)&1][i-1]
		for j := i; j <= n; j++ {
			dp[i&1][j] = dp[i&1][j-1] + arr[j-1]
			dp[i&1][j] = maxInt(dp[i&1][j], maxPre+arr[j-1])
			maxPre = maxInt(maxPre, dp[(i-1)&1][j])
		}
	}

	ans := -inf
	for i := m; i <= n; i++ {
		ans = maxInt(ans, dp[m&1][i])
	}
	return ans
}
