package main

import "sort"

// 题目链接: https://leetcode-cn.com/problems/maximum-number-of-tasks-you-can-assign/
// 解题思路: 算做是一道经典的二分+贪心算法的题目
//
// 1. 如果 m 个工人可以完成 k 个 task，则一定可以完成 k-1 个 task，存在单调性，可以用二分查找
// 2. 如何判定 k 个 task 是否可以被 m 个工人完成?（贪心思路）
//    a. 选择工作量最小的 k 个 task，并选择工作能力最强的 k 个工人
//    b. 如果当前工作量最大的 task，存在至少一个工人在不追加 pill 的情况下可以完成，则任何一个工人做都可以，我们不妨选择工作能力最强的那个工人
//    c. 如果不满足 b，那我们选择工作能力最小的那个工人，且该工人 + strength 是 >= 当前工作的工作量的，这样的选择是最优的，贪心很容易想到
//    d. 这个过程中需要用到 multiset 来进二分查找，由于单调性的原因，我们可以进行适当的优化（拆分成两个数组，后一个数组只需要首尾出队即可）

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	n, m := len(tasks), len(workers)
	sort.Ints(tasks)
	sort.Ints(workers)

	rarr := make([]int, minInt(n, m))
	var check = func(mid int) bool {
		warr := workers[m-mid : m]
		rstart, rend := 0, 0
		have := pills

		for i := mid - 1; i >= 0; i-- {
			if rend > rstart && rarr[rstart] >= tasks[i] {
				rstart++
				continue
			}

			if len(warr) > 0 && warr[len(warr)-1] >= tasks[i] {
				warr = warr[:len(warr)-1]
				continue
			}

			if have == 0 {
				return false
			}

			have--
			idx := sort.Search(len(warr), func(pos int) bool {
				return warr[pos] >= tasks[i]-strength
			})
			if idx < len(warr) {
				// 顺序是反的，卧槽，这个地方尤其要注意哈
				for j := len(warr) - 1; j > idx; j-- {
					rarr[rend] = warr[j]
					rend++
				}
				warr = warr[:idx]
				continue
			}

			if rstart >= rend {
				return false
			}
			rend--
		}
		return true
	}

	ans, l, r := 0, 1, minInt(n, m)
	for l <= r {
		mid := l + (r-l)/2
		if check(mid) {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
