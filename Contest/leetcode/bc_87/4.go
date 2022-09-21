package main

import "sort"

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimumMoney(transactions [][]int) int64 {
	n := len(transactions)
	borrow := make([][]int, 0)
	for i := 0; i < n; i++ {
		if transactions[i][0] > transactions[i][1] {
			borrow = append(borrow, append([]int(nil), transactions[i]...))
		}
	}

	sort.Slice(borrow, func(i, j int) bool {
		return borrow[i][0]-borrow[i][1]+borrow[j][0] >= borrow[j][0]-borrow[j][1]+borrow[i][0]
	})

	ans, need := 0, 0
	for i := 0; i < len(borrow); i++ {
		ans = maxInt(ans, need+borrow[i][0])
		need += borrow[i][0] - borrow[i][1]
	}

	maxB := 0
	for i := 0; i < n; i++ {
		if transactions[i][0] <= transactions[i][1] {
			maxB = maxInt(maxB, transactions[i][0])
		}
	}
	return int64(ans + maxB)
}

func minimumMoneyTwo(transactions [][]int) int64 {
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
