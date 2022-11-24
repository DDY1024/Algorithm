package main

// 题目链接: https://leetcode.cn/problems/soup-servings/description/
//
// 解题思路: https://leetcode.cn/problems/soup-servings/solutions/1981704/fen-tang-by-leetcode-solution-0yxs/
//
// 技巧总结
// 1. 像这类求解概率类的题目，如果没有什么好的想法，不妨直接试试概率 dp 的求解思路
// 2. 由于题目 n 给的数据范围很大，会让人怀疑概率 dp 是否能够搞定
// 		通过每次分配去掉 A 和 B 数学期望的大小比较并结合精度误差的要求，在 n 较大时，直接返回概率 1.0
// 		(这类结论发现往往需要通过数据测算)
// 3. 由于题目数据都是 25 操作数据都是 25 的倍数，我们可以直接除 25 来降低求解范围，降低求解复杂度

func soupServings(n int) float64 {
	n = (n + 24) / 25 // 当 n % 25 > 0 时，为确保结果正确性需要 + 1
	if n >= 179 {     // 当 n >= 179 * 25 时，所求概率已经 >= 0.99999，满足题目精度 10^(-5) 要求
		return 1.0
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, n+1)
	}

	for i := 1; i <= n; i++ {
		dp[0][i] = 1.0
	}
	dp[0][0] = 1.0 / 2 // A、B 同时分配完的概率除以 2

	// 一共存在四种转移策略
	// 1: (4, 0)
	// 2: (3, 1)
	// 3: (2, 2)
	// 4: (1, 3)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = 0.25 * (dp[max(0, i-4)][j] + dp[max(0, i-3)][max(0, j-1)] +
				dp[max(0, i-2)][max(0, j-2)] + dp[max(0, i-1)][max(0, j-3)])
		}
	}
	return dp[n][n]
}
