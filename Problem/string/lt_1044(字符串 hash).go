package main

import (
	"math"
	"math/rand"
	"time"
)

// 题目链接: https://leetcode-cn.com/problems/longest-duplicate-substring/
// 解题思路
// 1. 答案长度存在单调性，因此可以二分枚举答案长度
// 2. O(n) 复杂度判断一个字符串中是否存在长度为 l 的重复子串 --> 字符串 Hash --> 滚动计算 --> 避免冲突采用多 hash 函数

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

func randRange(l, r int) int {
	return l + rand.Intn(r-l+1)
}

type Pair struct {
	x, y int
}

func longestDupSubstring(s string) string {
	rand.Seed(time.Now().UnixNano())
	b1, b2 := randRange(26, 100), randRange(26, 100)
	m1, m2 := randRange(1e9+7, math.MaxInt32), randRange(1e9+7, math.MaxInt32)
	l, r, n := 1, len(s)-1, len(s)
	var check = func(x int) int {
		h1, h2 := 0, 0
		t1, t2 := powMod(b1, x, m1), powMod(b2, x, m2)
		mark := make(map[Pair]bool, n)
		for i := 0; i < x; i++ {
			h1 = (h1*b1 + int(s[i])) % m1
			h2 = (h2*b2 + int(s[i])) % m2
		}
		mark[Pair{h1, h2}] = true
		for i := x; i < n; i++ {
			h1 = ((h1*b1-int(s[i-x])*t1+int(s[i]))%m1 + m1) % m1
			h2 = ((h2*b2-int(s[i-x])*t2+int(s[i]))%m2 + m2) % m2
			if mark[Pair{h1, h2}] {
				return i - x + 1
			}
			mark[Pair{h1, h2}] = true
		}
		return -1
	}

	ridx, rlen := -1, -1
	for l <= r {
		mid := l + (r-l)/2
		idx := check(mid)
		if idx != -1 {
			ridx = idx
			rlen = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if rlen <= 0 {
		return ""
	}
	return s[ridx : ridx+rlen]
}
