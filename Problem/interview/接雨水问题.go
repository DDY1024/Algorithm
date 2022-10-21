package main

// https://leetcode.cn/problems/trapping-rain-water/

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

func trap(height []int) int {
	n := len(height)
	left := make([]int, n)
	right := make([]int, n)

	left[0] = height[0]
	for i := 1; i < n; i++ {
		left[i] = maxInt(height[i], left[i-1])
	}

	right[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = maxInt(height[i], right[i+1])
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += minInt(left[i], right[i]) - height[i]
	}

	return ans
}
