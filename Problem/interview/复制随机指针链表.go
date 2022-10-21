package main

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
		nd := &Node{Val: cur.Val}
		nd.Next = cur.Next
		cur.Next = nd
		cur = cur.Next.Next
	}

	// 2. 赋值随机指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	var (
		newHead *Node
		newTail *Node
	)

	// 3. 拆分链表节点
	cur = head
	for cur != nil {
		if newHead == nil {
			newHead, newTail = cur.Next, cur.Next
		} else {
			newTail.Next = cur.Next
			newTail = newTail.Next
		}
		cur.Next = newTail.Next
		cur = cur.Next
	}

	return newHead
}
