package main

import (
	"fmt"
	"math"
	"strconv"
)

// 题目链接: https://leetcode-cn.com/problems/find-the-closest-palindrome/
//
// 详细的解题报告可以参考这篇文章: https://leetcode-cn.com/problems/find-the-closest-palindrome/solution/gong-shui-san-xie-tan-xin-fen-xi-shang-x-vtr6/
// 其实总体思想主要是:
// 1. 贪心思想，优先选择高位进行变动，这样变动后产生的偏移最小，即从前半部分 不变、+1、-1 中选取，构造后半部分
// 2. 特殊 case 对于像 99, 100 这种特殊 case 如果按照方法 1 进行构造会得出错误的结果，因此针对这类情况，我们
// 引入两个边界对我们的结果进行相应的补充。假设原整数的长度为 n，我们引入 10^(n-1)-1 和 10^n+1 进行额外补充
//
// 综上所述，我们最终结果的选取便是从这些中进行最优选择

func nearestPalindromic(n string) string {
	m := len(n)

	// 候选集增加边界场景：10^(m-1) - 1, 10^m + 1
	// 9...9, 10...1
	candidates := []int{int(math.Pow10(m-1)) - 1, int(math.Pow10(m)) + 1}

	// 奇数长度: [0, (m+1)/2)
	// 偶数长度: [0, (m+1)/2)
	selfPrefix, _ := strconv.Atoi(n[:(m+1)/2])
	for _, x := range []int{selfPrefix - 1, selfPrefix, selfPrefix + 1} {
		y := x

		// 奇数长度去除中间位构造剩余部分
		if m&1 == 1 {
			y /= 10
		}

		for ; y > 0; y /= 10 {
			x = x*10 + y%10
		}
		candidates = append(candidates, x)
	}

	ans := -1
	selfNumber, _ := strconv.Atoi(n)
	for _, candidate := range candidates {
		if candidate != selfNumber {
			if ans == -1 ||
				abs(candidate-selfNumber) < abs(ans-selfNumber) ||
				abs(candidate-selfNumber) == abs(ans-selfNumber) && candidate < ans {
				ans = candidate
			}
		}
	}
	return strconv.Itoa(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(math.Pow10(1) - 1)
}
