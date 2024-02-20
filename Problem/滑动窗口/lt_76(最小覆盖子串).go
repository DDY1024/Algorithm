package main

// 题目链接：https://leetcode.cn/problems/minimum-window-substring/description/?envType=study-plan-v2&envId=top-100-liked
//
// 1. 利用滑动窗口进行统计
//
//

func minWindow(s string, t string) string {
	n, m := len(s), len(t)
	c1, c2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < m; i++ {
		c2[t[i]]++
	}

	bs := []byte(s)
	start, end := 0, m+n
	mark := make(map[byte]bool)

	for i, j := 0, 0; i < n; i++ {
		c1[s[i]]++
		if _, ok := c2[s[i]]; ok && c1[s[i]] >= c2[s[i]] {
			mark[s[i]] = true
		}

		for j <= i {
			if _, ok := c2[s[j]]; ok && c1[s[j]] <= c2[s[j]] {
				break
			}
			c1[s[j]]--
			j++
		}

		if len(mark) == len(c2) && i-j < end-start {
			start, end = j, i
		}
	}

	// 不存在
	if end-start+1 > n {
		return ""
	}
	return string(bs[start : end+1])
}
