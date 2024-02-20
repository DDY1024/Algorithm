package main

// 题目链接: https://leetcode.cn/problems/number-of-different-subsequences-gcds/description/

// 解题报告:
// https://leetcode.cn/problems/number-of-different-subsequences-gcds/solutions/2061079/ji-bai-100mei-ju-gcdxun-huan-you-hua-pyt-get7/
// https://leetcode.cn/problems/number-of-different-subsequences-gcds/solutions/2060230/xu-lie-zhong-bu-tong-zui-da-gong-yue-shu-ha1j/

// 题目总结
// 1. 数组的任意子序列的个数是 2^n，遇到这类题目时，如果题目给的数据范围 n 较大，必然不能采取枚举子序列的方式求解
// 		本题中 n <= 10^5，因此我们需要换种思考方式
// 		从最大公约数的结果出发，进行求解
// 2. 最大公约数性质
// 		gcd(a, b) = d1
//      gcd(a, b, c) = d2 = gcd(d1, c)
//    因此，多个数的最大公约数是单调递减性质；因此，可知
//      gcd(a, b) = d1
//      gcd(a, b, x*d1) = gcd(d1, x*d1) = d1
//    回到本题，我们只需要求解数组中所有 x 倍数的数的最大公约数是否等于 x 即可，不需要单独考虑某几个 x 的倍数
// 3. 枚举复杂度（调和级数复杂度）
//    x/1 + x/2 + ... + x/x 是 xlog(x) 复杂度
//    1/1 + 1/2 + ... + 1/x 是 log(x) 复杂度
//    因此，我们在枚举最大公约数时，计算的最终复杂度为 O(max*log(max)) 级别的，符合题目要求

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func countDifferentSubsequenceGCDs(nums []int) int {
	n, maxNum := len(nums), 0
	mark := make(map[int]bool, n)
	for i := 0; i < n; i++ {
		mark[nums[i]] = true
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	// var gcd func(a, b int) int
	// gcd = func(a, b int) int {
	// 	if b == 0 {
	// 		return a
	// 	}
	// 	return gcd(b, a%b)
	// }

	diffCnt := 0
	for i := 1; i <= maxNum; i++ {
		d := 0
		for j := i; j <= maxNum; j += i {
			if mark[j] {
				d = gcd(d, j)
			}
		}
		if d == i {
			diffCnt++
		}
	}
	return diffCnt
}
