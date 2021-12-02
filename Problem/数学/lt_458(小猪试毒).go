package main

import "math"

// 题目链接: https://leetcode-cn.com/problems/poor-pigs/
// 解题思路: https://leetcode-cn.com/problems/poor-pigs/solution/ke-lian-de-xiao-zhu-by-leetcode-solution-z0h7/
// 此题引导出一个经典的组合 DP 题目
// dp(i,j): i 只小猪经过 j 轮试毒做多可以校验多少桶水
// 边界条件: dp(0, j) = 1, dp(i, 0) = 1，没有小猪或没有校验轮数最多可以确认一桶水有毒
// dp(i,j) 状态表示 i 只小猪经过 j 轮测试最多可以确定多少桶水
// 在第 j 轮测试时可能存在 k 只小猪存活其余小猪全部死亡，存在 C(i,k) 种组合情况，枚举所有可能存活的猪的数量，可以得到以下递推式：
// f(i,j) = sum{f(k,j-1)*C(i,k)} {0 <= k <= i}  --> 为什么会是加法操作？？？
//
// f(i,1) = sum{f(k,0)*C(i,k)} = 2^i
// f(i,2) = sum{f(k,1)*C(i,k)} = 3^i
// ...
// f(i,j) = (j+1)^i  --> 因此本题目可以完全转化成纯数学计算题目
// 二项式推导: (x+1)^n = C(n,0) * x^0 + C(n,1) * x^1 + ... + C(n,n) * x^n
//
// x 个桶最多需要小猪的数量为 x - 1，因为 x-1 只小猪一轮试毒完全可以确定出哪只桶有毒
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	num := minutesToTest / minutesToDie
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(num+1))))
}

// f(i, 1) = 2^i
// f(i, 2) = 3^i
// f(i, 3) = 4^i
// 2 --> 4
// 1, 2, 3, 4
// 1, 2 | 2, 3 | 4 不需要试
// f(2, 1) = (1+1)^2 = 4
