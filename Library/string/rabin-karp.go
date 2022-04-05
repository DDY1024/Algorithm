package main

import "fmt"

// 单模式串匹配算法: K 进制滚动 Hash 算法判定
// 影响因素: Hash 算法冲突概率变大时，本身算法出现 Hash 值相同的概率会变大，最坏情况下会退化成 O(n*M)
// 算法大概原理:
// 1. 此算法使用散列函数以快速对每个位置能否匹配作大致的检测，此后只对通过了检测的位置进行匹配尝试
// 2. 相同字符串计算的哈希值肯定相同，哈希值相同字符串不一定相同(哈希冲突)
//
//
// https://juejin.cn/post/6844903638490415111
// https://zh.wikipedia.org/wiki/%E6%8B%89%E5%AE%BE-%E5%8D%A1%E6%99%AE%E7%AE%97%E6%B3%95

// 思考
// 1. RK 算法在单模式串匹配情况下性能不如 KMP、BM 算法
// 2. 多模式串匹配场景下我们可以提前预处理出多模式串的 Hash 集合，这样在 RK 算法计算滚动 Hash 时便可以高效计算
// 哪个模式串与当前字符串 Hash 值相同，然后进行进一步的相等性判定(多模式串匹配情况下更有优势)

const (
	mod  = 100007 // 模
	base = 256    // 进制
)

func quickPowMod(a, b, c int) int {
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

	constMod := quickPowMod(base, m-1, mod)
	for i := m; i < n; i++ {
		// 滚动 hash 计算方式
		dataHash = (((dataHash-int(data[i-m])*constMod%mod)+mod)%mod*base + int(data[i])) % mod
		if dataHash == patHash && check(i-m+1) {
			indexList = append(indexList, i-m+1)
		}
	}
	return indexList
}

func main() {
	fmt.Println(RabinKarp("aaa", "a"))
	fmt.Println(RabinKarp("abab", "ab"))
	fmt.Println(RabinKarp("abc", "d"))
	fmt.Println(RabinKarp("abc", "abc"))
	fmt.Println(RabinKarp("abc", "abcd"))
}
