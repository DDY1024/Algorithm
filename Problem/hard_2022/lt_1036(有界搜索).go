package main

// 题目链接: https://leetcode-cn.com/problems/escape-a-large-maze/
// 题目大意
// 在一个足够大的空间里，有少数的障碍物，问两点是否连通。
//
// 解题思路可以参考: https://leetcode-cn.com/problems/escape-a-large-maze/solution/gong-shui-san-xie-bfs-gei-ding-zhang-ai-8w63o/
// https://leetcode-cn.com/problems/escape-a-large-maze/solution/tao-chi-da-mi-gong-by-leetcode-solution-qxhz/
//
// 主要参考这两篇解题报告的求解上界的思，由 n 个砖块最大能够围成的封闭区域的面积（注意不是单纯的计算面积，而是数方格数）
// 1. 直边调整为斜边可以围成更大的面积
// 2. 借助边界可以围成更大的面积
// 最终我们得出一个上界的结论: 一个斜边三角形，其中包含的方格数为 (n-1)+...+1 = n*(n-1)/2
//
// 综上所述: 我们只需要判断
// 1. 源点 和 目标点 是否被包围。如果没有被包围，则必然存在一条连接路径。
// 2. 源点 和 目标点 是否在同一个包围圈中，可达。
//
// 带边界限制条件的搜索，仔细挖掘题目的性质

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	if len(blocked) < 2 {
		return true
	}

	n, base := len(blocked), 1000000

	var hash = func(x, y int) int {
		return x*base + y
	}

	isb := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		isb[hash(blocked[i][0], blocked[i][1])] = true
	}

	sx, sy, tx, ty, canTo, limit := source[0], source[1], target[0], target[1], false, n*(n-1)/2
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	var check = func(sx, sy, tx, ty int) bool {
		vis := make(map[int]bool, n*(n-1)/2)
		que := make([]int, 0, n*(n-1)/2+10)
		que = append(que, hash(sx, sy))
		vis[hash(sx, sy)] = true

		viscnt := 0
		for len(que) > 0 {
			pos := que[0]
			que = que[1:]

			x, y := pos/base, pos%base
			if x == tx && y == ty {
				canTo = true
				return true
			}

			// 超过限制，则没有被围
			viscnt++
			if viscnt > limit {
				return true
			}

			for i := 0; i < 4; i++ {
				xx, yy := x+dx[i], y+dy[i]
				if xx >= 0 && xx < base && yy >= 0 && yy < base && !isb[hash(xx, yy)] && !vis[hash(xx, yy)] {
					vis[hash(xx, yy)] = true
					que = append(que, hash(xx, yy))
				}
			}
		}
		return false
	}

	r1 := check(sx, sy, tx, ty)
	if canTo {
		return true
	}

	r2 := check(tx, ty, sx, sy)
	return r1 && r2
}

//
// [[10,9],[9,10],[10,11],[11,10]]
// [0,0]
// [10,10]
//
//
//         (9,10)
// (10,9)        (10,11)
//        (11,10)
//
