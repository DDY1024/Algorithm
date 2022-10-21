package main

// 1. 归并排序过程 O(n+m) 可以解决此问题
// 2. O(log(m+n)) 算法参考：
//      https://leetcode.cn/problems/median-of-two-sorted-arrays/solution/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-s-114/
//    O(log(n+m)) 复杂度寻找两个有序数组中的第 K 大数\
//    总体思想是寻找一种方法，每次排除大概一半不符合条件的数字（达到二分的效果），持续不断迭代，最终找到符合要求的答案

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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
			return minInt(nums1[idx1], nums2[idx2])
		}

		half := k / 2
		newIdx1 := minInt(idx1+half, len(nums1)) - 1
		newIdx2 := minInt(idx2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIdx1], nums2[newIdx2]
		if pivot1 <= pivot2 {
			k -= newIdx1 - idx1 + 1
			idx1 = newIdx1 + 1
		} else {
			k -= newIdx2 - idx2 + 1
			idx2 = newIdx2 + 1
		}
	}

	// not arrive
}

// 寻找第 k 大的数，通过二分的方法逐步缩减求解区间
