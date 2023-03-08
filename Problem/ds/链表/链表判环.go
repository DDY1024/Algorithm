package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	// 快、慢指针，查看最后是否相遇
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

// 判断是否存在环 & 寻找环的起点
func detectCycle(head *ListNode) *ListNode {
	n1, n2 := head, head
	for {
		if n1.Next == nil || n2.Next == nil || n2.Next.Next == nil {
			return nil
		}

		n1 = n1.Next
		n2 = n2.Next.Next

		if n1 == n2 {
			n2 = head
			for n1 != n2 {
				n1 = n1.Next
				n2 = n2.Next
			}
			return n2
		}
	}
}
