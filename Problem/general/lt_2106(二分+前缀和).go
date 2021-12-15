package main

import "sort"

// 题目链接: https://leetcode-cn.com/problems/maximum-fruits-harvested-after-at-most-k-steps/submissions/
// 解题思路:
// 最优解存在以下几种情况
// 1. 先左，后右
// 2. 一直左
// 3. 先右后左
// 4. 一直右
// 那我们完全可以从 startPos 开始枚举所有可能的情况，结合二分查找定位转向后最远可到达的位置，总时间复杂度为 O(nlogn)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Fruit struct {
	pos int
	val int
}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)

	fl := make([]*Fruit, 0, n+1)
	for i := 0; i < n; i++ {
		fl = append(fl, &Fruit{fruits[i][0], fruits[i][1]})
	}
	fl = append(fl, &Fruit{startPos, 0})
	sort.Slice(fl, func(i, j int) bool {
		if fl[i].pos == fl[j].pos {
			return fl[i].val <= fl[j].val
		}
		return fl[i].pos < fl[j].pos
	})

	n++
	pSum := make([]int, n)
	idx := -1
	for i := 0; i < n; i++ {
		if idx == -1 && fl[i].pos == startPos {
			idx = i
		}
		pSum[i] = fl[i].val
		if i-1 >= 0 {
			pSum[i] += pSum[i-1]
		}
	}

	var calc = func(i, j int) int {
		ret := pSum[j]
		if i-1 >= 0 {
			ret -= pSum[i-1]
		}
		return ret
	}

	var searchRight = func(idx int, left int) int {
		if left < 0 {
			return 0
		}
		l, r, ans := idx, n-1, 0
		for l <= r {
			mid := l + (r-l)/2
			if fl[mid].pos-fl[idx].pos <= left {
				ans = calc(idx, mid)
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return ans
	}

	var searchLeft = func(idx, left int) int {
		if left < 0 {
			return 0
		}
		l, r, ans := 0, idx, 0
		for l <= r {
			mid := l + (r-l)/2
			if fl[idx].pos-fl[mid].pos <= left {
				ans = calc(mid, idx)
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return ans
	}

	// 向左走
	ans, left, get := 0, k, 0
	for i := idx - 1; i >= 0; i-- {
		left -= (fl[i+1].pos - fl[i].pos)
		if left < 0 {
			break
		}
		get += fl[i].val
		ans = maxInt(ans, get+searchRight(idx, left-(fl[idx].pos-fl[i].pos)))
	}

	// 向右走
	left, get = k, 0
	for i := idx + 1; i < n; i++ {
		left -= (fl[i].pos - fl[i-1].pos)
		if left < 0 {
			break
		}
		get += fl[i].val
		ans = maxInt(ans, get+searchLeft(idx, left-(fl[i].pos-fl[idx].pos)))
	}
	return ans
}
