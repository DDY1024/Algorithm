package main

// https://leetcode.com/contest/biweekly-contest-44/problems/count-ways-to-make-array-with-product/
//
// 解题思路:
// 1. 素因子分解
// 2. 组合数学计数（插板法）
//
// 一个整数的素因子分解即一系列不可再分解的素数因子的乘积，将这些素因子分配到不同的篮子里面，对应的组合方案数采用插板法进行求解
// 解题思路可以参考这篇帖子：https://leetcode-cn.com/problems/count-ways-to-make-array-with-product/solution/guan-yu-zu-he-shu-de-ji-suan-by-jacky_50-d8w6/
//
// 采用什么样的组合方案数来求解，想到这个方案还是蛮困难，仔细想想。
//
// 素因子分解：最小的不可分解的因子表达式
//
// 素因子分解：最小的不可分解的因子表达式
//

func powerMod(a, b, c int) int {
	ret := 1
	for b > 0 {
		if b&1 > 0 {
			ret = ret * a % c
		}
		a = a * a % c
		b >>= 1
	}
	return ret
}

func combine(n, m, mod int, facMod []int) int {
	ret := facMod[n]
	ret = ret * powerMod(facMod[m], mod-2, mod) % mod
	ret = ret * powerMod(facMod[n-m], mod-2, mod) % mod
	return ret
}

func waysToFillArray(queries [][]int) []int {
	n, mod := len(queries), int(1e9+7)
	facMod := make([]int, 20000+1)
	facMod[0] = 1
	// n + k 最大范围为 20000，因此我们预先计算出 20000 范围内阶乘的取余值即可
	for i := 1; i <= 20000; i++ {
		facMod[i] = facMod[i-1] * i % mod
		// facMod[i] = facMod[i-1] * i % mod
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ni, ki, r := queries[i][0], queries[i][1], 1
		for p := 2; p*p <= ki; p += 2 {
			if ki%p == 0 {
				cnt := 0
				for ki%p == 0 {
					cnt++
					ki /= p
				}
				r = r * combine(cnt+ni-1, ni-1, mod, facMod) % mod
			}
			if p == 2 {
				p--
			}
		}
		if ki > 1 {
			r = r * combine(1+ni-1, ni-1, mod, facMod) % mod
		}
		ans[i] = r
		//ans = append(ans, r)
	}
	return ans
}

// 转化成素因子个数分配到哪个桶中的问题，C(x+n-1, n-1) -----> C(x+n-1, n-1)，最终每个桶中的数便是素因子相乘的结果
