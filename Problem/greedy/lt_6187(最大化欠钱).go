package main

// 题目链接：https://leetcode.cn/problems/minimum-money-required-before-transactions/
// 1. 任意一种操作顺序，求解 "最大值"
// 2. 至少一种操作顺序，求解 "最小值"

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 贪心思路
// 1. 先欠钱，再挣钱
// 2. 欠钱过程中，最大值必然出现在所有负收益操作的最后一次操作，即为 max{ borrow - (cost[i] - back[i]) + cost[i] }，其中 borrow 为所有亏损的金额总和
// 3. 挣钱过程中，我们需要使得第一次挣钱操作的 cost[i] 值最大，即 max{ borrow + cost[i] }，因为后续的挣钱操作只会更小
//
// 综上所述，我们在一次遍历的过程中维护这些变量，最终求解即可

func minimumMoney(transactions [][]int) int64 {
	n := len(transactions)
	// 借钱的时候，最坏情况肯定是出现在最后一比交易中
	borrow, maxBack, maxCost := 0, 0, 0
	for i := 0; i < n; i++ {
		if transactions[i][0] > transactions[i][1] {
			borrow += transactions[i][0] - transactions[i][1]
			maxBack = maxInt(maxBack, transactions[i][1])
			continue
		}
		maxCost = maxInt(maxCost, transactions[i][0])
	}
	return int64(maxInt(borrow+maxBack, borrow+maxCost))
}
