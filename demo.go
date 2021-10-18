package main

import "fmt"

//
//
// 欠题
// TODO: 后缀数组
// https://leetcode-cn.com/problems/longest-common-subpath/
//

// 3 ^ 5 = 243
// 进行预处理，优化状态转移的复杂度
//
//
//
func colorTheGrid(m int, n int) int {

	if m > n {
		m, n = n, m
	}

	// 3 种颜色
	var isOk = func(x, n int) bool {
		pBit := -1
		for i := 0; i < n; i++ {
			if pBit == x%3 {
				return false
			}
			pBit = x % 3
			x /= 3
		}
		return true
	}

	var isAdj = func(x, y, n int) bool {
		for i := 0; i < n; i++ {
			if x%3 == y%3 {
				return false
			}
			x /= 3
			y /= 3
		}
		return true
	}

	var power = func(a, b int) int {
		ret := 1
		for b > 0 {
			if b&1 > 0 {
				ret = ret * a
			}
			a = a * a
			b >>= 1
		}
		return ret
	}

	limit := power(3, m)
	okS := make([]int, 0, limit)
	for i := 0; i < limit; i++ {
		if isOk(i, m) {
			okS = append(okS, i)
		}
	}

	adj := make([][]int, limit)
	for _, i := range okS {
		adj[i] = make([]int, 0, len(okS))
		for _, j := range okS {
			if isAdj(i, j, m) {
				adj[i] = append(adj[i], j)
			}
		}
	}

	// 状态压缩动态规划
	dp, mod := make([][]int, n), int(1e9+7)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, limit)
	}
	for _, i := range okS {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for _, j := range okS {
			for _, k := range adj[j] {
				dp[i][j] = (dp[i][j] + dp[i-1][k]) % mod
			}
		}
	}

	ans := 0
	for _, i := range okS {
		ans = (ans + dp[n-1][i]) % mod
	}
	return ans
}

// 5 * 1000 = 5000
// dp[][]
//
// (i, j)    (i, j+1)
// (i+1, j)  (i+1, j+1)
//
//
// (a, b, c)
// a, b, c
//
//
//
//
//
//   x
// x

//
// 5 * 1000
// 5 * 1000 * 3 * 3
// dp[i][j][3][3]
// 5000 * 3 * 3 = 45000

func main() {
	fmt.Println("hello, world!")
}

//
//
//
//
// a1, b1, c1, d1, e1, f1, g1, h1
// a2, b2, c2, d2, e2, f2, g2, h2
// a3, b3, c3, d3, e3, f3, g3, h3
//
//
