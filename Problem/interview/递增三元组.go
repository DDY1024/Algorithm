package main

// 类似 LIS O(nlogn) 求解思路
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 维护长度为 1 的最小值
// 维护长度为 2 的最小值
// 判断是否存在 nums[i] > 当前维护的长度为 2 的最小值
func increasingTriplet(nums []int) bool {
	minV := make([]int, 3)
	minV[0] = -(1 << 31)
	minV[1], minV[2] = 1<<31, 1<<31
	for i := 0; i < len(nums); i++ {
		idx := 2
		for idx >= 0 && minV[idx] >= nums[i] {
			idx--
		}

		if idx == 2 {
			return true
		}

		minV[idx+1] = minInt(minV[idx+1], nums[i])
	}

	return false
}
