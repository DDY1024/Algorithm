package main

// 题目链接: https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation/

// 题目大意
// 		偶数大小的数组给定索引下标的转换规则，求解经过多少次转换，使得原数组变为原样

// 解题思路（一）
// 意就是要求一个排列重复作用几次后能够回到初始状态，用数学语言描述就是求一个置换的阶（周期）
// https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation/solution/by-vclip-iunv/
//	任意一个置换可以分解为多个独立的不相交轮换，而置换的阶即为这些不相交轮换周期的最小公倍数

// 题目下标转换关系
//   1. 2 * i	[0, n/2)
//   2. 2 * i - n + 1 	[n/2, n)

func reinitializePermutation(n int) int {
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}

	var lcm = func(a, b int) int {
		return a / gcd(a, b) * b
	}

	var convert = func(i int) int {
		if i < n/2 {
			return 2 * i
		}
		return 2*i - n + 1
	}

	mark := make([]bool, n)
	ret := 1
	for i := 0; i < n; i++ {
		if mark[i] { // 轮换内所有位置的转换周期是相同的
			continue
		}

		step, pos := 0, i
		for !mark[pos] {
			mark[pos] = true
			step++
			pos = convert(pos)
		}
		ret = lcm(ret, step) // 所有轮换周期的最小公倍数
	}
	return ret
}

// 解题思路（二）
// https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation/solutions/2051628/huan-yuan-pai-lie-de-zui-shao-cao-zuo-bu-d9cn/
//
// 下标转换关系
//   1. 2 * i	--> [0, n/2)
//   2. 2 * i - (n - 1) --> [n/2, n)
// 通过上述坐标转换关系，实际上一次转换操作可以转化为 2*i mod (n-1)
// 因此，对于每个 i 求解出最小的 k 使得 2^k * i = i (mod n-1) 即可
// 最终，问题转化为 2^k = 1 (mod n-1)
// 由于 gcd(2, n-1) 互质，因此由 【欧拉定理】 可知，k 的解在 [1, n-1) 范围内，枚举即可

func reinitializePermutation2(n int) int {
	if n == 2 {
		return 1
	}

	// 2^k = 1 (mod n-1)，最小的 k
	a, k := 2, 1
	for a != 1 {
		k++
		a = a * 2 % (n - 1)
	}
	return k
}
