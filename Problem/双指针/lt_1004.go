package main

import "fmt"

// 双指针思想博大精深，可以大大简化问题的处理流程
// https://leetcode.com/problems/max-consecutive-ones-iii/
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func longestOnes(A []int, K int) int {
	n, left, right, c0, ans := len(A), 0, 0, 0, 0
	for right < n {
		if A[right] == 0 {
			c0++
		}
		for c0 > K {  // make sure
			if A[left] == 0 {
				c0--
			}
			left++
		}
		ans = maxInt(ans, right - left + 1)
		right++
	}
	return ans
}

func main() {
	fmt.Println(longestOnes([]int{1,1,1,0,0,0,1,1,1,1,0}, 2))
	fmt.Println(longestOnes([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3))
}