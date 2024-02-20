package main

// 题目链接：https://leetcode.cn/problems/search-in-rotated-sorted-array/?envType=study-plan-v2&envId=top-100-liked
//
// 1. nums[l], nums[mid], nums[r]
// 		比较 nums[mid] 和 nums[r]；target 和 nums[r] 的大小关系

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		// nums[l], nums[mid], nums[r]
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			if nums[mid] < nums[r] {
				if target <= nums[r] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < nums[r] {
				r = mid - 1
			} else {
				if target <= nums[r] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
		}
	}
	return -1
}

// 题目链接：https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/?envType=study-plan-v2&envId=top-100-liked
//
// 采样同样的二分方式，搜索旋转数组的最小值
func findMin(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		// nums[l], nums[mid], nums[r]
		// if nums[l] < nums[mid] < nums[r] {
		//     return nums[l]
		// }

		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}
