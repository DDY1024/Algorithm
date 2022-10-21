package main

// 题目链接：
// 解题报告：https://leetcode.cn/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		// 结论: 按照 mid 划分，其中一半数组必然是有序的，有题目性质可知，无序的情况只会发生在一侧
		// 这样我们可以利用有序的一侧，进行区间排除，达到一个二分的效果
		if nums[l] <= nums[mid] { // 注意此处 l 和 mid 可能重合，应该是 <=
			if nums[l] <= target && nums[mid] > target { // 有序包含判断
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // 右侧有序
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
