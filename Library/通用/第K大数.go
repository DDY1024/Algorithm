package main

// 题目链接：https://leetcode.cn/problems/kth-largest-element-in-an-array/?envType=study-plan-v2&envId=top-100-liked
// 第 K 大元素 等价于 排序后 len(arr)-k 索引位置的元素
func findKthLargest(nums []int, k int) int {
	return quickSelect(0, len(nums)-1, len(nums)-k, nums)
}

func quickSelect(l, r, k int, arr []int) int {
	idx := partition(l, r, arr)
	if idx == k {
		return arr[idx]
	}

	if k < idx {
		return quickSelect(l, idx-1, k, arr)
	}
	return quickSelect(idx+1, r, k, arr)
}

func partition(l, r int, arr []int) int {
	i, j, base := l, r, arr[l]
	for i < j {
		// 1. 先 j
		for i < j && arr[j] >= base { // >=
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}

		// 2. 后 i
		for i < j && arr[i] <= base { // <=
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = base
	return i
}
