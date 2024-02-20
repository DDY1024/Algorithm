package main

// 题目链接：https://leetcode.cn/problems/median-of-two-sorted-arrays/description/

// 1. 如何利用二分巧妙的求解第 K 大数

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tl := len(nums1) + len(nums2)
	if tl&1 > 0 {
		return float64(findKthElement(nums1, nums2, (tl+1)/2))
	}
	return float64(findKthElement(nums1, nums2, tl/2)+findKthElement(nums1, nums2, tl/2+1)) / 2.0
}

func findKthElement(nums1, nums2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		if idx1 >= len(nums1) { // nums1 中的数已经被全部排除，直接返回 nums2 中的即可
			return nums2[idx2+k-1]
		}

		if idx2 >= len(nums2) { // nums2 中的数已经被全部删除，直接返回 nums1 中的即可
			return nums1[idx1+k-1]
		}

		// 当前迭代到寻找两个数组中最小的数，直接将头部元素做比较即可
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}

		half := k / 2
		newIdx1 := min(idx1+half, len(nums1)) - 1
		newIdx2 := min(idx2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIdx1], nums2[newIdx2]
		if pivot1 <= pivot2 {
			k -= newIdx1 - idx1 + 1
			idx1 = newIdx1 + 1
		} else {
			k -= newIdx2 - idx2 + 1
			idx2 = newIdx2 + 1
		}
	}
}
