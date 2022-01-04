package main

// 题目链接: https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
// 解题思路

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var find = func(cur *ListNode) (int, *ListNode, *ListNode) {
		head, tail, cnt := cur, cur, 1
		for tail.Next != nil && cnt < k {
			tail = tail.Next
			cnt++
		}
		return cnt, head, tail
	}

	var reverse = func(head *ListNode) (*ListNode, *ListNode) {
		var (
			newHead *ListNode
			newTail *ListNode
			cur     = head
			tmp     *ListNode
		)
		// 已知 k 个节点
		for i := 0; i < k; i++ {
			if newHead == nil {
				newHead, newTail, cur = cur, cur, cur.Next
				continue
			}
			tmp = cur.Next
			cur.Next = newHead
			newHead = cur
			cur = tmp
		}
		return newHead, newTail
	}

	var rHead *ListNode
	var rTail *ListNode
	cur := head
	for cur != nil {
		cnt, thead, ttail := find(cur)
		cur = ttail.Next
		if cnt >= k {
			thead, ttail = reverse(thead)
		}
		if rHead == nil {
			rHead = thead
			rTail = ttail
		} else {
			rTail.Next = thead
			rTail = ttail
		}
	}
	rTail.Next = nil // 注意别漏了，否则容易导致死循环
	return rHead
}
