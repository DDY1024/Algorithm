package main

// 欧拉函数：phi(x) --> <= x 且与 x 互质的正整数个, phi(1) = 1
// 计算公式
//
//	phi(x) = x * ((p1-1)/p1) * ((p2-1)/p2) * ... * ((pk-1)/pk)
//
// 积性函数：如果 gcd(a, b) = 1，则 phi(a*b) = phi(a) * phi(b)
func calc(x int) int {
	ret := x
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			ret = ret / i * (i - 1)
			for x%i == 0 {
				x /= i
			}
		}
	}
	if x > 1 {
		ret = ret / x * (x - 1)
	}
	return ret
}

// 素数筛法求解
func solve1(n int) {
	phi := make([]int, n+1)
	for i := 1; i <= n; i++ {
		phi[i] = i
	}

	for i := 2; i <= n; i++ {
		if phi[i] == i {
			for j := i; j <= n; j += i {
				phi[j] = phi[j] / i * (i - 1)
			}
		}
	}
}
