package main

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

func manacher(s string) int {
	var (
		R         []int // 回文半径
		maxPos    int   // 当前最长回文子串的中心位置
		maxRight  int   // 当前最长回文子串能够触及到的最右位置
		maxRadius int
		bs        []byte
		n         int
	)

	for i := 0; i < len(s); i++ {
		bs = append(bs, '*')
		bs = append(bs, s[i])
	}
	bs = append(bs, '*')

	n = len(bs)
	R = make([]int, n)
	maxPos, maxRight, maxRadius = 0, 0, 1

	for i := 0; i < n; i++ {
		if i < maxRight {
			R[i] = minInt(R[2*maxPos-i], maxRight-i+1) // 递推求解
		} else {
			R[i] = 1
		}

		// 扩展半径
		for i+R[i] < n && i-R[i] >= 0 && bs[i-R[i]] == bs[i+R[i]] {
			R[i]++
		}

		if maxRight < i+R[i]-1 {
			maxRight = i + R[i] - 1
			maxPos = i
		}

		maxRadius = maxInt(maxRadius, R[i])
	}
	return maxRadius - 1
}
