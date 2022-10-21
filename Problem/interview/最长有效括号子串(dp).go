package main

// 解题报告: https://leetcode.cn/problems/longest-valid-parentheses/solution/zui-chang-you-xiao-gua-hao-by-leetcode-solution/

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

// ()() --> 4
// ())()() --> 4
// ()(() --> 2 not 4

func longestValidParentheses(s string) int {
	ret, n := 0, len(s)
	stk := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stk = append(stk, i)
		} else {
			if len(stk) == 0 {
				continue
			}

			cnt := 0
			for len(stk) > 0 && stk[len(stk)-1] < 0 { // 实际上只可能一个
				cnt += -stk[len(stk)-1]
				stk = stk[:len(stk)-1]
			}
			if len(stk) == 0 {
				continue
			}

			// 匹配 ()
			cnt += 2
			stk = stk[:len(stk)-1]
			for len(stk) > 0 && stk[len(stk)-1] < 0 {
				cnt += -stk[len(stk)-1]
				stk = stk[:len(stk)-1]
			}
			ret = maxInt(ret, cnt)
			stk = append(stk, -cnt)
		}
	}
	return ret
}

// 解题思路：https://leetcode.cn/problems/longest-valid-parentheses/solution/zui-chang-you-xiao-gua-hao-by-leetcode-solution/
// dp[i]: 表示以 i 结尾的最长有效括号的长度

func longestValidParenthesesDP(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)

	ret := 0
	// 思考：括号匹配情况下，如何进行有效的状态转移？
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
