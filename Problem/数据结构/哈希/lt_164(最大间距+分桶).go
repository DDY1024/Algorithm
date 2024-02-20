package main

import (
	"math"
	"strconv"
)

// 题目链接：https://leetcode.cn/problems/maximum-gap/submissions/
//
//
// 按照一定规则分桶，相邻最大间距值只可能出现在桶之间，进而 O(n) 复杂度求解
// https://leetcode.cn/problems/maximum-gap/solution/zui-da-jian-ju-by-leetcode-solution/

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	minV, maxV := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		minV = min(minV, nums[i])
		maxV = max(maxV, nums[i])
	}
	if maxV == minV {
		return 0
	}

	// 注意: d 最小为 1, 例如 [1,1,1,1,1,5,5,5,5,5]
	d := max(1, (maxV-minV)/(n-1))
	bs := (maxV-minV)/d + 1
	maxBuckets := make([]int, bs)
	minBuckets := make([]int, bs)
	for i := 0; i < bs; i++ {
		maxBuckets[i] = -1
		minBuckets[i] = math.MaxInt
	}

	for i := 0; i < n; i++ {
		bidx := (nums[i] - minV) / d
		maxBuckets[bidx] = max(maxBuckets[bidx], nums[i])
		minBuckets[bidx] = min(minBuckets[bidx], nums[i])
	}

	i, j, ret := -1, 0, 0
	for j < bs {
		if maxBuckets[j] >= 0 {
			if i != -1 {
				ret = max(ret, minBuckets[j]-maxBuckets[i])
			}
			i = j
		}
		j++
	}
	return ret
}

func do() {
	strconv.FormatFloat()
}
