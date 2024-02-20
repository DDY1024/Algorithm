package main

import "sort"

// 题目链接：https://leetcode.cn/problems/palindrome-rearrangement-queries/description/

func canMakePalindromeQueries(s string, queries [][]int) []bool {
	return nil
}

func maximumLength(s string) int {
	n := len(s)
	stats, cnt := make([][]int, 26), 0
	for i := 0; i < n; i++ {
		if i-1 >= 0 && s[i-1] != s[i] {
			j := int(s[i-1] - 'a')
			stats[j] = append(stats[j], cnt)
			if cnt > 2 && maxL < cnt-2 {
				maxL = cnt - 2
			}
			cnt = 1
		} else {
			cnt++
		}
	}
	j := int(s[n-1] - 'a')
	stats[j] = append(stats[j], cnt)
	if cnt > 2 && maxL < cnt-2 {
		maxL = cnt - 2
	}

	for i := 0; i < 26; i++ {
		sort.Ints(stats[i])
		m := len(stats[i])
		if m >= 2 {
			// 不相等
			if stats[i][m-1] != stats[i][m-2] {
				if maxL < stats[i][m-2] {
					maxL = stats[i][m-2]
				}
			} else {
				// 相等
				if maxL < stats[i][m-1]-1 {
					maxL = stats[i][m-1] - 1
				}
			}
		}
		if m >= 3 && maxL < stats[i][m-3] {
			maxL = stats[i][m-3]
		}
	}
	if maxL <= 0 {
		maxL = -1
	}
	return maxL
}
