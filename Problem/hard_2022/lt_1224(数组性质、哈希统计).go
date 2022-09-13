package hard2022

// 题目链接: https://leetcode.cn/problems/maximum-equal-frequency/

// 解题思路
// 1. 我们最多只需要关注两个频率 maxFreq 和 maxFreq - 1 即可
// 		需要确保这两个频率出现的元素构成全部的元素 maxFreq * freq[maxFreq] + (maxFreq-1) * freq[maxFreq-1] == i+1
// 2. 另外，我们需要关注一个特殊频率 freq[1]（每次操作只会删除一个元素）

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxEqualFreq(nums []int) int {
	n := len(nums)
	stats := make(map[int]int, n)
	freq := make(map[int]int, n)
	maxLen, maxFreq := 0, 0

	for i := 0; i < n; i++ {
		stats[nums[i]]++
		maxFreq = maxInt(maxFreq, stats[nums[i]]) // 最多允许存在两种不同频率，我不妨维护一个 maxFreq，剩余一个肯定为 maxFreq - 1
		freq[stats[nums[i]]]++
		freq[stats[nums[i]]-1]--
		if freq[1] == i+1 || // 全部只出现一次
			(freq[1] == 1 && maxFreq*freq[maxFreq]+1 == i+1) || // 一个数出现一次，剩余数出现次数相同
			(maxFreq*freq[maxFreq]+(maxFreq-1)*freq[maxFreq-1] == i+1 && freq[maxFreq] == 1) { // maxFreq - 1，maxFreq 两种频次，且最大频次只有一个数
			maxLen = i + 1
		}
	}
	return maxLen
}
