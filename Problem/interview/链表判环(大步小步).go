package main

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

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	// 大步、小步算法，查看最终两者能否相遇
	n1, n2 := head, head
	for {
		if n1.Next == nil || n2.Next == nil || n2.Next.Next == nil {
			return false
		}
		n1 = n1.Next
		n2 = n2.Next.Next
		if n1 == n2 {
			return true
		}
	}
}

// 链表判环，并寻找环

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		if fast == slow { // 链表存在环
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p // 寻找到链表环的起点
		}
	}

	return nil
}
