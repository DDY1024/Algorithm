package main

import (
	"math"
)

// x^y % n
func powMod(x, y, n int) int {
	x %= n
	r := 1
	for y > 0 {
		if y&1 > 0 {
			r = r * x % n
		}
		x = x * x % n
		y >>= 1
	}
	return r
}

// x * y % n
func mulMod(x, y, n int) int {
	x %= n
	r := 0
	for y > 0 {
		if y&1 > 0 {
			r = (r + x) % n
		}
		x = x * 2 % n
		y >>= 1
	}
	return r
}

// 离散对数问题
// a^k = b (mod m)，求解 k 使得上述同余方程成立, 其中 gcd(a, m) = 1
// 求解该问题通常采用 Baby-step giant-step algorithm 算法
// 算法原理参考：https://www.geeksforgeeks.org/discrete-logarithm-find-integer-k-ak-congruent-modulo-b/
// 如果 k 在 [0, m) 区间内不存在解，则该同余方程不存在解。离散对数的求解算法便是将 k 换一种表示方式，通过 i * n - j
// 其中 i 在 [1, n) 区间内， j 在 [0, n) 区间内, n = ceil(sqrt(m))
// 求解 a^k = b (mod m)
// 注意: a 和 m 必须互素，即 gcd(a, m) = 1
func discreteLog(a, b, m int) int {
	// n := int(math.Sqrt(float64(m))) + 1
	n := int(math.Ceil(math.Sqrt(float64(m))))
	mark := map[int]int{}
	for i := 1; i <= n; i++ { // LHS: a^(i*n) % m
		x := powMod(a, i*n, m)
		if _, ok := mark[x]; !ok { // 注意: 此处要进行判重操作
			mark[x] = i
		}
	}
	for i := 0; i < n; i++ {
		cur := powMod(a, i, m) * b % m // a^j * b % m , 如果此处发生整数溢出的情况，则乘法操作的取余要单独处理
		if _, ok := mark[cur]; ok {
			res := mark[cur]*n - i
			if res < m { // < m 判断必须存在
				return res
			}
		}
	}
	return -1
}
