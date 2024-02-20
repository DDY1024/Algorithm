package main

import "sort"

// 题目链接：https://leetcode.cn/problems/rearranging-fruits/description/
//
// 题目大意
//		两个长度相同的数组，通过交换两个数组中的数字 nums1[i] 和 nums2[j] 多次，使得最终
// 两个数组经过排序后完全相同。每次交换的代价为 min{ nums1[i], nums2[j] }
//
// 贪心算法（https://leetcode.cn/problems/rearranging-fruits/solutions/2093855/si-wei-by-tsreaper-5oi7/）
// 		1. 首先两个数组中均存在的数全部去除掉，因为这些数本身不需要处理
//      2. 经过步骤 1 处理后，两个数组分别为 [x1, x1, ..., xk, xk] 和 [y1, y1, ..., yk, yk]，要使得两个数组最终相等
//			便是将 [x1,...,xk] 和 [y1,...,yk] 交换下位置
// 		3. 因此，最终问题变成 [x1,...,xk] 和 [y1,...,yk] 交换的最小代价
//			a. 将 x 中最小的和 y 中最大的进行匹配，一次性消除两个；这样最终代价为 x 和 y 中最小的【一半数】之和
//			b. 由 a 知道，我们每次操作时的代价为 min{ xi, yj }，同时存在另外一种策略，可能使得结果更优，即利用【全部元素的最小值】进行转移，这样交换
//				(xi, yj) 的代价为 (min, yj) + (min, xi) = 2*min
//         综上所述，在交换时，我们选择 a 和 b 策略中的最优者

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(basket1 []int, basket2 []int) int64 {
	n := len(basket1)
	minVal := int(1e9) // 全局最小值
	stats1 := make(map[int]int, n)
	stats2 := make(map[int]int, n)
	stats := make(map[int]int, n)
	for i := 0; i < n; i++ {
		minVal = minInt(minVal, basket1[i])
		minVal = minInt(minVal, basket2[i])
		stats1[basket1[i]]++
		stats2[basket2[i]]++
		stats[basket1[i]]++
		stats[basket2[i]]++
	}

	// 所有数字出现次数必然为偶数次，否则不存在
	for _, cnt := range stats {
		if cnt&1 > 0 {
			return -1
		}
	}

	// 由于 [x1, ..., xk] 和 [y1, ..., yk] 交换只需要考虑其中【最小的一半数字】，因此我们直接一维数组排序即可
	arr := make([]int, 0)
	for v := range stats1 {
		d1 := stats1[v] - stats2[v]
		for j := 0; j < d1/2; j++ {
			arr = append(arr, v)
		}
	}
	for v := range stats2 {
		d2 := stats2[v] - stats1[v]
		for j := 0; j < d2/2; j++ {
			arr = append(arr, v)
		}
	}

	sort.Ints(arr)
	ret := 0
	for i := 0; i < len(arr)/2; i++ {
		ret += minInt(arr[i], 2*minVal) // 取 a 和 b 策略的最小值
	}
	return int64(ret)
}
