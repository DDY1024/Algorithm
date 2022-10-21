package main

import (
	"container/heap"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type Item struct {
	val  int
	node *ListNode
}

type MinPQ []*Item

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Less(i, j int) bool {
	return pq[i].val < pq[j].val
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
	var head, tail *ListNode

	pq := make(MinPQ, 0)
	for _, nd := range lists {
		if nd != nil {
			heap.Push(&pq, &Item{
				val:  nd.Val,
				node: nd,
			})
		}
	}

	for pq.Len() > 0 {
		item := pq.Top().(*Item)
		heap.Pop(&pq)

		if head == nil {
			head, tail = item.node, item.node
		} else {
			tail.Next = item.node
			tail = tail.Next
		}

		if item.node.Next != nil {
			heap.Push(&pq, &Item{
				val:  item.node.Next.Val,
				node: item.node.Next,
			})
		}
	}

	return head
}
