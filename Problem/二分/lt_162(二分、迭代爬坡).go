package main

import "math"

// 题目链接：https://leetcode.cn/problems/find-peak-element/?envType=daily-question&envId=Invalid%20Date
// 解题思路：https://leetcode.cn/problems/find-peak-element/solutions/998152/xun-zhao-feng-zhi-by-leetcode-solution-96sj/?envType=daily-question&envId=Invalid+Date

// 迭代爬坡 + 二分优化
func findPeakElement(nums []int) int {
	n := len(nums)

	get := func(idx int) int {
		if idx < 0 || idx >= n {
			return math.MinInt64
		}
		return nums[idx]
	}

	// 迭代爬坡：人往高处走，水往低处流，朝着更高的方向
	// i-1 < i < i+1
	// i-1 < i > i+1
	// i-1 > i > i+1
	// i-1 > i < i+1  这种情况向左走
	//
	//
	// idx := rand.Intn(n)
	// for !(get(idx) > get(idx-1) && get(idx) > get(idx+1)) {
	//     if get(idx) < get(idx+1) {
	//         idx++
	//     } else {
	//         idx--
	//     }
	// }

	// 迭代爬坡二分优化
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return mid
		}

		if get(mid) < get(mid+1) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
