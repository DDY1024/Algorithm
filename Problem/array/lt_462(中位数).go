package array

import "sort"

// 题目链接 ：https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/
//
// 给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最小操作数；在一次操作中，你可以使数组中的一个元素加 1 或者减 1
//
//
// 题目结论：数组中所有的数向中位数靠拢，最终结果是最小的；即中位数贪心

func minMoves2(nums []int) int {
	n := len(nums)

	var abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	sort.Ints(nums)
	ret, choose := 0, nums[len(nums)/2]
	for i := 0; i < n; i++ {
		ret += abs(nums[i] - choose)
	}

	return ret
}
