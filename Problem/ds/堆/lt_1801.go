package main

import (
	"container/heap"
)

// https://leetcode.cn/problems/number-of-orders-in-the-backlog/description/
// buy：sell 中最低的开始
// sell: buy 中最高的开始
// 最低、最高便可以采用优先队列进行求解

type Item struct {
	price  int
	amount int
}

type MaxPQ []*Item

func (pq MaxPQ) Len() int { return len(pq) }

func (pq MaxPQ) Less(i, j int) bool {
	if pq[i].price == pq[j].price {
		return pq[i].amount > pq[j].amount
	}
	return pq[i].price > pq[j].price
}

func (pq MaxPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MaxPQ) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *MaxPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *MaxPQ) Top() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	return (*pq)[0]
}

type MinPQ []*Item

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Less(i, j int) bool {
	if pq[i].price == pq[j].price {
		return pq[i].amount > pq[j].amount
	}
	return pq[i].price < pq[j].price
}

func (pq MinPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinPQ) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *MinPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *MinPQ) Top() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	return (*pq)[0]
}

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

func getNumberOfBacklogOrders(orders [][]int) int {
	n := len(orders)
	mod := int(1e9) + 7
	sell := make(MinPQ, 0, n)
	buy := make(MaxPQ, 0, n)
	for i := 0; i < n; i++ {
		price, amount, tp := orders[i][0], orders[i][1], orders[i][2]
		if tp == 0 {
			for sell.Len() > 0 && amount > 0 {
				item := heap.Pop(&sell).(*Item)
				if item.price <= price {
					sa := minInt(amount, item.amount)
					item.amount -= sa
					amount -= sa
					if item.amount > 0 {
						heap.Push(&sell, &Item{item.price, item.amount})
					}
				} else {
					heap.Push(&sell, item)
					break
				}
			}
			if amount > 0 {
				heap.Push(&buy, &Item{price, amount})
			}
		} else {
			for buy.Len() > 0 && amount > 0 {
				item := heap.Pop(&buy).(*Item)
				if item.price >= price {
					ba := minInt(amount, item.amount)
					item.amount -= ba
					amount -= ba
					if item.amount > 0 {
						heap.Push(&buy, &Item{item.price, item.amount})
					}
				} else {
					heap.Push(&buy, item)
					break
				}
			}
			if amount > 0 {
				heap.Push(&sell, &Item{price, amount})
			}
		}
	}

	ret := 0
	for sell.Len() > 0 {
		item := heap.Pop(&sell).(*Item)
		ret += item.amount
		ret %= mod
	}
	for buy.Len() > 0 {
		item := heap.Pop(&buy).(*Item)
		ret += item.amount
		ret %= mod
	}
	return ret
}
