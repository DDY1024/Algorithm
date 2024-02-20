package main

// https://blog.csdn.net/mylinchi/article/details/79508112
// 8 个常用字符串哈希函数 https://developer.aliyun.com/article/252773
func BKDRHash(s string) uint {
	seed := uint(131) // 31, 131, 1313, 13131, 131313 etc..
	hash := uint(0)
	for i := 0; i < len(s); i++ {
		hash = hash*seed + uint(s[i])
	}
	return hash & 0x7FFFFFFF
}
