package main

// 参考资料: https://zhuanlan.zhihu.com/p/100051075

// 1. 朴素素数筛法
// 复杂度 O(n * logn * logn)
func doPrimeList(n int) {
	isNP := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if !isNP[i] { // 素数
			for j := i * i; j <= n; j += i {
				isNP[j] = true
			}
		}
	}
}

// 2. 欧拉筛、线性筛
// 复杂度 O(n)
// 工作原理: 让每一个合数只被其最小质因数筛到
func doLinearPrimeList(n int) {
	isNP := make([]bool, n+1)
	primes := make([]int, 0, n/2)
	for i := 2; i <= n; i++ {
		if !isNP[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p*i > n {
				break
			}
			isNP[p*i] = true
			// i % p == 0
			// i * next_p 的最小素因子是 p，而不是 next_p，不满足我们只能被最小素因子 p 筛到的要求
			if i%p == 0 {
				break
			}
		}
	}
}
