package main

import "sort"

// 题目链接：https://leetcode.cn/problems/minimum-number-of-operations-to-make-arrays-similar/
//
// 题目大意
// 		两个长度相同的数组，对第一个数组的元素进行操作，使其变为第二个数组，求最少操作次数
//      每次操作选择 nums[i], nums[j]；使得 nums[i] += 2，nums[j] -= 2
//
// 解题报告
// 		1. https://leetcode.cn/problems/minimum-number-of-operations-to-make-arrays-similar/solution/by-endlesscheng-lusx/
//      2. https://leetcode.cn/problems/minimum-number-of-operations-to-make-arrays-similar/solution/by-tsreaper-sl0y/

//
// 1. 子问题
// 		给定两个数组 a 和 b，每次操作可以选择两个下标将 ai + 1, bj - 1；求最少操作次数
// 	 结论：将 a 和 b 分别从小到大排序，最优方案为将 ai 变成 bi
//
// 2. 原问题
//		将每次操作变为 ai + k, aj - k
//  结论：只有 mod k 相等的元素才能相互转化；将原数组 a 和 b 按照 mod k 的余数分类；每类子数组排序，并一一对应进行转化即可
//
//  由于本题中 k = 2，因此将 a 和 b 按照奇偶性分类，并进行排序，然后一一对应转化即可
//
func makeSimilar(nums []int, target []int) int64 {
	n := len(nums)
	a1 := make([]int, 0, n)
	a2 := make([]int, 0, n)
	b1 := make([]int, 0, n)
	b2 := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if nums[i]&1 > 0 {
			a1 = append(a1, nums[i])
		} else {
			a2 = append(a2, nums[i])
		}

		if target[i]&1 > 0 {
			b1 = append(b1, target[i])
		} else {
			b2 = append(b2, target[i])
		}
	}

	// Tips：为了避免多次排序，可以对原数组直接排序，然后按照奇偶性从小到大筛选出对应的数即可
	sort.Ints(a1)
	sort.Ints(a2)
	sort.Ints(b1)
	sort.Ints(b2)

	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	ret := 0
	for i := 0; i < len(a1); i++ {
		ret += abs(b1[i] - a1[i])
	}

	for i := 0; i < len(a2); i++ {
		ret += abs(b2[i] - a2[i])
	}

	return int64(ret >> 2) // ret / 4，包括 + 和 -
}
