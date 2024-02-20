package main

// 最长回文子序列（区间 DP）
// 1. 初始条件
// 		dp[i][i] = 1
// 2. 状态转移
//      如果 s[i] == s[j]，则 dp[i][j] = dp[i+1][j-1] + 2  (贪心计算)
//      如果 s[i] != s[j]，则 dp[i][j] = max{ dp[i+1][j], dp[i][j-1] }

func longestPalindrome(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+n2)
	for i := 0; i < n1+n2; i++ {
		dp[i] = make([]int, n1+n2)
	}

	word := make([]byte, 0, n1+n2)
	word = append(word, []byte(word1)...)
	word = append(word, []byte(word2)...)

	n := n1 + n2
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			if word[i] == word[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	ans := 0
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if word1[i] == word2[j] {
				ans = max(ans, dp[i][n1+j])
			}
		}
	}
	return ans
}
