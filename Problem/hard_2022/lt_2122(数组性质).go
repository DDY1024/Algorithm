package main

import "sort"

// 题目链接: https://leetcode-cn.com/problems/recover-the-original-array/
// 解题思路:
// 1. K 为正整数，从最小的整数开始匹配（肯定属于第一数组），逐步寻找是否存在完整方案

func recoverArray(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	lower := make([]int, n/2)
	higher := make([]int, n/2)
	for i := n / 2; i > 0; i-- {
		if (nums[i]-nums[0])&1 > 0 {
			continue
		}
		cnt := make(map[int]int, n)
		for j := 0; j < n; j++ {
			cnt[nums[j]]++
		}
		idx, ok, K := 1, true, (nums[i]-nums[0])>>1
		lower[0] = nums[0]
		higher[0] = nums[i]
		cnt[nums[0]]--
		cnt[nums[i]]--
		for j := 0; j < n; j++ {
			if cnt[nums[j]] == 0 {
				continue
			}
			lower[idx] = nums[j]
			if cnt[nums[j]+2*K] == 0 {
				ok = false
				break
			}
			cnt[nums[j]]--
			cnt[nums[j]+2*K]--
			higher[idx] = nums[j] + 2*K
			idx++
		}
		if ok {
			break
		}
	}
	ret := make([]int, n/2)
	for i := 0; i < n/2; i++ {
		ret[i] = (lower[i] + higher[i]) / 2
	}
	return ret
}
