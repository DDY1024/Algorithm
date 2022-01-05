package main

// 题目链接: https://leetcode-cn.com/problems/unique-binary-search-trees/
// 求解不同二叉搜索树的个数
//
// 正规解法: 卡特兰数，由于此题数据范围比较小，直接采用递推方法求解
//

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
