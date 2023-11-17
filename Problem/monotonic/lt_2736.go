package main

import "sort"

// 最大和查询：https://leetcode.cn/problems/maximum-sum-queries/description/?envType=daily-question&envId=2023-11-17
//
// 解题思路
//   		排序  +  单调性  +   二分查找
//
//
//   求解  nums1[i] + nums2[i] 的最大值，且满足 nums1[i] >= x && nums2[i] >= y ?
//
//	1. 将 (nums1[i], nums2[i]) 的组合按照 nums1[i] 降序
//  2. 将 (x, y) 的查询组合降序
//  3. 通过排序，我们可以确保 nums1[i] 是满足条件的，接下来只需要求解 nums2[i] 的【最优解】
//  4. 存在单调性
//			x1 + y1 <= x2 + y2 且 x1 >= x2，则 y1 <= y2，因此 (x2, y2) 的组合更优，(x1, y1) 可以丢弃
//          x1 + y1 > x2 + y2 且 y1 >= y2，则 (x1, y1) 更优  (x2, y2) 可以丢弃
//    综上所述，单调性必然是  x1 + y1 > x2 + y2 且 y1 < y2；x2 + y2 > x3 + y3 且 y2 < y3  这类形式
//    因此，给出一个查询 (x, y) 在单调栈中第一个满足 yk >= y 的组合 (xk, yk) 必然是最优的，由于单调性序列 【y单调递增】，可以利用二分查找进行加速

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	n, m := len(nums1), len(queries)
	sortN := make([][]int, n)
	for i := 0; i < n; i++ {
		sortN[i] = []int{nums1[i], nums2[i]}
	}
	sort.Slice(sortN, func(i, j int) bool {
		return sortN[i][0] > sortN[j][0]
	})

	sortQ := make([][]int, m)
	for i := 0; i < m; i++ {
		sortQ[i] = []int{queries[i][0], queries[i][1], i}
	}
	sort.Slice(sortQ, func(i, j int) bool {
		return sortQ[i][0] > sortQ[j][0]
	})

	stk := [][]int{}
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		ans[i] = -1
	}

	for i, j := 0, 0; i < m; i++ {
		x, y, idx := sortQ[i][0], sortQ[i][1], sortQ[i][2]
		for j < n && sortN[j][0] >= x {
			xx, yy := sortN[j][0], sortN[j][1]
			// 1. x + y 严格递减
			for len(stk) > 0 && stk[len(stk)-1][1] <= xx+yy {
				stk = stk[:len(stk)-1]
			}

			// 2. y 严格递增
			if len(stk) == 0 || stk[len(stk)-1][0] < yy {
				stk = append(stk, []int{yy, xx + yy})
			}

			j++
		}

		// 3. 根据 y 进行二分查找
		sidx := sort.Search(len(stk), func(i int) bool {
			return stk[i][0] >= y
		})
		if sidx < len(stk) {
			ans[idx] = stk[sidx][1]
		}
	}
	return ans
}
