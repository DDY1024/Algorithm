package main

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	match := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		match[i] = make([]bool, m+1)
	}
	match[0][0] = true

	// 注意边界条件的处理
	// 例如: "" 匹配 "a*"，"" 匹配 ".*", "" 匹配 "a*b*.*" 的情况
	// match[i][j]: default false
	for j := 1; j <= m; j++ {
		if p[j-1] == '*' {
			match[0][j] = match[0][j-2]
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 分类讨论而已
			if p[j-1] >= 'a' && p[j-1] <= 'z' {
				match[i][j] = match[i-1][j-1] && (s[i-1] == p[j-1])
			} else if p[j-1] == '.' {
				match[i][j] = match[i-1][j-1]
			} else if p[j-1] == '*' {
				match[i][j] = match[i][j] || match[i][j-2] // * 匹配零次
				if p[j-2] == '.' {
					for k := 0; k < i; k++ {
						match[i][j] = match[i][j] || match[k][j-2]
					}
				} else {
					for k := i; k > 0 && s[k-1] == p[j-2]; k-- {
						match[i][j] = match[i][j] || match[k-1][j-2]
					}
				}
			}
		}
	}

	return match[n][m]
}
