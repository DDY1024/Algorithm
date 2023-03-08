package main

import (
	"sort"
)

// https://leetcode.cn/problems/remove-invalid-parentheses/
//
// 搜索 + 剪枝
//
// 		左括号数量 >= 右括号数量

func removeInvalidParentheses(s string) []string {
	n := len(s)
	ans := make([]string, 0)
	bs := make([]byte, 0, n)
	rrc := n + 1

	var dfs func(pos, lbc, rbc, rc int)
	dfs = func(pos, lbc, rbc, rc int) {
		if pos >= n {
			if lbc == rbc {
				if rc < rrc {
					rrc = rc
					ans = make([]string, 0)
					ans = append(ans, string(bs))
				} else if rc == rrc {
					ans = append(ans, string(bs))
				}
			}
			return
		}

		if rc > rrc { // 剪枝
			return
		}

		if rbc > lbc { // 非法
			return
		}

		tlen := len(bs)

		if s[pos] >= 'a' && s[pos] <= 'z' {
			bs = append(bs, s[pos])
			dfs(pos+1, lbc, rbc, rc)
			return
		}

		if s[pos] == '(' {
			bs = append(bs, s[pos])
			dfs(pos+1, lbc+1, rbc, rc)

			bs = bs[:tlen]
			dfs(pos+1, lbc, rbc, rc+1)
			return
		}

		if lbc == rbc {
			dfs(pos+1, lbc, rbc, rc+1)
			return
		}

		bs = append(bs, s[pos])
		dfs(pos+1, lbc, rbc+1, rc)

		bs = bs[:tlen]
		dfs(pos+1, lbc, rbc, rc+1)
	}

	dfs(0, 0, 0, 0)

	// 去重
	sort.Strings(ans)
	i, j := 0, 0
	for j < len(ans) {
		if ans[j] == ans[i] {
			j++
			continue
		}
		i++
		ans[i] = ans[j]
		j++
	}
	ans = ans[:i+1]
	return ans
}

// 也可以采用 bfs 进行搜索
