package main

// 题目链接：https://leetcode.com/problems/trapping-rain-water/

// 1. 常规解法
func trap(height []int) int {
	n := len(height)
	lmax := make([]int, n)
	rmax := make([]int, n)
	for i := 0; i < n; i++ {
		lmax[i] = height[i]
		if i-1 >= 0 {
			lmax[i] = max(lmax[i], lmax[i-1])
		}
	}
	for i := n - 1; i >= 0; i-- {
		rmax[i] = height[i]
		if i+1 < n {
			rmax[i] = max(rmax[i], rmax[i+1])
		}
	}

	ret := 0
	for i := 1; i < n-1; i++ {
		ret += max(0, min(lmax[i], rmax[i])-height[i])
	}
	return ret
}

// 2. 双指针解法
func trap1(height []int) int {
	n := len(height)
	l, r, lmax, rmax, ret := 0, n-1, 0, 0, 0
	for l < r {
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])
		// min{ lmax(i),  rmax(i) } - height(i)
		if height[l] < height[r] { // 谁小，谁挪动
			ret += lmax - height[l] // lmax 必然 <= rmax
			l++
		} else {
			ret += rmax - height[r] // rmax 必然 <= lmax
			r--
		}
	}
	return ret
}
