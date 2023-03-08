package main

// 题目连接：https://leetcode.cn/problems/triples-with-bitwise-and-equal-to-zero/description/
// 		求解 nums[i]&nums[j]&nums[k]=0 的三元组的个数
//
// 解题思路
// 		1. 子集枚举优化方式

func countTriplets(nums []int) int {
	n := len(nums)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cnt[nums[i]&nums[j]]++
		}
	}

	ret := 0
	// mask := (1 << 16) - 1
	for i := 0; i < n; i++ {
		x := nums[i] ^ 0xffff
		for sub := x; sub > 0; sub = (sub - 1) & x { // 子集枚举
			ret += cnt[sub]
		}
		ret += cnt[0] // 切勿忘记 0
	}
	return ret
}
