package main

// 题目链接：https://leetcode.cn/problems/maximum-rows-covered-by-columns/?envType=daily-question&envId=2024-01-04
//
// Gosper's Hack 算法优化 n 元集合 k 元子集的生成

func maximumRows(matrix [][]int, numSelect int) int {
	n, m, ret := len(matrix), len(matrix[0]), 0
	rows := make([]int, n)
	for i := 0; i < n; i++ {
		mask := 0
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				mask |= 1 << j
			}
		}
		rows[i] = mask
	}

	cur, limit := (1<<numSelect)-1, 1<<m
	for cur < limit {
		cnt := 0
		for i := 0; i < n; i++ {
			if rows[i]&cur == rows[i] {
				cnt++
			}
		}
		ret = max(ret, cnt)

		lowbit := cur & (-cur)
		high := cur + lowbit
		cur = (((high ^ cur) >> 2) / lowbit) | high
	}
	return ret
}
