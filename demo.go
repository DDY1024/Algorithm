package main

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

func waysToBuildRooms(prevRoom []int) int {
	n, mod := len(prevRoom), int(1e9+7)

	// 预处理
	factor := make([]int, n+1)
	invFactor := make([]int, n+1)
	factor[0] = 1
	invFactor[0] = 1
	for i := 1; i <= n; i++ {
		factor[i] = factor[i-1] * i % mod
		invFactor[i] = powMod(factor[i], mod-2, mod)
	}

	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		if prevRoom[i] >= 0 {
			adj[prevRoom[i]] = append(adj[prevRoom[i]], i)
		}
	}

	// 计算组合数
	// var C = func(n, m int) int {
	// 	return factor[n] * invFactor[m] % mod * invFactor[n-m] % mod
	// }

	var solve func(u int) (int, int)
	solve = func(u int) (int, int) {
		if len(adj[u]) == 0 {
			return 1, 1
		}

		ans, childTotal := 1, 0
		for _, v := range adj[u] {
			ca, cb := solve(v)
			ans = ans * ca % mod * invFactor[cb] % mod
			childTotal += cb
		}
		ans = ans * factor[childTotal] % mod
		return ans, childTotal + 1
	}

	ans, _ := solve(0)
	return ans
}

func main() {

}
