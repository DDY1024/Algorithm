package main

import "sort"

// 题目链接: https://leetcode.cn/problems/rectangle-area-ii/
// 解题思路
// 		运用扫描线的思路来求解本题

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

type Interval struct {
	l, r int
}

func rectangleArea(rectangles [][]int) (ans int) {
	n, mod := len(rectangles), int(1e9+7)
	xarr := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		xarr = append(xarr, rectangles[i][0])
		xarr = append(xarr, rectangles[i][2])
	}
	sort.Ints(xarr)

	for i := 1; i < len(xarr); i++ {
		if xarr[i] == xarr[i-1] {
			continue
		}

		ilist := make([]Interval, 0, n)
		for j := 0; j < n; j++ {
			if rectangles[j][0] <= xarr[i-1] && xarr[i] <= rectangles[j][2] {
				ilist = append(ilist, Interval{rectangles[j][1], rectangles[j][3]})
			}
		}
		if len(ilist) == 0 {
			continue
		}

		sort.Slice(ilist, func(i, j int) bool {
			if ilist[i].l == ilist[j].l {
				return ilist[i].r <= ilist[j].r
			}
			return ilist[i].l < ilist[j].l
		})

		l, r := ilist[0].l, ilist[0].r
		for j := 1; j < len(ilist); j++ {
			if ilist[j].l <= r {
				r = maxInt(r, ilist[j].r)
			} else {
				ans += (xarr[i] - xarr[i-1]) * (r - l) % mod
				ans %= mod
				l, r = ilist[j].l, ilist[j].r
			}
		}
		ans += (xarr[i] - xarr[i-1]) * (r - l) % mod
		ans %= mod
	}
	return ans
}
