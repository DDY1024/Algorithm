package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// 前序遍历
func preOrder(root *TreeNode) {
	stk := make([]*TreeNode, 0)
	stk = append(stk, root)
	res := make([]int, 0)

	for len(stk) > 0 {
		nd := stk[len(stk)-1]
		res = append(res, nd.val)

		// 1. 先进右，后出右
		if nd.right != nil {
			stk = append(stk, nd.right)
		}

		// 2. 后进左，先出左
		if nd.left != nil {
			stk = append(stk, nd.left)
		}
	}
}

// 中序遍历
func inOrder(root *TreeNode) {
	res := make([]int, 0)
	stk := make([]*TreeNode, 0)
	cur := root

	for cur != nil || len(stk) > 0 {
		for cur != nil {
			stk = append(stk, cur)
			cur = cur.left
		}

		cur = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		res = append(res, cur.val)
		cur = cur.right
	}
}
