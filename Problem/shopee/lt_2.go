package main

// 题目链接: https://leetcode-cn.com/problems/add-two-numbers/
// 题目大意
// 1. 两链表元素相加，直接模拟操作即可

type ListNode struct {
	Val  int
	Next *ListNode
}

// 模拟加法操作有点归并排序的味道
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	var append = func(val int) {
		if head == nil {
			head = &ListNode{Val: val}
			tail = head
			return
		}
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
	}

	carry := 0
	for l1 != nil && l2 != nil {
		ret := carry + l1.Val + l2.Val
		carry = ret / 10
		ret = ret % 10
		append(ret)
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		ret := carry + l1.Val
		carry = ret / 10
		ret = ret % 10
		append(ret)
		l1 = l1.Next
	}
	for l2 != nil {
		ret := carry + l2.Val
		carry = ret / 10
		ret = ret % 10
		append(ret)
		l2 = l2.Next
	}
	// 注意: 别忘记处理 carry
	if carry != 0 { // 最多只会产生一次进位操作
		append(carry)
	}
	return head
}
