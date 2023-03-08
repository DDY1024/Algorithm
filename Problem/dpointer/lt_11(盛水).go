package main

// 题目链接：https://leetcode.cn/problems/container-with-most-water/
//
// 解题思路
// 		双指针

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	n := len(height)
	i, j, ret := 0, n-1, 0
	for i < j {
		ret = maxInt(ret, minInt(height[i], height[j])*(j-i))
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return ret
}
