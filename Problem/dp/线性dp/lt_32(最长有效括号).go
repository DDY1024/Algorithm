package main

// 题目链接：https://leetcode.cn/problems/longest-valid-parentheses/
//
// 解题思路
// 		dp[i]: 表示以 i 结尾的最长有效括号长度
//
// https://leetcode.cn/problems/longest-valid-parentheses/solution/zui-chang-you-xiao-gua-hao-by-leetcode-solution/

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	ret := 0
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			continue
		}

		// s[i-1] = '(' && s[i] = ')'
		if s[i-1] == '(' {
			dp[i] = 2
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			}
			ret = maxInt(ret, dp[i])
			continue
		}

		// s[i] = ')' && s[i-1] = ')'
		if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
			ret = maxInt(ret, dp[i])
		}
	}

	return ret
}
