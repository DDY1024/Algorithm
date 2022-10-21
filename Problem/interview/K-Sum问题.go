package main

import "sort"

// 题意：寻找所有三个数之和为 0 且不重复的整数对
//
// K-Sum 问题及其演变

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	ans := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		l, r := i+1, n-1
		for l < r {
			if nums[l]+nums[r] == target {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l+1 < r && nums[l+1] == nums[l] {
					l++
				}
				l++

				for r-1 > l && nums[r-1] == nums[r] {
					r--
				}
				r--
			} else if nums[l]+nums[r] > target {
				r--
			} else {
				l++
			}
		}
	}

	return ans
}
