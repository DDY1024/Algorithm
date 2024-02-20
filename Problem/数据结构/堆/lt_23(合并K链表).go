package main

import "container/heap"

// 题目链接：https://leetcode.cn/problems/merge-k-sorted-lists/description/?envType=study-plan-v2&envId=top-100-liked
//
// 利用堆进行合并处理

type ListNode struct {
	Val  int
	Next *ListNode
}

type Item struct {
	ln *ListNode
}

type MinPQ []*Item

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Less(i, j int) bool {
	return pq[i].ln.Val < pq[j].ln.Val
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	pq := make(MinPQ, 0)
	head := &ListNode{}
	tail := head

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(&pq, &Item{
				lists[i],
			})
		}
	}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		tail.Next = item.ln
		tail = tail.Next
		if item.ln.Next != nil {
			heap.Push(&pq, &Item{
				item.ln.Next,
			})
		}
	}

	tail.Next = nil
	return head.Next
}
