package main

// 题目链接：https://leetcode.cn/problems/stone-game-iii/

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

// dp[i] 表示从第 i 堆开始取，先手比后手最多可以多取多少石头
func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n+1)
	dp[n], dp[n-1] = 0, stoneValue[n-1]
	for i := n - 2; i >= 0; i-- {
		dp[i] = -0x3f3f3f3f
		for j, s := 1, 0; j < 4 && i+j <= n; j++ {
			s += stoneValue[i+j-1]
			dp[i] = maxInt(dp[i], s-dp[i+j])
		}
	}
	if dp[0] > 0 {
		return "Alice"
	}
	if dp[0] == 0 {
		return "Tie"
	}
	return "Bob"
}
