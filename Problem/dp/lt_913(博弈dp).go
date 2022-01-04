package main

// 题目链接: https://leetcode-cn.com/problems/cat-and-mouse/
// 题目大意
// 0. 无向图，猫和老鼠每次必须选择一条边进行移动
// 1. 猫和老鼠游戏，猫起始位置为 2，老鼠起始位置为 1，洞位置为 0
// 2. 老鼠能够移动到 0 则获胜
//
// 解题思路  https://leetcode-cn.com/problems/cat-and-mouse/solution/mao-he-lao-shu-by-leetcode-solution-444x/
// 1. 必败状态: 所有操作均为对方的必胜状态 --> 猫和老鼠重合
// 2. 必胜状态: 存在一种操作到达对方的必败状态 --> 老鼠进洞
// 3. 必和状态: 非必胜状态，且存在一种操作到达对方的必和状态 -->  ">= 2n 轮仍然无法分出胜负" 此处比较难以证明，可以直接参考题解
//
// 解决方法: 动态规划、记忆化搜索

// 0: 平局 1: 必胜 2: 必败
// dp[i][j][k]: 进行到第 i 轮，老鼠位置 j，猫位置 k 时的状态
func catMouseGame(graph [][]int) int {
	n := len(graph)
	dp := make([][][]int, 2*n+1)
	for i := 0; i <= 2*n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
			for k := 0; k < n; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	// 获取转移状态
	var getNext = func(i, x, y int) []int {
		if i&1 > 0 { // 奇数轮
			return graph[y]
		}
		return graph[x] // 偶数轮
	}

	var solve func(i, x, y int) int
	solve = func(i, x, y int) int {
		if dp[i][x][y] != -1 {
			return dp[i][x][y]
		}

		// 为什么 >= 2n 轮肯定是平局???
		if i >= 2*n {
			dp[i][x][y] = 0
			return dp[i][x][y]
		}

		// 老鼠入洞
		if x == 0 {
			if i&1 > 0 {
				dp[i][x][y] = 2 // 必败
			} else {
				dp[i][x][y] = 1 // 必胜
			}
			return dp[i][x][y]
		}

		// 老鼠和猫相遇
		if x == y {
			if i&1 > 0 {
				dp[i][x][y] = 1 // 必胜
			} else {
				dp[i][x][y] = 2 // 必败
			}
			return dp[i][x][y]
		}

		ret, win, tie := 0, false, false
		for _, v := range getNext(i, x, y) {
			if i&1 > 0 && v == 0 { // 猫不能进入老鼠洞(位置0)
				continue
			}

			var tret int
			if i&1 > 0 {
				tret = solve(i+1, x, v)
			} else {
				tret = solve(i+1, v, y)
			}

			if tret == 0 { // 存在平局状态 --> 次优选择
				tie = true
			}

			if tret == 2 { // 存在必败状态 --> 最优选择
				win = true
			}
		}

		if win {
			ret = 1
		} else if tie {
			ret = 0
		} else {
			ret = 2
		}

		dp[i][x][y] = ret
		return dp[i][x][y]
	}
	return solve(0, 1, 2)
}
