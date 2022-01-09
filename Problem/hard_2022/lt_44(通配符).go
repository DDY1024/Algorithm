package main

// 题目链接: https://leetcode-cn.com/problems/wildcard-matching/
// 题目大意
// 实现 '?' 和 '*' 的通配符匹配，典型的动态规划问题

// 通配符动态规划匹配求解
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	ok := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		ok[i] = make([]int, m+1)
	}

	ok[0][0] = 1
	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i == 0 && j == 0 {
				continue
			}

			switch p[j-1] {
			case '?':
				if i == 0 {
					ok[i][j] = 0
				} else {
					ok[i][j] = ok[i-1][j-1]
				}
			case '*':
				ok[i][j] = ok[i][j-1]
				if i-1 >= 0 {
					ok[i][j] |= ok[i-1][j]
				}
			default:
				ok[i][j] = 0
				if i-1 >= 0 && s[i-1] == p[j-1] {
					ok[i][j] = ok[i-1][j-1]
				}
			}
		}
	}
	return ok[n][m] == 1
}
