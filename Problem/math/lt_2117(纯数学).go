package main

import (
	"fmt"
	"math"
)

// 题目链接: https://leetcode-cn.com/problems/abbreviating-the-product-of-a-range/
// 1. math.Log10 求解乘积的前几位数字
// 2. c2, c5 乘积末尾 0 的个数

func count(a, b int) (int, int) {
	cnt := 0
	for a > 0 && a%b == 0 {
		cnt++
		a /= b
	}
	return a, cnt
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abbreviateProduct(left int, right int) string {
	c2, c5, sum, product, eps, l5 := 0, 0, 0.0, 1, 1e-9, 1
	for i := left; i <= right; i++ {
		x1, x2 := count(i, 2)
		c2 += x2

		x1, x2 = count(x1, 5)
		c5 += x2

		sum += math.Log10(float64(x1))
		product *= x1 // 溢出不需要考虑
		l5 = l5 * x1 % 100000
	}

	c0 := minInt(c2, c5)
	for i := 0; i < c2-c0; i++ {
		sum += math.Log10(2.0)
		product *= 2
		l5 = l5 * 2 % 100000
	}
	for i := 0; i < c5-c0; i++ {
		sum += math.Log10(5.0)
		product *= 5
		l5 = l5 * 5 % 100000
	}

	if sum > 10.0 {
		sum = sum - math.Floor(sum)
		f5 := int(math.Floor((math.Pow(10.0, sum) + eps) * 10000.0)) // eps 浮点数误差
		return fmt.Sprintf("%d...%05de%d", f5, l5, c0)               // %05d
	}

	return fmt.Sprintf("%de%d", product, c0)
}
