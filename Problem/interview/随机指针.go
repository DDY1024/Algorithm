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

	cur := head
	for cur != nil {
		nd := &Node{Val: cur.Val}
		nd.Next = cur.Next
		cur.Next = nd
		cur = cur.Next.Next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	var newHead *Node
	var newTail *Node
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
