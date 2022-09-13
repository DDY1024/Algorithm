package main

// https://leetcode-cn.com/problems/maximize-palindrome-length-from-subsequences/
// 首先我们需要知道一个结论，如何求解一个字符串的最长回文子序列，通过 O(n^2) 动态规划预处理求解
// dp(i, j) = dp(i+1, j-1) + 1 (s[i] == s[j])
// dp(i, j) = max(dp(i+1, j), dp(i, j-1)) (s[i] != s[j])
// 这样我们便预处理出 (i, j) 区间最长回文子序列了
// 由于最终要求解回文子序列需要两个原始字符串共同贡献，因此我们只需要枚举最终的回文子序列的左右端点在两个串中的哪个位置，然后通过
// 预处理结果直接 O(1) 求解最长回文子序列的长度即可。

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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
				dp[i][j] = maxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	// 也可以在 dp 过程中直接求解，不需要单独拎出来
	ans := 0
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if word1[i] == word2[j] {
				ans = maxInt(ans, dp[i][n1+j])
			}
		}
	}

	return ans
}
