package main

import "fmt"

const maxn = 10010

// 1. 最大公约数
// 算法复杂度: O(log(max(a,b)))
// 常用结论/题目
// a. gcd(n,n+1)=1，即相邻整数是互质的
// b. 浮点数 fgcd: https://www.geeksforgeeks.org/program-find-gcd-floating-point-numbers/
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 2. 最小公倍数
// 常用结论/题目
// a. lcm(a1/b1, a2/b2, ..., ak/bk) = lcm(a1, a2, ..., ak)/gcd(b1, b2, ..., bk)
// b. 求解 lcm(1, 2, ..., n)，我们提供一种素因子分解的思路，枚举 [1,n] 范围内的所有素数，计算每个素数的最大幂次 x 使得 p^x <= n，最终结果便是 2^x1 * 3^x3 * ... * p^xp --> https://www.geeksforgeeks.org/lcm-first-n-natural-numbers/
// c. https://www.geeksforgeeks.org/maximum-sum-distinct-number-lcm-n/ 一堆不同整数的 lcm 为 n，求解这些整数的最大和。利用素因子分解的思路，n = p1^x1 * p2^x2 * ... * pk^xk，则这些整数的最大和为 (1+p1+p1^2+..+p1^x1) * ... * (1+pk^1+...+pk^xk)，利用排列组合的思想想想这是为什么？
// d. 每个正整数均可以被素因子分解唯一表示
// e. https://www.geeksforgeeks.org/given-gcd-g-lcm-l-find-number-possible-pairs-b/，求解出所有 (a, b) 对，满足 gcd(a, b) = x 且 lcm(a, b) = y
//    利用素因子分解的思想求解: gcd(a, b) = p1^min(a1,b1) * p2^min(a2,b2) * ... * pk^min(ak,bk) = p1^x1 * p2^x2 * ... * pk^xk
//                         	lcm(a, b) = p1^max(a1,b1) * p2^max(a2,b2) * ... * pk^max(ak,bk) = p1^y1 * p2^y2 * ... * pk^yk
//    由上述表示可知，当 x1 != y1 时，方案数乘 2（一个数大，一个数小，两种方式）; 当 x1 = y1 时，方案数乘 1
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

// 另外一种求解思路
// 通过素因子分解的方式来理解最大公约数和最小公倍数
// a = p1^a1 * p2^a2 * ... * pk^ak
// b = p1^b1 * p2^b2 * ... * pk^bk
// gcd(a, b) = p1^min(a1,b1) * p2^min(a2,b2) * ... * pk^min(ak,bk)
// lcm(a, b) = p1^max(a1,b1) * p2^max(a2,b2) * ... * pk^max(ak,bk)

// 最大公约数 && 最小公倍数 性质
// a * b = lcm(a, b) * gcd(a, b)
// gcd(a, b, c) = gcd(gcd(a, b),c)
// lcm(a, b, c) = lcm(lcm(a, b),c)
// gcd(lcm(x,y),lcm(x,z))=lcm(x,gcd(y,z)) --> 这条性质可以利用素因子分解的方式来求证 min、max 操作的交换运算

// 3. 扩展欧几里得算法  https://zhuanlan.zhihu.com/p/100567253
// 不定方程 ax + by = gcd(a,b), d = gcd(a,b)
// a. 初始解 x0, y0
// b. 通解公式: x = x0 + (b/d) * t, y = y0 - (a/d)*t
// 不定方程 ax + by = c，存在解的前提是 gcd(a,b)|c
// 初始解: x0 = x0 * c/d, y0 = y0 * c/d
// 通解公式: x = x0 + (b/d)*t, y = y0 - (a/d)*t
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

// 4. 扩展欧几里德算法求解同余方程
// 求解 ax = b (mod n) 方程转化 ax - ny = b，利用扩展欧几里得算法求解
// 最小整数解: ((x * b / d) % n + n) % n
func modEquation(a, b, n int) int {
	var d, e, x, y int
	d = extendGCD(a, n, &x, &y)
	if b%d > 0 {
		return -1 // 不存在解
	}

	// 存在 d 个模 n 不同余的解
	e = (x * b / d) % n
	for i := 0; i < d; i++ {
		fmt.Println(e + i*(n/d)) // 通解公式进行求解
	}

	// 最小整数解
	// ((x * b / d) % n + n) % n
	return e // 暂且返回第一个解即可
}

// 5. 模逆元
// 求解 ax = 1 (mod n) 作为一类特殊的同余方程求解

// 模逆元
// ax = 1 (mod n)
func inv(a, n int) int {
	var d, x, y int
	d = extendGCD(a, n, &x, &y)
	if d != 1 {
		return -1
	}
	return (x%n + n) % n
}

// 大整数取余运算
// a % b
func bigIntMod(a []byte, b int) int {
	c := 0
	for i := range a {
		c = (c*10 + int(a[i]-'0')) % b
	}
	return c
}

// 欧拉函数 f(x): 1 ~ x 中与 x 互质的数的个数
// f(x) = x * (1 - 1/p1) * (1 - 1/p2) * ... * (1-1/pk)
// f(1) = 1
// 欧拉函数应用实例
// 1 <= i, j <= N, 求解 SUM(gcd(i, j))
// F(i) : 所有 gcd(x, i) 的和
// 枚举 i 的所有约数 x, 计算 F(i) = SUM(x*phi(i/x))
// 对于一个正整数 N, 其约数枚举的时间复杂度为 O(sqrt(N))
func Phi(x int) int {
	res := x
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			res = res / i * (i - 1)
			for x%i == 0 {
				x /= i
			}
		}
	}
	if x > 1 {
		res = res / x * (x - 1)
	}
	return res
}

// 递推打表求解欧拉函数
var phi [maxn + 10]int

func calcEulerTable() {
	for i := 1; i < maxn; i++ {
		phi[i] = i
	}
	for i := 2; i < maxn; i += 2 {
		phi[i] /= 2
	}
	for i := 3; i < maxn; i += 2 {
		if phi[i] == i { // 素因子
			for j := i; j < maxn; j += i {
				phi[j] = phi[j] / i * (i - 1)
			}
		}
	}
}

// https://www.geeksforgeeks.org/summation-gcd-pairs-n/
// 求解 sigma(gcd(i,j)) 1 <= i, j <= n
// f(n) = sigma(gcd(i,n))
// s(n) = sigma(f(i))
var f [maxn]int
var s [maxn]int

func solveGcdSum() {
	for i := 1; i < maxn; i++ {
		f[i] = phi[i] // gcd 为 1
	}
	for i := 2; i*i <= maxn; i++ {
		f[i*i] += i * phi[i] // x = i * i
		for j, k := i*i+i, i+1; j < maxn; j, k = j+i, k+1 {
			f[j] += i*phi[k] + k*phi[i]
		}
	}
	s[1] = f[1]
	for i := 2; i < maxn; i++ {
		s[i] = s[i-1] + f[i]
	}
	return
}

// 利用欧拉函数求解另外一个例子
// https://www.geeksforgeeks.org/count-pairs-natural-numbers-gcd-equal-given-number/
// 上题中如果 [L, R] 区间范围很小，则我们可以直接枚举掉。如果区间范围很大，则此时我们需要利用欧拉函数
// 进行优化求解
