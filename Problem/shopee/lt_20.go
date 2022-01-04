package main

// 题目链接: https://leetcode-cn.com/problems/valid-parentheses/
// 解题思路
// 1. 括号配对: 栈的应用

func isValid(s string) bool {
	n := len(s)
	stk := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		switch s[i] {
		case '(', '[', '{':
			stk = append(stk, s[i])
		case ')':
			if len(stk) == 0 || stk[len(stk)-1] != '(' {
				return false
			}
			stk = stk[:len(stk)-1]
		case ']':
			if len(stk) == 0 || stk[len(stk)-1] != '[' {
				return false
			}
			stk = stk[:len(stk)-1]
		case '}':
			if len(stk) == 0 || stk[len(stk)-1] != '{' {
				return false
			}
			stk = stk[:len(stk)-1]
		}
	}
	return len(stk) == 0
}
