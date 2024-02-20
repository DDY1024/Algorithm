package main

// https://leetcode.cn/problems/shortest-common-supersequence/

// 求解两个字符串的 LCS，然后由 LCS 逆向构造最短字符串

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

func shortestCommonSupersequence(str1 string, str2 string) string {
	n, m := len(str1), len(str2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	ans := make([]byte, 0, n+m)
	i, j := n, m
	for i > 0 || j > 0 {
		if i == 0 {
			for k := j - 1; k >= 0; k-- {
				ans = append(ans, str2[k])
			}
			break
		}

		if j == 0 {
			for k := i - 1; k >= 0; k-- {
				ans = append(ans, str1[k])
			}
			break
		}

		// 注意此处的判断条件不能是 dp[i][j] = dp[i-1][j-1] + 1
		// 因为 dp[i][j] --> dp[i-1][j], dp[i][j-1] --> dp[i-1][j-1] + 1，这种情况是完全存在的
		if str1[i-1] == str2[j-1] { // 相等，直接按照最优进行转移
			ans = append(ans, str1[i-1])
			i--
			j--
		} else if dp[i][j] == dp[i-1][j] {
			ans = append(ans, str1[i-1])
			i--
		} else {
			ans = append(ans, str2[j-1])
			j--
		}
	}

	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return string(ans)
}
