package main

// 参考资料
// https://oi-wiki.org/string/minimal-string/

// 原串 S，循环同构 T
// T = S[i:n-1] + S[0:i-1]
//
// 最小表示即为 S 的所有循环同构 T 中字典序最小的字符串

// 1. 暴力解法
/*
	bs := []byte(s)
	i, j, k, n := 0, 1, 0, len(bs)
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
			if i == j { // 存在 i 追上 j 的情况
				j++
			}
		}
	}
	return string(bs[i:]) + string(bs[:i])
*/

// 2. 优化解法
// 利用之前的匹配结果，避免一些无用比较
// 例如 S[i..i+k-1] == S[j...j+k-1]
// 当 S[i+k] > S[j+k] 时，则在 [i,i+k] 区间的 i 不需要再遍历，因为总存在一个 l 属于 [j,j+k] 其结果更优，则 i 直接跳到 i + k + 1 即可
// 当 S[i+k] < S[j+k] 时，则 j = j + k + 1
/*
	// 注意: 优化后的算法，并不保证 i <= j 一直成立，因为存在跳跃
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

	// 最终结果只会存在下面三种情况
	// 1. i >= n，则 j 为最小字符串表示
	// 2. j >= n, 则 i 为最小字符串表示
	// 3. i < n && j < n，则选择 i 或 j 任意一个即可，我们通常下标更小的
	idx := minInt(i, j)

	return string(bs[idx:]) + string(bs[:idx])
*/

// 相关题目
// 1. https://leetcode.cn/problems/orderly-queue/
// a. 对于 k > 1 时，我们在任意次操作后，总会把字符串转化成原字符串字符重新排列后的最小字符串
// b. 对于 k = 1，便是字符串最小表示问题，由于题目数据范围为 1000，我们可以直接暴力求解 O(n^2)
