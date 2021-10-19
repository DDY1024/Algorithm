package main

import "fmt"

// 参考资料:
// 核心流程
// 1. 动态规划求解失效函数: 最长(前缀 = 后缀) --> 对应最右端点下标 i
// 2. 根据求解的失效函数确定不匹配时移动的距离，遍历主串完成匹配

func calcNext(pattern string) []int {
	n := len(pattern)
	next := make([]int, n)
	next[0] = -1
	k := -1
	for i := 1; i < n; i++ {
		// 思考如何从 next[i]、next[next[i]]、... 推导出 next[i+1]
		for k != -1 && pattern[k+1] != pattern[i] {
			k = next[k]
		}
		// 新位置处的匹配情况
		if pattern[k+1] == pattern[i] {
			k++
		}
		// 计算结果
		next[i] = k
	}
	return next
}

func kmp(text, pattern string) int {
	next := calcNext(pattern)
	// fmt.Println("Next Result:", next)
	n, m, idx := len(text), len(pattern), 0
	for i := 0; i < n; i++ {
		// 失效函数的价值体现出来了
		for idx > 0 && text[i] != pattern[idx] {
			idx = next[idx-1] + 1 // 失配计算
		}
		// 重新
		if text[i] == pattern[idx] {
			idx++
		}
		if idx == m { // 找到一处匹配处
			return i - m + 1
		}
	}
	return -1
}

// test
func main() {
	fmt.Println(kmp("aaa", "aa"))
	fmt.Println(kmp("abab", "ab"))
	fmt.Println(kmp("abc", "d"))
	fmt.Println(kmp("abc", "abc"))
	fmt.Println(kmp("abc", "abcd"))
	fmt.Println(kmp("abcab", "b"))
}
