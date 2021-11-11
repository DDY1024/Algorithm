package main

// 题目链接：https://leetcode-cn.com/problems/check-if-an-original-string-exists-given-two-encoded-strings/
// 解题思路
// 1. 算做是字符串编辑距离问题的变种，编辑距离问题中我们有 '?' 和 '*' 来匹配任意字符，此问题中我们变为数字字符串匹配指定任意字母的个数
// 2. 状态表示: dp(i, j, d) 表示 s1 的前 i 个字符串与 s2 的前 j 个字符串匹配最终多出 d 个任意字符是否可行
// 3. 状态转移（当前表示字符串长度较短的一方优先进行扩展）
// d = 0
// d <= 0
// d >= 0
// 上述三种情况分别讨论转移情况
// 搜索递推前进方向 --> 短串优先进行扩展 --> 最多三位连续数字 --> d 取值范围 (-1000, 1000)

func possiblyEquals(s1 string, s2 string) bool {
	n, m, offset := len(s1), len(s2), 1000
	vis := make([][][]bool, n+1)
	for i := 0; i <= n; i++ {
		vis[i] = make([][]bool, m+1)
		for j := 0; j <= m; j++ {
			vis[i][j] = make([]bool, 2010)
		}
	}

	var isDigit = func(ch byte) bool {
		return ch >= '0' && ch <= '9'
	}

	var dfs func(i, j, d int) bool
	dfs = func(i, j, d int) bool {
		// 边界条件
		if i >= n && j >= m {
			return d == 0
		}

		// s2 串表示的个数肯定比 s1 串多
		if i >= n && d <= 0 {
			return false
		}

		// s1 传表示的个数肯定比 s2 串多
		if j >= m && d >= 0 {
			return false
		}

		// 此处 vis 保存的是被计算过的 false 状态，因为一旦出现 true 状态，我们便及时 return 了
		if vis[i][j][d+offset] {
			return false
		}

		// d = 0
		// 1. s[i] == s[j] --> d(i+1, j+1, 0)
		// 2. s[i] != s[j] --> false（s[i], s[j] 均为字母）
		// 该情况下 i < n && j < m
		// 无论是字符还是数字一旦相等便可以进一步递推
		if d == 0 && s1[i] == s2[j] && dfs(i+1, j+1, 0) {
			return true
		}

		// d <= 0 --> 拓展 s1 长度短的
		if d <= 0 {
			// 数字拓展
			if isDigit(s1[i]) {
				for p, v := i, 0; p < n && isDigit(s1[p]); p++ {
					v = v*10 + int(s1[p]-'0')
					if dfs(p+1, j, d+v) {
						return true
					}
				}
			} else {
				// 字符拓展，由于 d 的存在是由于数字拓展出来的，因此只有 d != 0 时当前字符才能被当做字符进行拓展
				if d < 0 && dfs(i+1, j, d+1) { // 注意此处: d 不能为 0, 我们的 d 不为 0 的情况只存在于由数字构成的任意字符串
					return true
				}
			}
		}

		// d >= 0 --> 拓展 s2 长度短的
		if d >= 0 {
			// 数字拓展
			if isDigit(s2[j]) {
				for p, v := j, 0; p < m && isDigit(s2[p]); p++ {
					v = v*10 + int(s2[p]-'0')
					if dfs(i, p+1, d-v) {
						return true
					}
				}
			} else {
				// 字符拓展，同理 s1
				if d > 0 && dfs(i, j+1, d-1) { // 注意此处: d 不能为 0, 我们的 d 不为 0 的情况只存在于由数字构成的任意字符串
					return true
				}
			}
		}

		// 标记当前已经被计算过的 false 状态
		vis[i][j][d+offset] = true
		return false
	}
	return dfs(0, 0, 0)
}
