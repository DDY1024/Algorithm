package main

import "fmt"

// https://juejin.cn/post/6844903638490415111
// 总结: 利用滚动 hash 进行判定，过滤掉大部分不匹配的场景

const (
	mod  = 100007
	base = 256
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
