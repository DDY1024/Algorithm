package main

// K 进制滚动 Hash
// https://juejin.cn/post/6844903638490415111
// https://zh.wikipedia.org/wiki/%E6%8B%89%E5%AE%BE-%E5%8D%A1%E6%99%AE%E7%AE%97%E6%B3%95

// 1. RK 算法在单模式串匹配情况下性能不如 KMP、BM 算法
// 2. 多模式串匹配场景下我们可以提前预处理出多模式串的 Hash 值，在 RK 算法进行滚动 Hash 计算时可以高效地进行匹配

const (
	mod  = 1e9 + 7 // 通常为素数
	base = 256     // 进制数，通常也为素数
)

func powMod(a, b, c int) int {
	ret := 1
	for b > 0 {
		if b&1 > 0 {
			ret = ret * a % c
		}
		a = a * a % c
		b >>= 1
	}
	return ret
}

func RabinKarp(data string, pattern string) []int {
	n, m := len(data), len(pattern)
	if m > n {
		return nil
	}

	patHash := 0
	indexList := make([]int, 0)
	for i := 0; i < m; i++ {
		patHash = (patHash*base + int(pattern[i])) % mod
	}

	var check = func(idx int) bool {
		for i, j := idx, 0; j < m; i, j = i+1, j+1 {
			if data[i] != pattern[j] {
				return false
			}
		}
		return true
	}

	dataHash := 0
	for i := 0; i < m; i++ {
		dataHash = (dataHash*base + int(data[i])) % mod
	}
	if dataHash == patHash && check(0) {
		indexList = append(indexList, 0)
	}

	cmod := powMod(base, m-1, mod)
	for i := m; i < n; i++ {
		dataHash = ((dataHash-int(data[i-m])*cmod%mod+mod)%mod*base + int(data[i])) % mod
		if dataHash == patHash && check(i-m+1) { // 在滚动哈希值相同的基础上，进一步判断是否相等
			indexList = append(indexList, i-m+1)
		}
	}
	return indexList
}
