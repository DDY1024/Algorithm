package main

// https://segmentfault.com/a/1190000003914228
//
// 维护两个变量
// 1. 当前回文子串能够触及的最右索引位置
// 2. (1) 情况下，该回文子串的对称中心
//
// 存在如下递推关系
// 1. i < maxRight: RL[i] = min{RL[2*maxPos-i], maxRight - i + 1}
// 2. i >= maxRight: RL[i] = 1
// 3. 在 1 和 2 判定的基础上，向两边扩展回文串的边界

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

	// 通过填充构造奇数长度的回文串
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
		// 递推求解
		if i < maxRight {
			R[i] = minInt(R[2*maxPos-i], maxRight-i+1)
		} else {
			R[i] = 1
		}

		// 向两边扩展
		for i+R[i] < n && i-R[i] >= 0 && sb[i-R[i]] == sb[i+R[i]] {
			R[i]++
		}

		// 维护更新 maxPos 和 maxRight
		if maxRight < i+R[i]-1 {
			maxRight = i + R[i] - 1
			maxPos = i
		}

		// 统计最长回文半径
		maxLen = maxInt(maxLen, R[i])
	}
	return maxLen - 1 // 最长回文子串的长度为 maxLen - 1
}
