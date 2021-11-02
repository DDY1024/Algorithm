package main

import (
	"sort"
)

func kthDistinct(arr []string, k int) string {
	// sort.Strings()
	cnt := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		cnt[arr[i]]++
	}

	ans, kth := "", 0
	for i := 0; i < len(arr); i++ {
		if cnt[arr[i]] == 1 && kth < k {
			ans = arr[i]
			kth++
		}
	}
	if kth < k {
		return ""
	}
	return ans
}

type AC struct {
	s, e, v int
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxTwoEvents(events [][]int) int {
	n := len(events)
	acList := make([]AC, 0, n)
	for i := 0; i < n; i++ {
		acList = append(acList, AC{events[i][0], events[i][1], events[i][2]})
	}
	sort.Slice(acList, func(i, j int) bool {
		if acList[i].e == acList[j].e {
			return acList[i].s < acList[j].s
		}
		return acList[i].e < acList[j].e
	})

	maxV := make([]int, n)
	maxV[0] = acList[0].v
	for i := 1; i < n; i++ {
		maxV[i] = maxInt(maxV[i-1], acList[i].v)
	}

	ans := acList[0].v
	for i := 0; i < n; i++ {
		l, r, tans := 0, n-1, -1
		for l <= r {
			mid := l + (r-l)/2
			if acList[mid].e < events[i][0] {
				tans = maxInt(tans, maxV[mid])
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		ans = maxInt(ans, events[i][2])
		ans = maxInt(ans, tans+events[i][2])
	}
	return ans
}

func platesBetweenCandles(s string, queries [][]int) []int {

	n := len(s)
	pSum := make([]int, n)
	lPos := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if s[i] == '|' {
			lPos = append(lPos, i)
		}
		if s[i] == '*' {
			pSum[i]++
		}
		if i-1 >= 0 {
			pSum[i] += pSum[i-1]
		}
	}

	var minInt = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var calc = func(i, j int, sum []int) int {
		ret := sum[j]
		if i > 0 {
			ret -= sum[i-1]
		}
		return ret
	}

	// >= x 最小值
	var b1 = func(x int) int {
		l, r, ans := 0, len(lPos)-1, n
		for l <= r {
			mid := l + (r-l)/2
			if lPos[mid] >= x {
				ans = minInt(ans, lPos[mid])
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		return ans
	}

	// <= x 最大值
	var b2 = func(x int) int {
		l, r, ans := 0, len(lPos)-1, 0
		for l <= r {
			mid := l + (r-l)/2
			if lPos[mid] <= x {
				ans = maxInt(ans, lPos[mid])
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return ans
	}

	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		s, e := queries[i][0], queries[i][1]
		l, r := b1(s), b2(e)
		if l < r {
			ans[i] = calc(l, r, pSum)
		}
	}

	return ans
}

func countCombinations(pieces []string, positions [][]int) int {
	n := len(pieces)
	for i := 0; i < n; i++ {
		positions[i][0]--
		positions[i][1]--
	}

	dx := []int{-1, 1, 0, 0, 1, 1, -1, -1}
	dy := []int{0, 0, -1, 1, 1, -1, 1, -1}
	dir := make([]int, n)
	step := make([]int, n)
	curPos := make([][]int, n)
	tstep := make([]int, n)
	for i := 0; i < n; i++ {
		curPos[i] = make([]int, 2)
	}

	var isValid = func() int {
		for i := 0; i < n; i++ {
			tstep[i] = step[i]
			curPos[i][0], curPos[i][1] = positions[i][0], positions[i][1]
		}
		mark := [8][8]int{}
		for {
			isMoved := false
			for i := 0; i < n; i++ {
				if tstep[i] > 0 {
					isMoved = true
					tstep[i]--
					curPos[i][0] += dx[dir[i]]
					curPos[i][1] += dy[dir[i]]
				}
				mark[curPos[i][0]][curPos[i][1]]++
			}
			if !isMoved { // 移动过程中没有发生重叠
				return 1
			}

			for i := 0; i < n; i++ {
				if mark[curPos[i][0]][curPos[i][1]] > 1 {
					return 0
				}
				mark[curPos[i][0]][curPos[i][1]] = 0
			}
		}
	}

	var calc func(idx int) int
	calc = func(idx int) int {
		if idx >= n {
			return isValid()
		}
		var l, r int
		switch pieces[idx][0] {
		case 'r':
			l, r = 0, 3
		case 'q':
			l, r = 0, 7
		case 'b':
			l, r = 4, 7
		}

		ans := 0
		x, y := positions[idx][0], positions[idx][1]
		for i := l; i <= r; i++ {
			for j := 1; j <= 8; j++ {
				if x+j*dx[i] >= 0 && x+j*dx[i] < 8 && y+j*dy[i] >= 0 && y+j*dy[i] < 8 {
					dir[idx] = i
					step[idx] = j
					ans += calc(idx + 1)
				}
			}
		}

		// 不移动随便选择一个方向即可
		dir[idx] = 0
		step[idx] = 0
		ans += calc(idx + 1)
		return ans
	}

	return calc(0)
}
