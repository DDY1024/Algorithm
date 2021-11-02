package main

import "fmt"

const (
	maxn = 1010
)

// Gcd 最大公约数
// 算法复杂度: O(log(max(a,b)))
// Gcd(n, n+1) = 1 即相邻的两个整数是互质的
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// 浮点数 最大公约数
// fgcd(a, b)
// a = p1^a1 * p2^a2 * ... * pk^ak
// b = p1^b1 * p2^b2 * ... * pk^bk
// gcd(a, b) = p1^min(a1,b1) * p2^min(a2,b2) * ... * pk^min(ak,bk)
// lcm(a, b) = p1^max(a1,b1) * p2^max(a2,b2) * ... * pk^max(ak,bk)
// 最终取 <= min(a, b) 的最大浮点数即可，小数点左移多少位的问题
// https://www.geeksforgeeks.org/program-find-gcd-floating-point-numbers/

// Lcm 最小公倍数
// Lcm(a1/b1, a2/b2, ..., ak/bk) = Lcm(a1, a2, ..., ak) / Gcd(b1, b2, ..., bk)
// Lcm(1, 2, ..., n) 另一种求解方法
// a = p1^x1 * p2^x2 * ... * pk^xk
// b = p1^y1 * p2^y2 * ... * pk^yk
// Lcm(a, b) = p1^max(x1, y1) * p2^max(x2,y2) * ... * pk^max(xk, yk)
// 因此，对于求解 Lcm(1, 2, ..., n) 我们只需要枚举 [1, n] 范围内的所有素数，并且针对该素数求解最大的 x 使得 p^x <= n
// 最终结果为 ans = 2^x1 * 3^x2 ... *
// https://www.geeksforgeeks.org/lcm-first-n-natural-numbers/
//
//
// Maximum sum of distinct numbers with LCM as N
// https://www.geeksforgeeks.org/maximum-sum-distinct-number-lcm-n/
// 利用 Lcm 求解的素因子性质, maxSum = (1 + p1^1 + ... + p1^x1) * (1 + p2^x2 + ... p2^x2) * ...
// 其实上面公式所求解的即 N 的所有约数的和(包括 N)
//
//
// Given GCD G and LCM L, find number of possible pairs (a, b)
// https://www.geeksforgeeks.org/given-gcd-g-lcm-l-find-number-possible-pairs-b/
// 求解出所有 (a, b) 对，满足 gcd(a, b) = x 且 lcm(a, b) = y
// x = p1^x1 * p2^x2 * ... * pk^xk
// y = p1^y1 * p2^y2 * ... * pk^yk
// 如果 x1 < y1，则方案数 * 2；如果 x1 = y1，则方案数 * 1
func Lcm(a, b int) int {
	return a / Gcd(a, b) * b
}

// 最大公约数和最小公倍数的一些性质
// a * b = Lcm(a, b) * Gcd(a, b)
// Gcd(a, b, c) = Gcd(Gcd(a, b), c)
// Lcm(a, b, c) = Lcm(Lcm(a, b), c)
// Gcd(Lcm(x,y),Lcm(x,z)) = Lcm(x,Gcd(y,z))

// https://www.cnblogs.com/frog112111/archive/2012/08/19/2646012.html
// 扩展欧几里得算法
// ax + by = gcd(a, b)
// 初始解 x0, y0
// 通解公式: x = x0 + (b/d) * t, y = y0 - (a/d)*t
// ax + by = c --> c % gcd(a, b) == 0
// 初始解: x0 = x0 * c / d, y0 = y0 * c/d
// 通解公式: x = x0 + (b/d)*t, y = y0 - (a/d)*t
func extendGcd(a, b int, x *int, y *int) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	d := extendGcd(b, a%b, y, x)
	*y -= *x * (a / b)
	return d
}

// ax = b (mod n)
// 方程转化: ax - ny = b --> 扩展欧几里得
// 最小整数解: ((x * b / d) % n + n) % n
func modEquation(a, b, n int) int {
	var d, e, x, y int
	d = extendGcd(a, n, &x, &y)
	if b%d > 0 {
		return -1 // 不存在解
	}
	// 存在 d 个 模 n 不同余的解
	e = (x * b / d) % n
	for i := 0; i < d; i++ {
		fmt.Println(e + i*(n/d))
	}
	return e // 暂且返回第一个解即可
}

// 模逆元
// ax = 1 (mod n)
func inv(a, n int) int {
	var d, x, y int
	d = extendGcd(a, n, &x, &y)
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
//
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

func main() {
	fmt.Println("Yes")
}
