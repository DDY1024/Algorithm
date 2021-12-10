package main

import "fmt"

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
