package hard2022

// 题目链接: https://leetcode-cn.com/problems/grid-illumination/
// 网格中关于行、列、对角线的 Hash 表示

type Pair struct {
	x, y int
}

func gridIllumination(n int, lamps, queries [][]int) []int {
	rowMark := make(map[int]int)
	colMark := make(map[int]int)
	diagMark := make(map[int]int)
	rDiagMark := make(map[int]int)
	qMark := make(map[Pair]bool)
	for i := 0; i < len(lamps); i++ {
		x, y := lamps[i][0], lamps[i][1]
		if qMark[Pair{x, y}] { // 1. 存在重复的灯，需要进行去重
			continue
		}
		rowMark[x]++
		colMark[y]++
		diagMark[y-x]++
		rDiagMark[x+y]++
		qMark[Pair{x, y}] = true
	}

	dx, dy := []int{0, -1, 1, 0, 0, -1, -1, 1, 1}, []int{0, 0, 0, -1, 1, -1, 1, -1, 1}
	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		p := Pair{queries[i][0], queries[i][1]}
		if rowMark[p.x] > 0 || colMark[p.y] > 0 || diagMark[p.y-p.x] > 0 || rDiagMark[p.x+p.y] > 0 {
			ans[i] = 1
		}
		for j := 0; j < 9; j++ {
			xx, yy := p.x+dx[j], p.y+dy[j]
			// 2. 边界判断
			if xx < 0 || yy < 0 || xx >= n || yy >= n || !qMark[Pair{xx, yy}] {
				continue
			}
			delete(qMark, Pair{xx, yy})
			// qMark[Pair{xx, yy}] = false
			rowMark[xx]--
			colMark[yy]--
			diagMark[yy-xx]--
			rDiagMark[xx+yy]--
		}
	}
	return ans
}
