package main

// Boyer-Moora 字符串匹配算法
// 参考：https://zh.wikipedia.org/wiki/%E5%8D%9A%E8%80%B6-%E7%A9%86%E5%B0%94%E5%AD%97%E7%AC%A6%E4%B8%B2%E6%90%9C%E7%B4%A2%E7%AE%97%E6%B3%95
// 核心思想通过两种方法来确定当字符串匹配发生失效时，下次匹配的位置，避免一些无效的匹配计算
// 1. 坏字符
// 2. 好后缀 --> 类似 KMP 算法中的 next 数组计算

// 一种单纯利用坏字符规则优化匹配过程
// Index returns the first index substr found in the s.
// function should return same result as `strings.Index` function
func Index(s string, substr string) int {
	d := CalculateSlideTable(substr)
	return IndexWithTable(&d, s, substr)
}

// IndexWithTable returns the first index substr found in the s.
// It needs the slide information of substr
func IndexWithTable(d *[256]int, s string, substr string) int {
	lsub := len(substr)
	ls := len(s)
	// 前置判断
	switch {
	case lsub == 0:
		return 0
	case lsub > ls:
		return -1
	case lsub == ls:
		if s == substr {
			return 0
		}
		return -1
	}

	i := 0
	for i+lsub-1 < ls {
		// 从后向前匹配
		j := lsub - 1
		for ; j >= 0 && s[i+j] == substr[j]; j-- {
		}

		// 模式串完全匹配
		if j < 0 {
			return i
		}

		// 计算坏字符串距离
		slid := j - d[s[i+j]]
		// 存在 <= 0，不会移动的情况，默认移动距离为 1
		if slid < 1 {
			slid = 1
		}

		// 移动
		i += slid
	}
	return -1
}

// CalculateSlideTable builds sliding amount per each unique byte in the substring
// 计算模式串的字符映射表 char --> pos
func CalculateSlideTable(substr string) [256]int {
	var d [256]int
	for i := 0; i < 256; i++ {
		d[i] = -1
	}
	for i := 0; i < len(substr); i++ {
		d[substr[i]] = i
	}
	return d
}

// func main() {
// 	debug.SetGCPercent(0)
// }
