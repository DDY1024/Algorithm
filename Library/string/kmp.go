package main

// 核心思想
// 1. 利用动态规划求解模式串 pattern 的失效函数：每个【前缀字符串】的 【最长前缀 = 最长后缀】（不包含 pattern 本身）
// 2. 主串与模式串进行匹配时，如果发生某个位置的不匹配，主串直接前移指定距离（失效函数），避免无效的累加操作

func calcNext(pattern string) []int {
	n := len(pattern)
	next := make([]int, n)

	next[0] = -1
	k := -1
	for i := 1; i < n; i++ {
		for k != -1 && pattern[k+1] != pattern[i] {
			k = next[k]
		}
		if pattern[k+1] == pattern[i] {
			k++
		}
		next[i] = k
	}
	return next
}

func kmp(text, pattern string) int {
	next := calcNext(pattern)
	n, m, idx := len(text), len(pattern), 0
	for i := 0; i < n; i++ {
		for idx > 0 && text[i] != pattern[idx] {
			idx = next[idx-1] + 1 // 失配计算
		}

		if text[i] == pattern[idx] {
			idx++
		}

		if idx == m { // 存在一处匹配
			// 如果需要找到所有匹配处：则 idx = next[idx-1] + 1
			return i - m + 1
		}
	}
	return -1
}
