package main

// 题目链接: https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/
//
// 解题思路:
// 		二分枚举答案，判定二维矩阵中 <= x 的元素个数
//		由于矩阵满足同行从左到右单调不减，同列从上到下单调不减，因此可以在 O(n+m) 复杂度内求解出 <= x 元素个数

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)

	var calc = func(x int) int {
		cnt, row, col := 0, 0, n-1
		for ; row < n; row++ {
			for col >= 0 && matrix[row][col] > x {
				col--
			}
			cnt += col + 1
		}
		return cnt
	}

	l, r, ret := matrix[0][0], matrix[n-1][n-1], -1
	// 注意：二分查找是寻找 >= k 的最小元素，该元素必然为矩阵中的元素
	for l <= r {
		mid := l + (r-l)/2
		if calc(mid) >= k {
			ret = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ret
}
