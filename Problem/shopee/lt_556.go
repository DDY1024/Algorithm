package main

import (
	"fmt"
	"math"
	"sort"
)

// 题目链接: https://leetcode-cn.com/problems/next-greater-element-iii/
// 题目大意
// 用原数字重新排列生成的下一个更大的数字（下一个排列算法经典应用）

func nextGreaterElement(n int) int {
	bits := make([]int, 0, 20)
	for n > 0 {
		bits = append(bits, n%10)
		n /= 10
	}

	if len(bits) < 2 {
		return -1
	}

	// 下一个排列算法生成更大的数字
	for i, j := 0, len(bits)-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}

	idx := len(bits) - 2
	for idx >= 0 && bits[idx] >= bits[idx+1] {
		idx--
	}
	if idx < 0 {
		return -1
	}

	// 倒序
	for j := len(bits) - 1; j > idx; j-- {
		if bits[j] > bits[idx] {
			bits[idx], bits[j] = bits[j], bits[idx]
			break
		}
	}
	sort.Ints(bits[idx+1:])

	ret := 0
	for i := 0; i < len(bits); i++ {
		ret = ret*10 + bits[i]
	}

	if ret > math.MaxInt32 {
		return -1
	}
	return ret
}

func main() {
	fmt.Println(nextGreaterElement(2147483476))
}
