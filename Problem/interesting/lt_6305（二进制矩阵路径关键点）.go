package main

// 题目链接：https://leetcode.cn/problems/disconnect-path-in-a-binary-matrix-by-at-most-one-flip/description/
//
// 解题思路
// 		相当于找出左上角到右下角的路径中的关键节点，去掉该关键节点，则路径不连通
//
// 综合整理解题区各位大神的一些思路点子
//
// 1. 组合计数验证法
// 		f[i][j]: 表示从 (0, 0) 出发到达 (x, y) 路径方案数
//      g[i][j]: 表示从 (m-1, n-1) 出发到达 (x, y) 路径方案数
//  对于路径上的关键节点，必然满足
// 		f[x][y] * g[x][y] = f[m-1][n-1]
//  因为，所有可达路径必然经过该点，所以上述等式必然成立；但是组合计数存在一个问题，方案数太大，容易整数溢出；且利用
//  哈希取模方法，存在冲突的可能性

// 2. 上下轮廓相交判定法
//  不得不佩服灵神的思维能力呀
// 		https://leetcode.cn/problems/disconnect-path-in-a-binary-matrix-by-at-most-one-flip/solutions/2093243/zhuan-huan-cheng-qiu-lun-kuo-shi-fou-xia-io8x/
//  首先，我们知道题目要求只能向右或向下走动；那我们不妨采用两种策略进行走动
// 		a. 优先考虑向下走动，直到不能向下再考虑向右（下轮廓）
// 		b. 优先考虑向右走动，直到不能向右再考虑向下（上轮廓）
//  由上述两种走动策略可知，其可达路径经过的节点是尽可能分离的；如果这样两条路径存在交集，那这些交集节点必然是关键点
//	具体实现时，由于我们优先采用 a 策略，这样在去除 a 策略下的中间点后，只需要验证剩下那些点是否可达即可，并不需要按照 b 策略寻找

func isPossibleToCutPath(grid [][]int) bool {
	n, m := len(grid), len(grid[0])
	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if x == n-1 && y == m-1 {
			return true
		}
		grid[x][y] = 0 // 先判断后修改，确保 (n-1,m-1) 永远不会被修改
		return (x+1 < n && grid[x+1][y] == 1 && dfs(x+1, y)) ||
			(y+1 < m && grid[x][y+1] == 1 && dfs(x, y+1)) // 利用逻辑运算的短路特性
	}
	return !dfs(0, 0) || !dfs(0, 0)
}

// 3. 副对角线性质
// 		参考：https://leetcode.cn/problems/disconnect-path-in-a-binary-matrix-by-at-most-one-flip/solutions/2093392/zhao-bi-jing-zhi-lu-bu-yong-ren-he-shu-x-qc0z/
//	a. 剔除掉所有不可达点，只保留可达点（很重要，不可达点会干扰 b 的结论）
//  b. 存在一条副对角线上的可达点为 1 个，则这个点为关键点
func isPossibleToCutPathOther(grid [][]int) bool {
	n, m := len(grid), len(grid[0])
	mark := make([][]bool, n)
	for i := 0; i < n; i++ {
		mark[i] = make([]bool, m)
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		mark[x][y] = true
		if x-1 >= 0 && grid[x-1][y] == 1 && !mark[x-1][y] {
			dfs(x-1, y)
		}
		if y-1 >= 0 && grid[x][y-1] == 1 && !mark[x][y-1] {
			dfs(x, y-1)
		}
	}
	dfs(n-1, m-1)
	if !mark[0][0] {
		return true // 原图不连通，直接返回 true
	}

	stats := make([]int, n+m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mark[i][j] {
				stats[i+j]++ // 可达点副对角线统计
			}
		}
	}

	for i := 1; i < n+m-2; i++ {
		if stats[i] == 1 {
			return true
		}
	}
	return false
}
