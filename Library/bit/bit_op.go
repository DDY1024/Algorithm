package main

import (
	"fmt"
	"math/bits"
)

// go 标准库 math/bits 提供的位运算操作
// 注意: bits.Len(0) = 0

// 运算符优先级: https://golang.org/ref/spec#Operators

// 32 位 or 64 位
const bitSize = 32 << (^uint(0) >> 32 & 1)

// & 和 | 操作在区间求和上具有单调性，其中 & 操作单调不增，| 操作单调不减
// ^ 操作 max_xor 通过 0/1 trie 实现

// 常用等式操作
// a|b = (a^b) + (a&b)
// a^b = (a|b) - (a&b)
// a+b = (a|b) + (a&b) = (a&b)*2 + (a^b) = (a|b)*2 - (a^b)
// (a&b)^(a&c) = a&(b^c) 结合律

func main() {
	fmt.Println(bits.Len(17))
	fmt.Println(bits.Len(0))
	fmt.Println(bits.Len(7))
	// bits.LeadingZeros()
}
