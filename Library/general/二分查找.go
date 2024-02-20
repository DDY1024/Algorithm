package main

import (
	"sort"
)

// mid 选取存在两种方式：
// 	靠近左端点：l + (r-l)>>1
//  靠近右端点：r - (r-l)>>1

// 情况一：查找确定值 target
func searchOne(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 情况二：lower_bound，查找第一个 >= target 的元素
func searchTwo(nums []int, target int) int {
	low, high, idx := 0, len(nums)-1, -1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] >= target {
			idx = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return idx
}

// 情况三：uppder_bound，查找第一个 > target 的元素
func searchThree(nums []int, target int) int {
	low, high, idx := 0, len(nums)-1, -1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			idx = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return idx
}

// 情况四：查找最后一个 <= target 的元素
func searchFour(nums []int, target int) int {
	low, high, idx := 0, len(nums)-1, -1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] <= target {
			idx = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}

// f(n) = true --> f(n+1) = true ... f(n+k) = true
func stdTest() {
	var (
		nums   []int // 升序数组
		target int   // 目标值
		n      = len(nums)
	)

	// 情况一：查找指定值 target
	idx := sort.Search(n, func(i int) bool {
		return nums[i] >= target
	})
	if idx >= n || nums[idx] != target {

	}

	// 情况二：查找 >=  target
	sort.Search(n, func(i int) bool {
		return nums[i] >= target
	})

	// 情况三：查找 > target
	sort.Search(n, func(i int) bool {
		return nums[i] > target
	})

	// 情况四：查找 <= target 的最大值
	//	1. reverse
	//  2. search
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	sort.Search(n, func(i int) bool {
		return nums[i] <= target
	})
}
