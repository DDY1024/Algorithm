package main

// 题目链接：https://leetcode.com/problems/trapping-rain-water/

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	n := len(height)
	left, right, leftMax, rightMax, ret := 0, n-1, 0, 0, 0
	for left < right {
		leftMax = maxInt(leftMax, height[left])
		rightMax = maxInt(rightMax, height[right])
		if height[left] < height[right] {
			ret += leftMax - height[left]
			left++
		} else {
			ret += rightMax - height[right]
			right--
		}
	}
	return ret
}
