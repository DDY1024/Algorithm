package main

// 借助于 LIS O(nlogn) 求解方法，进行求解

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func increasingTriplet(nums []int) bool {
	minV := make([]int, 3)
	minV[0] = -(1 << 31)
	minV[1], minV[2] = 1<<31, 1<<31
	for i := 0; i < len(nums); i++ {
		idx := 2
		for minV[idx] >= nums[i] {
			idx--
		}
		if idx == 2 {
			return true
		}
		minV[idx+1] = minInt(minV[idx+1], nums[i])
	}
	return false
}
