package main

// 核心思想
//  	1. 动态规划求解【模式串】所有【前缀子串】的失效函数：即求解字符串的【最长前缀】=【最长后缀】
//      2. 【主串】与【模式串】进行匹配时，如果发生【不匹配】，根据【失效函数】进行转移

func calc(pattern string) []int {
	var (
		n    = len(pattern)
		next = make([]int, n)
	)

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
	next := calc(pattern)
	n, m, idx := len(text), len(pattern), 0
	for i := 0; i < n; i++ {
		for idx > 0 && text[i] != pattern[idx] {
			idx = next[idx-1] + 1 // next[0] = -1
		}

		if text[i] == pattern[idx] {
			idx++
		}

		if idx == m { // 存在一处完全匹配
			// 如果需要输出所有匹配位置：idx = next[idx-1] + 1
			return i - m + 1
		}
	}
	return -1
}
