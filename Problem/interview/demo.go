package main

import "sort"

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

func searchRange(nums []int, target int) []int {
	n := len(nums)
	idx := sort.SearchInts(nums, target)
	if idx >= n || nums[idx] != target {
		return []int{-1, -1}
	}
	return []int{idx, sort.SearchInts(nums, target+1) - 1}
}
