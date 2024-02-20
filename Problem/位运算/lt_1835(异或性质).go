package main

// 题目链接：https://leetcode.cn/problems/find-xor-sum-of-all-pairs-bitwise-and/description/
//
// 解题思路
// 		1. 异或性质: 0^0 = 0, 0^1 = 1, 1^0=1, 1^1=0; 即奇数个1的异或结果为 1，偶数个1的异或结果为 0
// 		2. 求解 arr1[i]&arr2[j] 的异或和，考虑每一个 bit 位中的 1 在最终异或序列中出现的次数
//      3. 依次考虑每个 arr[i] 对最终结果的贡献
//      4. 由于数据范围为 10^9，因此我们只需要考虑 30 个 bit 位即可

func getXORSum(arr1 []int, arr2 []int) int {
	n, m := len(arr1), len(arr2)
	stats := make([]int, 30)

	// 1. 预处理计算 bit[i] = 1 的 arr2 中存在的元素个数
	for i := 0; i < m; i++ {
		for j := 0; j < 30; j++ {
			stats[j] += (arr2[i] >> j) & 1
		}
	}

	tot := make([]int, 30)
	// 2. 计算与 arr1[i] 与结果中每个 bit 位对最终结果的贡献
	for i := 0; i < n; i++ {
		for j := 0; j < 30; j++ {
			if arr1[i]&(1<<j) > 0 {
				tot[j] += stats[j]
			}
		}
	}

	ret := 0
	// 3. 根据 1 出现的次数，判定该位最终的结果
	for i := 0; i < 30; i++ {
		if tot[i]&1 > 0 {
			ret |= 1 << i
		}
	}
	return ret
}
