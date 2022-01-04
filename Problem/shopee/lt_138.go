package main

// 题目链接: https://leetcode-cn.com/problems/copy-list-with-random-pointer/
// 解题思路
// 1. 复制带随机指针的链表: 技巧性题目
// 2. 复制节点 --> 方便处理 random 指针
// 3. 拆分链表

// A --> B
// A --> A' --> B --> B'

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 1. 复制链表节点
	cur := head
	for cur != nil {
		node := &Node{Val: cur.Val}
		node.Next = cur.Next
		cur.Next = node
		cur = node.Next
	}

	// 2. 处理 random 指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 3. 拆分链表
	var newHead, newTail, tail *Node
	cur = head
	for cur != nil {
		if newHead == nil {
			newHead, newTail, tail = cur.Next, cur.Next, cur
		} else {
			newTail.Next = cur.Next
			newTail = newTail.Next
			tail.Next = cur
			tail = tail.Next
		}
		cur = cur.Next.Next
	}
	newTail.Next = nil // 清空尾节点指针
	tail.Next = nil    // 清空尾节点指针

	return newHead
}
