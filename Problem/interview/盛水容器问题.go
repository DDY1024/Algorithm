package main

// 题目链接: https://leetcode.cn/problems/container-with-most-water/?favorite=2cktkvj

// 解题思路：单调性 + 二分

func maxArea(height []int) int {
	n := len(height)
	ret := 0
	larr := make([]int, 0)
	rarr := make([]int, 0)

	// 左 --> 右
	for i := 0; i < n; i++ {
		l, r, idx := 0, len(larr)-1, -1
		for l <= r {
			mid := l + (r-l)>>1
			if height[larr[mid]] >= height[i] {
				idx = larr[mid]
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if idx != -1 {
			ret = maxInt(ret, height[i]*(i-idx))
		}

		if len(larr) == 0 || height[larr[len(larr)-1]] < height[i] {
			larr = append(larr, i)
		}
	}

	// 右 --> 左
	for i := n - 1; i >= 0; i-- {

		l, r, idx := 0, len(rarr)-1, -1
		for l <= r {
			mid := l + (r-l)>>1
			if height[rarr[mid]] >= height[i] {
				idx = rarr[mid]
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if idx != -1 {
			ret = maxInt(ret, height[i]*(idx-i))
		}

		if len(rarr) == 0 || height[rarr[len(rarr)-1]] < height[i] {
			rarr = append(rarr, i)
		}
	}

	return ret
}

// 双指针解法
// https://leetcode.cn/problems/container-with-most-water/solution/sheng-zui-duo-shui-de-rong-qi-by-leetcode-solution/
// 首先初始区间为 [0, n-1]，取两者中的较小者进行移动，在这个过程不断更新最大值

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

func maxAreaTwo(height []int) int {
	n := len(height)
	l, r, ret := 0, n-1, 0
	for l < r {
		ret = maxInt(ret, minInt(height[l], height[r])*(r-l))
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return ret
}
