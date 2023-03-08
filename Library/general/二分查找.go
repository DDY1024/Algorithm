package main

import (
	"fmt"
	"sort"
)

// 情况一：查找确定值 target
func searchOne(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
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

// 标准库 sort.Search 应用
//  1. 当 sort.Search 没有查找到指定目标元素时，返回数组大小 n
//     2.
func main() {
	var (
		nums   []int // 有序数组（升序）
		target int   // 目标值
	)

	// 情况一：查找指定值 target
	idx := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if idx >= len(nums) || nums[idx] != target {
		fmt.Println("No")
	}

	// 情况二：查找 >=  target
	sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})

	// 情况三：查找 > target
	sort.Search(len(nums), func(i int) bool {
		// nums[i] >= target+1
		return nums[i] > target
	})

	// 情况四：查找 <= target 的最大值
	// 1. 反转数组（降序）
	// 2. sort.Search
	sort.Search(len(nums), func(i int) bool {
		return nums[i] <= target
	})
}
