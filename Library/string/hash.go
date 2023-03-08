package main

// 字符串 Hash 介绍: https://oi-wiki.org/string/hash/

const (
	base = 131313  // 素数
	mod  = 1e9 + 7 // 素数
)

var (
	prefixHash []int
	baseMod    []int
)

func prepare(s string) {
	n := len(s)
	prefixHash = make([]int, n+1)
	baseMod = make([]int, n+1)

	prefixHash[0] = 0
	baseMod[0] = 1
	for i := 1; i <= n; i++ {
		prefixHash[i] = (prefixHash[i-1]*base + int(s[i-1])) % mod
		baseMod[i] = baseMod[i-1] * base % mod
	}
}

func calcSegment(l, r int) int {
	return (prefixHash[r] - prefixHash[l-1]*baseMod[r-l+1]%mod + mod) % mod
}

// 参考：https://blog.csdn.net/mylinchi/article/details/79508112
func BKDRHash(s string) uint {
	seed := uint(131) // 31, 131, 1313, 13131, 131313 etc..
	hash := uint(0)
	for i := 0; i < len(s); i++ {
		hash = hash*seed + uint(s[i]) // 利用整数溢出的性质
	}
	return hash & 0x7FFFFFFF
}
