package main

// 题目链接: https://leetcode-cn.com/problems/maximum-number-of-achievable-transfer-requests/
// 解题思路
// 由于题目数据范围，很容易想到采用二进制枚举+判定方式进行求解

func countBit(x int) int {
	cnt := 0
	for x > 0 {
		x &= (x - 1)
		cnt++
	}
	return cnt
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumRequests(n int, requests [][]int) int {
	m := len(requests)
	limit, ret := 1<<uint(m), 0
	for i := 1; i < limit; i++ {
		delta := make([]int, n)
		for j := 0; j < m; j++ {
			if i&(1<<uint(j)) > 0 {
				delta[requests[j][0]]--
				delta[requests[j][1]]++
			}
		}

		ok := true
		for j := 0; j < n; j++ {
			if delta[j] != 0 {
				ok = false
				break
			}
		}
		if ok {
			ret = maxInt(ret, countBit(i))
		}
	}
	return ret
}
