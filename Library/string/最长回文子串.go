package main

import "fmt"

// https://segmentfault.com/a/1190000003914228
// 将求解传统算法中求解最长回文子串的奇偶性问题，通过填充特殊字符转换成求解回文半径的问题
//
// 维护两个变量
// 1. 当前回文子串能够触及的最右索引位置
// 2. (1) 情况下，该回文子串的对称中心
//
// 存在如下递推关系
// 1. i < maxRight: RL[i] = min(RL[2*pos-i], maxRight - pos)
// 2. i >= maxRight: RL[i] = 1
// 3. 上述两个过程推导完成以后，均应该从当前位置扩展回文子串能够扩展到的最右位置

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Manacher(s string) int {

	var (
		R        []int // 回文半径
		maxPos   int
		maxRight int
		maxLen   int
		sb       []byte
		n        int
	)

	n = len(s)
	sb = make([]byte, 2*n+1)
	for i := 0; i < n; i++ {
		sb[2*i] = '*'
		sb[2*i+1] = s[i]
	}
	sb[2*n] = '*'

	n = 2*n + 1
	R = make([]int, n)
	maxPos, maxRight, maxLen = 0, 0, 1
	for i := 0; i < n; i++ {
		if i < maxRight {
			R[i] = minInt(R[2*maxPos-i], maxRight-i+1) // 递推
		} else {
			R[i] = 1
		}
		for i+R[i] < n && i-R[i] >= 0 && sb[i-R[i]] == sb[i+R[i]] {
			R[i]++
		}
		if maxRight < i+R[i]-1 {
			maxRight = i + R[i] - 1
			maxPos = i // 回文半径最右端的中心点位置
		}
		maxLen = maxInt(maxLen, R[i])
	}
	return maxLen // 最长回文子串的长度为 maxLen - 1
}

// maxPos --> maxRight
// maxPos --> i ---> maxRight
// maxPos-(maxRight-maxPos) --> maxPos-(i-maxPos) --> maxPos --> i --> maxRight
// maxPos - (i - maxPos)

func main() {
	s1 := "a"
	fmt.Println(Manacher(s1) - 1)
	s2 := "abc"
	fmt.Println(Manacher(s2) - 1)
	s3 := "aaa"
	fmt.Println(Manacher(s3) - 1)
	s4 := "abba"
	fmt.Println(Manacher(s4) - 1)
	s5 := "abbabbbbb"
	fmt.Println(Manacher(s5) - 1)
	s6 := "abbbbba"
	fmt.Println(Manacher(s6) - 1)
	s7 := "abxbbx"
	fmt.Println(Manacher(s7) - 1)

	// mp := map[int]int{}
	// mp[1] = 1
	// delete(mp, 1)
	// mp[2] = 2
	// delete(mp, 2)
	// for k, v := range mp {
	// 	fmt.Println(k, v)
	// }
}
