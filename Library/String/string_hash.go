package main

// 字符串 Hash 介绍: https://oi-wiki.org/string/hash/
// Base 进制: 素数
// Mod 模数：素数

const (
	base = 131313  // 先随便整一个，但是最好取素数
	mod  = 1e7 + 7 // 取一个较大的素数
)

var (
	prefixHash []int // base 进制下的前缀 hash 值
	baseMod    []int //
)

// Rabin-Karp 滚动 Hash 函数计算
func prepareCalc(s string) {
	n := len(s)
	prefixHash = make([]int, n+1)
	baseMod = make([]int, n+1)
	prefixHash[0] = 0
	baseMod[0] = 1
	for i := 1; i <= n; i++ {
		prefixHash[i] = (prefixHash[i-1]*base + int(s[i-1])) % mod // s[i-1]
		baseMod[i] = baseMod[i-1] * base % mod
	}
}

func calcSegment(l, r int) int {
	return ((prefixHash[r]-prefixHash[l-1]*baseMod[r-l+1])%mod + mod) % mod
}

// 2. 另外一篇介绍各种字符串 Hash 函数: https://blog.csdn.net/mylinchi/article/details/79508112
//
// https://blog.csdn.net/MyLinChi/article/details/79509455
func BKDRHash(s string) uint {
	seed := uint(131) // 31 131 1313 13131 131313 etc..
	hash := uint(0)
	for i := 0; i < len(s); i++ {
		hash = hash*seed + uint(s[i]) // 利用整数溢出的性质
	}
	return hash & 0x7FFFFFFF
}

// APHash 算法
/*
unsigned int APHash(char *str)
{
    unsigned int hash = 0;
    int i;

    for (i=0; *str; i++)
    {
        if ((i & 1) == 0)
        {
            hash ^= ((hash << 7) ^ (*str++) ^ (hash >> 3));
        }
        else
        {
            hash ^= (~((hash << 11) ^ (*str++) ^ (hash >> 5)));
        }
    }

    return (hash & 0x7FFFFFFF);
}
*/

// DJBHash 算法
/*
unsigned int DJBHash(char *str)
{
    unsigned int hash = 5381;

    while (*str)
    {
        hash += (hash << 5) + (*str++);
    }

    return (hash & 0x7FFFFFFF);
}
*/

// RSHash 算法
/*
unsigned int RSHash(char *str)
{
    unsigned int b = 378551;
    unsigned int a = 63689;
    unsigned int hash = 0;

    while (*str)
    {
        hash = hash * a + (*str++);
        a *= b;
    }

    return (hash & 0x7FFFFFFF);
}
*/

// https://leetcode-cn.com/problems/longest-duplicate-substring/  lt_1044 二分 + 字符串滚动 hash 计算
