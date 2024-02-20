package main

// 题目链接：https://leetcode.cn/problems/container-with-most-water/

func maxArea(height []int) int {
	ret, i, j := 0, 0, len(height)-1
	for i < j {
		ret = max(ret, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ret
}
