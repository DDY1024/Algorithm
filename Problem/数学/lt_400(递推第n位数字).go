package main

// 题目链接: https://leetcode-cn.com/problems/nth-digit/
// 解题思路: 递推
// 1 ~ 9: 9
// 10 ~ 99: 90 * 2
// 100 ~ 999: 900 * 3
// 1000 ~ 9999: 9000 * 4
// 10000 ~ 99999: 90000 * 5
// 100000 ~ 999999: 900000 * 6
// ...
// 100000000 ~ 999999999: 900000000 * 9
// 最多求解第 2^31 - 1 位
// 我们可以通过递推计算出第 x 位落在哪个整数区间内，由于每个整数区间内的整数位数是相同的，因此我们可以直接根据偏移量计算出是整数 x 的从右往左数的第 y 位
func findNthDigit(n int) int {
	sum, tmp, base := make([]int, 10), 9, make([]int, 10)
	base[1] = 1
	for i := 1; i < 10; i++ {
		sum[i] = sum[i-1] + tmp*i
		tmp *= 10
		if i > 1 {
			base[i] = base[i-1] * 10
		}
	}

	idx := -1
	for i := 1; i < 10; i++ {
		if n <= sum[i] {
			idx = i
			break
		}
	}

	// 返回 x 十进制表示从低到高的第 y 位
	var get = func(x, y int) int {
		for i := 1; i < y; i++ {
			x /= 10
		}
		return x % 10
	}

	remain := n - sum[idx-1]
	x, y := base[idx]+remain/idx-1, remain%idx
	if y == 0 {
		y = 1
	} else {
		x++
		y = idx - y + 1
	}

	return get(x, y)
}
