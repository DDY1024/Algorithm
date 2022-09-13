package string

import "sort"

// https://leetcode.cn/problems/orderly-queue/
//
// 字符串最小表示问题

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func orderlyQueue(s string, k int) string {
	bs := []byte(s)
	if k > 1 {
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		return string(bs)
	}

	// i, j, k, n := 0, 1, 0, len(bs)
	// for i < n && j < n && k < n {
	// 	if bs[(i+k)%n] == bs[(j+k)%n] {
	// 		k++
	// 	} else {
	// 		if bs[(i+k)%n] > bs[(j+k)%n] {
	// 			i++
	// 		} else {
	// 			j++
	// 		}
	// 		k = 0
	// 		if i == j { // 存在 i 追上 j 的情况
	// 			j++
	// 		}
	// 	}
	// }

	// 注意: 优化后的算法，并不保证 i <= j 一直成立，因为存在跳跃
	i, j, k, n := 0, 1, 0, len(bs)
	for i < n && j < n && k < n {
		if bs[(i+k)%n] == bs[(j+k)%n] {
			k++
		} else {
			if bs[(i+k)%n] > bs[(j+k)%n] {
				i = i + k + 1
			} else {
				j = j + k + 1
			}
			k = 0
			if i == j {
				j++
			}
		}
	}

	// 最终结果只会存在下面三种情况
	// 1. i >= n，则 j 为最小字符串表示
	// 2. j >= n, 则 i 为最小字符串表示
	// 3. i < n && j < n，则选择 i 或 j 任意一个即可，我们通常下标更小的
	idx := minInt(i, j)

	return string(bs[idx:]) + string(bs[:idx])
}
