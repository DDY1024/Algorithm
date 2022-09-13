package main

// 题目链接: https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/
// 解题思路:
// https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/solution/you-xu-ju-zhen-zhong-di-kxiao-de-yuan-su-by-leetco/
//
// 二分枚举第 k 小的值，介于 [ m[0][0], m[n-1][n-1] ] 之间
// 然后 O(n) 复杂度判定二维矩阵中 <= mid 的元素个数
//
// 利用二维矩阵的特性
// 1. 同一行从左到右单调不减
// 2. 同一列从上到下单调不减

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)

	var le = func(x int) int {
		cc, row, col := 0, 0, 0
		for col < n && matrix[row][col] <= x {
			col++
		}

		for row < n {
			for col-1 >= 0 && matrix[row][col-1] > x {
				col--
			}
			cc += col
			row++
		}
		return cc
	}

	l, r, ret := matrix[0][0], matrix[n-1][n-1], -1
	for l <= r {
		mid := l + (r-l)/2
		if le(mid) >= k {
			ret = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ret
}
