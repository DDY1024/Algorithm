package main

// 贪心解题思路: https://leetcode-cn.com/problems/longest-chunked-palindrome-decomposition/solution/tan-xin-fa-zheng-que-xing-de-xiang-xi-zheng-ming-b/
// https://leetcode-cn.com/problems/longest-chunked-palindrome-decomposition/submissions/
// https://zhuanlan.zhihu.com/p/83334559
// 贪心算法的证明类似于 kmp 算法中利用动态规划求解前缀函数的过程
// 此题我们最终采用贪心算法来求解，找到一个 s[i...i1] == s[j1...j] 成立，答案立即 + 2
func longestDecomposition(text string) int {
	n := len(text)
	match := make([][]int, n)
	for i := 0; i < n; i++ {
		match[i] = make([]int, n)
	}

	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if text[i] == text[j] {
				match[i][j] = 1
				if i+1 < j && j+1 < n {
					match[i][j] += match[i+1][j+1]
				}
			}
		}
	}

	start, end, ans := 0, n-1, 0
	for start <= end {
		l, ok := end-start+1, false
		for st := 1; st <= l/2; st++ {
			if match[start][end-st+1] >= st {
				ok = true
				start += st
				end -= st
				ans += 2
				break
			}
		}
		if !ok {
			ans++
			break
		}
	}
	return ans
}
