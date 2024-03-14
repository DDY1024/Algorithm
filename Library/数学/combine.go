package main

// 参考: https://blog.csdn.net/ACdreamers/article/details/8037918

// 情况一： 1 <= n, m <= 1000
// 直接根据杨辉三角公式递推求解: C[i][j] = C[i-1][j] + C[i-1][j-1]

// 情况二：n, m 较大，C(n, m) % p，其中 p 为素数
// Lucas 定理: Lucas(n, m) % p = Lucas(n/p, m/p) * C(n%p, m%p) % p

func powMod(a, b, p int) int {
	res := 1
	a %= p
	for b > 0 {
		if b&1 > 0 {
			res = res * a % p
		}
		a = a * a % p
		b >>= 1
	}
	return res
}

// 进一步优化方向：C(n, m) = n!/(m!*(n-m)!)，n <= 10^5 时，预处理出阶乘，不需要循环遍历了
// factor[i] = factor[i-1] * i % p
//
// 互质关系、费马小定理
// invFactor[i] = powMod(factor[i], p-2, p)
func C(n, m, p int) int {
	if m > n {
		return 0
	}
	ans := 1
	for i := 1; i <= m; i++ {
		a, b := (n+i-m)%p, i%p
		ans = ans * (a * powMod(b, p-2, p) % p) % p
	}
	return ans
}

func Lucas(n, m, p int) int {
	if m == 0 {
		return 1
	}
	return C(n%p, m%p, p) * Lucas(n/p, m/p, p) % p
}
