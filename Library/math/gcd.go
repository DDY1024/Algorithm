package main

import "fmt"

// 最大公约数 O(log(max(a,b)))
//
//	a. gcd(x, x + 1) = 1
//	b. fgcd 参考 https://www.geeksforgeeks.org/program-find-gcd-floating-point-numbers/
//	c. 素因子分解法求解最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 最小公倍数
// a. lcm(a1/b1, a2/b2, ..., ak/bk) = lcm(a1, a2, ..., ak) / gcd(b1, b2, ..., bk)  --> 分子大，分母小
// b. 素因子分解法求解最小公倍数
// c. https://www.geeksforgeeks.org/maximum-sum-distinct-number-lcm-n/ 一些整数 lcm 为 n，这些整数的最大和为 (1+...+p1^x1)*(1+...+p2^x2) ... *(1+...+pk^xk)，其中 n = p1^x1 * ... * pk^xk
// d. 每个【正整数】均可被唯一的素因子分解表示
// e. https://www.geeksforgeeks.org/given-gcd-g-lcm-l-find-number-possible-pairs-b/，求解出所有 (a, b) 对，满足 gcd(a, b) = x 且 lcm(a, b) = y
//
//			gcd(a, b) = p1^min(a1,b1) * p2^min(a2,b2) * ... * pk^min(ak,bk) = p1^x1 * p2^x2 * ... * pk^xk
//	        lcm(a, b) = p1^max(a1,b1) * p2^max(a2,b2) * ... * pk^max(ak,bk) = p1^y1 * p2^y2 * ... * pk^yk
//		由上述素因子分解表达式可知，当 xk != yk 时，方案数 * 2；当 xk = yk 时，方案数 * 1
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

// 素因子分解
// 	a = p1^a1 * p2^a2 * ... * pk^ak
// 	b = p1^b1 * p2^b2 * ... * pk^bk
// 	gcd(a, b) = p1^min(a1,b1) * p2^min(a2,b2) * ... * pk^min(ak,bk)
// 	lcm(a, b) = p1^max(a1,b1) * p2^max(a2,b2) * ... * pk^max(ak,bk)

// 扩展欧几里德算法
//
//		ax + by = c，其中 d = gcd(a, b)
//		 1. 初始解 x0, y0
//	     2. 通解公式：x = x0 + (b/d)*t, y = y0 - (a/d)*t
//	方程 ax + by = c，存在解的前提是 gcd(a,b)|c
//		x0 = x0*x/d，y0 = y0*c/d
//	通解公式: x = x0 + (b/d)*t, y = y0 - (a/d)*t
func extendGCD(a, b int, x *int, y *int) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	d := extendGCD(b, a%b, y, x)
	*y -= *x * (a / b)
	return d
}

func extendGCD2(a, b int, x *int, y *int) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	d, x0, y0 := extendGCD2(b, a%b, x, y), *x, *y
	*x = y0
	*y = x0 - (a/b)*y0
	return d
}

// 同余方程
//
//	ax = b (mod n)   -->   a*x - n*y = b
func modEquation(a, b, n int) int {
	var d, e, x, y int
	d = extendGCD(a, n, &x, &y)
	if b%d > 0 {
		return -1
	}

	// 存在 d 个模 n 不同余的解
	e = (x * b / d) % n
	// e = ((x*b/d)%n + n) % n
	for i := 0; i < d; i++ {
		fmt.Println(e + i*(n/d))
	}

	// 最小整数解：((x * b / d) % n + n) % n
	return e
}

// 模逆元（特殊同余方程）
//
//	a*x = 1 (mod n)
func inv(a, n int) int {
	var d, x, y int
	d = extendGCD(a, n, &x, &y)
	if d != 1 {
		return -1
	}

	return (x%n + n) % n
}

// 大整数取余
func bigIntMod(b []byte, mod int) int {
	ret := 0
	for i := 0; i < len(b); i++ {
		ret = (ret*10 + int(b[i]-'0')) % mod
	}
	return ret
}
