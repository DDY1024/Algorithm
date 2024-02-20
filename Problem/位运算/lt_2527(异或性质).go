package main

// 题目链接：https://leetcode.cn/problems/find-xor-beauty-of-array/description/
//
// 解题思路
//		1. 异或性质：x^x = 0
//      2. (i, j, k) ^ (j, i, k) = ((nums[i]|nums[j])&nums[k])^((nums[j]|nums[i])&nums[k]) = 0，当 i, j, k 互不相同；
//		3. (i, j, k) ^ (j, i, k) = (nums[i]&nums[k])^(nums[j]&nums[k]) = 0，当 j = i 且 i != k
//      4. (i, j, k) ^ (j, i, k) = nums[k]^nums[k] = 0，当 i = j = k
// 		5. 2 和 3 情况均存在偶数种，因此异或结果为 0; 4 情况仅仅存在1种；因此最终结果便是数组元素的异或和

func xorBeauty(nums []int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}
