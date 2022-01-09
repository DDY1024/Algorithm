package main

// 题目链接: https://leetcode-cn.com/problems/linked-list-cycle/
// 题目大意
// 链表判环: 大步小步算法

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	// do-while 循环
	// while 循环
	p1, p2 := head, head
	for p1 != nil && p2 != nil {
		if p1.Next == nil {
			p1 = nil
		} else {
			p1 = p1.Next
		}

		if p2.Next == nil || p2.Next.Next == nil {
			p2 = nil
		} else {
			p2 = p2.Next.Next
		}

		if p1 == p2 && p1 != nil {
			return true
		}
	}
	return false
}
