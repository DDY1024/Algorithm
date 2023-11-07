package main

import "fmt"

// 十进制 <--> 二进制
//	 a. 十进制整数转二进制整数，不断除 2 取余，逆序输出
//   b. 十进制小数转二进制小数，不断整 2 取整，正序输出

// 位运算常见公式
// 		a|b = (a^b) + (a&b)
//		a^b = (a|b) - (a&b)
// 		a+b = (a|b) + (a&b) = (a&b)*2 + (a^b) = (a|b)*2 - (a^b)
//		(a&b)^(a&c) = a&(b^c)

// 判断 32 位 或 64 位
const bitSize = 32 << (^uint(0) >> 32 & 1)

func bitAbs(n int) int {
	return (n ^ (n >> 31)) - (n >> 31)
}

func bitMax(a, b int) int { return (b & ((a - b) >> 31)) | (a & (^(a - b) >> 31)) }
func bitMin(a, b int) int { return (a & ((a - b) >> 31)) | (b & (^(a - b) >> 31)) }

func sign(x, y int) bool {
	return (x ^ y) >= 0
}

func swap(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

func getBit(a, b int) int {
	return (a >> b) & 1
}

// 1. a &^ (1<<b)      清除
// 		11 &^ 10 = 01
//
// 2. a | (1<<b)       设置
// 		10 | 01 = 11
//
// 3. a ^ (1<<b)       反转
//      11 ^ 01 = 10
//

func bitCountOne(x int) int {
	cnt := 0
	for x > 0 {
		cnt += x & 1
		x >>= 1
	}
	return cnt
}

func bitCountTwo(x int) int {
	cnt := 0
	for x > 0 {
		x -= x & (-x)
		cnt++
	}
	return cnt
}

func bitCountThree(x int) int {
	cnt := 0
	for x > 0 {
		x &= (x - 1)
		cnt++
	}
	return cnt
}

// 快速统计 32 位整数中 1 的个数，复杂度 O(loglogn)
func countBitTwo(n uint32) uint32 {
	// 01,01,01,...,01
	// 0011,...,0011
	// 00001111,...,00001111
	// 0000000011111111,...,0000000011111111
	// 00000000000000001111111111111111
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	n = (n & 0x0f0f0f0f) + ((n >> 4) & 0x0f0f0f0f)
	n = (n & 0x00ff00ff) + ((n >> 8) & 0x00ff00ff)
	n = (n & 0x0000ffff) + ((n >> 16) & 0x0000ffff)
	return n
}

// 非空子集枚举、全部子集枚举
func findSubset(x int) []int {
	arr := make([]int, 0)
	for i := x; i > 0; i = (i - 1) & x {
		arr = append(arr, i)
	}

	// 0 代表空集
	arr = append(arr, 0)
	return arr
}

// 汉明权重排列
// 00, 01, 10, 11
// 000, 001, 010, 100, 011, 101, 110, 111
func solve(n int) {
	for i := 0; (1<<i)-1 <= n; i++ {
		x, t := (1<<i)-1, 0
		for x <= n {
			fmt.Println(x)
			t = x + (x & -x)
			if x > 0 {
				x = t | (((t & -t) / (x & -x)) >> 1) - 1
			} else {
				x = n + 1
			}
		}
	}
}

func main() {

	// fmt.Println(bitCountOne(8))
	// fmt.Println(bitCountTwo(8))
	// fmt.Println(bitCountThree(8))

	// fmt.Println(bitAbs(-4))
	// fmt.Println(bitAbs(-1))
	// fmt.Println(bitAbs(9))

	// fmt.Println(bitMax(-1, 5))
	// fmt.Println(bitMin(-1, 5))
	// fmt.Println(bitMax(1, 5))
	// fmt.Println(bitMin(1, 5))

	// fmt.Println(sign(1, 3))
	// fmt.Println(sign(-1, 3))

	// a, b := 1, 2
	// swap(&a, &b)
	// fmt.Println(a, b)
	// fmt.Println(getBit(2, 0))
	// fmt.Println(getBit(2, 1))
	// fmt.Println(3 &^ (1 << 1))

	// fmt.Println(findSubset(3)) // 3, 2, 1, 0
	// fmt.Println(findSubset(6)) // 6, 4, 2, 0
	solve(7)
}
