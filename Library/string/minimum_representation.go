package main

import "sort"

// S[1...n]
// 字符串 S 的循环同构即 S[i...n] + S[1...i-1]
// 字符串 S 的最小表示即 S 的所有循环同构字符串中字典序最小的

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve1(s string) string {
	bs := []byte(s)
	i, j, k, n := 0, 1, 0, len(bs)
	// 如果 k >= n; 可以推导出整个字符串只由一类字符组成
	for i < n && j < n && k < n {
		if bs[(i+k)%n] == bs[(j+k)%n] {
			k++
		} else {
			if bs[(i+k)%n] > bs[(j+k)%n] {
				i++
			} else {
				j++
			}
			k = 0
			if i == j { // i 追上 j，j++
				j++
			}
		}
	}
	return string(bs[i:]) + string(bs[:i])
}

// 优化解法
//
//	假设当前匹配到 S[i,...,i+k-1] = S[j,...,j+k-1]
//		当 S[i+k] > S[j+k] 时，[i+1, i+k] 直接跳过，因为后续必然存在一个 j 其结果更优；i = i + k + 1
//		当 S[i+k] < S[j+k] 时，同理 j = j + k + 1
func solve2(s string) string {
	bs := []byte(s)
	i, j, k, n := 0, 1, 0, len(bs)
	for i < n && j < n && k < n {
		if bs[(i+k)%n] == bs[(j+k)%n] {
			k++
		} else {
			if bs[(i+k)%n] > bs[(j+k)%n] {
				i = i + k + 1
			} else {
				j = j + k + 1
			}
			k = 0
			if i == j {
				j++
			}
		}
	}

	// 最终结果存在三种情况：
	// 1. i >= n，则 j 为最小字符串起始下标
	// 2. j >= n, 则 i 为最小字符串起始下标
	// 3. i < n && j < n，则 i 或 j 任意即可，通常选择索引下标更小的
	idx := minInt(i, j)
	return string(bs[idx:]) + string(bs[:idx])
}

// 题目链接：https://leetcode.cn/problems/orderly-queue/
//
//  1. 当 k > 1 时，经过任意次操作后，必然会转化成【字母排列最小】的字典序
//  2. 当 k = 1 时，字符串【最小表示】
func orderlyQueue(s string, k int) string {
	bs := []byte(s)
	if k > 1 {
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		return string(bs)
	}

	var minInt = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	i, j, k, n := 0, 1, 0, len(s)
	for i < n && j < n && k < n {
		if s[(i+k)%n] == s[(j+k)%n] {
			k++
		} else {
			if s[(i+k)%n] > s[(j+k)%n] {
				i = i + k + 1
			} else {
				j = j + k + 1
			}
			k = 0
			if i == j {
				j++
			}
		}
	}
	idx := minInt(i, j)
	return s[idx:] + s[:idx]
}
