package main

// 题目链接: https://leetcode.cn/problems/selling-pieces-of-wood/
// 注意理解题目要求，否则容易想高复杂度
// 每一次操作中，你必须按下述方式之一执行切割操作，以得到两块更小的矩形木块：
//     a. 沿垂直方向按高度 完全 切割木块，或
//     b. 沿水平方向按宽度 完全 切割木块
// 所以对于一个方块 (h, w) 切割一次的结果为两个小方块且为 (h1,w)、(h2,w) 或 (h,w1)、(h,w2)
// 因此，我们完全可以枚举切割方案，利用 dp 来求解最优值；
// 另外，对于 (h, w) 刚好可以出售的情况，我们可以直接出售，或者切割成更小的出售
// 注意题目确保 (hi, wi) 唯一

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sellingWood(m int, n int, prices [][]int) int64 {
	pn := len(prices)
	pmap := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		pmap[i] = make([]int, n+1)
	}
	for i := 0; i < pn; i++ {
		pmap[prices[i][0]][prices[i][1]] = prices[i][2]
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = pmap[i][j] // 不存在，则为 0
			// 常数优化: 枚举一半即可
			for k := 1; k < i; k++ {
				dp[i][j] = maxInt(dp[i][j], dp[k][j]+dp[i-k][j])
			}
			for k := 1; k < j; k++ {
				dp[i][j] = maxInt(dp[i][j], dp[i][k]+dp[i][j-k])
			}
		}
	}

	return int64(dp[m][n])
}
