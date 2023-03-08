package main

// 题目链接：https://leetcode.cn/problems/can-i-win/description/

// 1. 由于数字不能重复使用，需要用到状态压缩
// 2. 由于子集 i 固定以后，数字和也随之固定，因此对于数字和没必要多一维保存其状态
// 3. 博弈论中
//		必胜点可以经过转移到达必败点
//      必败点只能转移到必胜点
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	n, m := maxChoosableInteger, desiredTotal
	if n >= m {
		return true
	}
	if (n+1)*n/2 < m {
		return false
	}

	dp := make([]int, 1<<n)
	for i := 0; i < (1 << n); i++ {
		dp[i] = -1
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j >= m { // 无论 i 是否有剩余，面临当前局面，必败
			return 0
		}

		if dp[i] != -1 {
			return dp[i]
		}

		ret := 0
		for x := 0; x < n; x++ {
			// 必胜点【可以】到达必败点
			// 必败点【只能】到达必胜点
			if i&(1<<x) == 0 && dfs(i|(1<<x), j+x+1) == 0 { // j+x+1
				ret = 1
				break
			}
		}

		dp[i] = ret
		return ret
	}

	return dfs(0, 0) == 1
}
