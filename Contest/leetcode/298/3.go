package main

import (
	"math/bits"
	"strconv"
	"strings"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 提供一种动态规划解题思路
// dp[i] 表示长度为 i 的子序列的最小值，容易知道这个子序列肯定是以某个 s[i] 为结尾的
// 这样在从前往后枚举子序列的过程中，对于每个 i，我们判断由其为结尾构成长度为 j 的子序列的最小值
// 最终枚举到 s[n-1] 时，我们判断 dp[1] ... dp[n] 最大长度的子序列 <= k 即为答案
func longestSubsequence(s string, k int) int {
	n := len(s)
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = k + 0xff // 随便找个 > k 的值即可
	}

	dp[0] = 0
	for i := 0; i < n; i++ { // 此处要逆序，防止 j 在前面状态被选择后，从而对后面的状态产生影响（满足无后效性）
		for j := n; j > 0; j-- {
			num := (dp[j-1] << 1) + int(s[i]-'0')
			if num <= k {
				dp[j] = minInt(dp[j], num)
			}
		}
	}

	for j := n; j > 0; j-- {
		if dp[j] <= k {
			return j
		}
	}

	// 由题目数据范围可知，结果肯定 >= 1，因此此处实际上是不可达的
	return -1
}

// 这道题目的状态转移方程与 LIS 的 O(nlogn) 解法有点相似

// 实际上这道题目是可以用贪心思路来求解的：https://leetcode.cn/problems/longest-binary-subsequence-less-than-or-equal-to-k/solution/fen-lei-tao-lun-tan-xin-by-endlesscheng-vnlx/
// 1. 如果 k 的二进制字符串的表示长度 > len(s)，那 s 直接满足返回即可
// 2. 假设 k 的二进制字符串长度为 m，我们需要做的事情是在 m 中找一个最靠后的后缀，长度 <= m 且数值 <= k
// 这样我们便可以利用前面更多的前导 0 来进行填充使得最终结果更大。为什么这样是正确的？我们分类讨论下
// a. 假设 s 的最后 m 位组成的二进制数 <= k，则直接使用；很显然在 [0,m-1) 区间内寻找一个后缀长度 <= k 的结果不会比现在结果更好
// b. 假设 s 的最后 m 位组成的二进制数 > k，则我们保留 m-1 位，剩余区间堆前导 0; 因为在 [0,m-1) 寻找一个 m 长度的后缀 <= k，最优情况下会多占一位
// 且多占的这一位肯定是 0（因为为 1 肯定是 > k 的）
// 综上所述，我们便有了最终的贪心方案
// 1. 先选最后 m 位，如果 <= k，则为 m + 前导 0 个数；如果 > k，则为 m-1 + 前导 0 个数

func longestSubsequence2(s string, k int) int {
	n, m := len(s), bits.Len(uint(k))
	if m > n {
		return n
	}

	ret := m
	x, _ := strconv.ParseInt(s[n-m:], 2, 64) // 2 进制解析
	if int(x) > k {
		ret--
	}

	return ret + strings.Count(s[:n-m], "0")
}
