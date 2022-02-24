package hard2022

// 题目链接: https://leetcode-cn.com/problems/where-will-the-ball-fall/
// 解题思路
// 1. 模拟小球掉落的过程
// 1 -> 1 向右
// -1 --> -1 向左
// 否则在同一行会形成 v 型，无法掉落
//
// 仔细观察形成 v 形结构的条件
//

func findBall(grid [][]int) []int {
	n, m := len(grid), len(grid[0])
	ans := make([]int, m)
	for j := 0; j < m; j++ {
		pos := j
		for i := 0; i < n; i++ {
			dir := grid[i][pos] // 移动方向，向左或向右
			pos += dir
			if pos < 0 || pos >= m || grid[i][pos] != dir { // 出界或形成v形结构
				pos = -1
				break
			}
		}
		ans[j] = pos
	}
	return ans
}
