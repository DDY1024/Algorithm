package main

// 参考资料: https://zhuanlan.zhihu.com/p/108422764
// 欧拉函数: phi(x) 为 <=x 且与 x 互质的正整数个数，其中phi(1) = 1
// 欧拉函数为积性函数，其证明参考: https://blog.csdn.net/summonlight/article/details/51967425
// 假设 x 素因子分解为 x = p1^x1 * p2^x2 * ... + pk^xk，则欧拉函数 phi(x) 求解公式为
// phi(x) = phi(p1^x1) * phi(p2^x2) * ... * phi(pk^xk)
// phi(x) = p1^(x1-1)*(p1-1) * ... * pk^(xk-1)*(pk-1)
// phi(x) = x*((p1-1)/p1) * ... * ((pk-1)/pk)

// 常见结论:
// phi(p^n) = (p^(n-1))*(p-1)，其中 p 为素数，因为与 p^n 不互质的正整数只可能是 p, 2*p, 3*p, ..., p^(n-1)*p
// 如果a|x，则 phi(ax) = a*phi(x), 其中 a|x 表示 x 被 a 整除，x 是 a 的倍数，a 是 x 的因子
// 如果gcd(a,b)=1, 则phi(a*b)=phi(a)*phi(b)

// 1. 直接求解欧拉函数
func calcPhi(x int) int {
	res := x
	for i := 2; i*i <= x; i += 2 {
		if x%i == 0 {
			res = res / i * (i - 1)
		}
		for x%i == 0 {
			x /= i
		}
		if i == 2 {
			i--
		}
	}
	if x > 1 {
		res = res / x * (x - 1)
	}
	return res
}

// 2. 素数筛法求解欧拉函数
// 复杂度: O(n * logn * logn)
func solvePhi(n int) {
	phi := make([]int, n+1)
	for i := 1; i <= n; i++ {
		phi[i] = i
	}
	for i := 2; i <= n; i += 2 {
		if phi[i] == i { // 素数
			for j := i; j <= n; j += i {
				phi[j] = phi[j] / i * (i - 1)
			}
		}
		if i == 2 {
			i--
		}
	}
}

// 3. 线性递推求解欧拉函数
// 算法复杂度 O(n)
func solveFaster(n int) {
	phi := make([]int, n+1)
	isNP := make([]bool, n+1)
	primes := make([]int, 0, n/2)

	phi[1] = 1
	for i := 2; i <= n; i++ {
		if !isNP[i] { // 素数
			primes = append(primes, i)
			phi[i] = i - 1
		}

		for j := 0; j < len(primes); j++ {
			if primes[j]*i > n {
				break
			}
			isNP[primes[j]*i] = true
			if i%primes[j] == 0 {
				phi[primes[j]*i] = phi[i] * primes[j]
				break
			} else {
				phi[primes[j]*i] = phi[primes[j]] * phi[i]
			}
		}
	}
}
