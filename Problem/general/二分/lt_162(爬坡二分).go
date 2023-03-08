package main

// 题目链接：https://leetcode.cn/problems/find-peak-element/
//
// 解题思路
// 	   最终的解题思路恰好对应到二分查找上，可以实现在 O(logN) 复杂度内解决该问题
// https://leetcode.cn/problems/find-peak-element/solution/xun-zhao-feng-zhi-by-leetcode-solution-96sj/

//
// 寻找峰值元素

func findPeakElement(nums []int) int {

	n := len(nums)

	var get = func(idx int) int {
		if idx < 0 || idx >= n {
			return -(1 << 40)
		}
		return nums[idx]
	}

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

	return -1 // 由于 nums[i] != nums[i+1]，因此可以确定答案是肯定存在的
}
