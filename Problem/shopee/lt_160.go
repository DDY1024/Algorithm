package main

// 题目链接: https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
// 题目大意
// 找出两个单向链表的交点

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1. 先计算长度，确认每个链表遍历的起始点
// 2. 遍历比较确认交点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var calcNum = func(nd *ListNode) int {
		cnt := 0
		for nd != nil {
			nd = nd.Next
			cnt++
		}
		return cnt
	}

	var minInt = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	l1, l2 := calcNum(headA), calcNum(headB)
	sameL := minInt(l1, l2)
	for l1 > sameL {
		headA = headA.Next
		l1--
	}
	for l2 > sameL {
		headB = headB.Next
		l2--
	}

	for l1 > 0 {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
		l1--
	}
	return nil
}

// 提供另外一种比较有趣的解题思路
// 1. a + c = m, b + c = m; a + c + b, b + c + a，如果相交则肯定会存在交点
func getIntersectionNodeTwo(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headA
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headB
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

// 1. 不存在交点
// a + b
// b + a
//
//
// 2. 存在交点
// a + c + b
// b + c + a
//
//
