package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// 前序遍历：根节点、左子树、右子树
func preOrder(root *TreeNode) {
	result := make([]int, 0)
	stk := make([]*TreeNode, 0)
	stk = append(stk, root)

	for len(stk) > 0 {
		nd := stk[len(stk)-1]
		result = append(result, nd.val)

		// 1. 先进右
		if nd.right != nil {
			stk = append(stk, nd.right)
		}

		// 2. 后进左
		if nd.left != nil {
			stk = append(stk, nd.left)
		}
	}
}

// 中序遍历：左子树、根节点、右子树
func inOrder(root *TreeNode) {
	stk := make([]*TreeNode, 0)
	result := make([]int, 0)

	cur := root
	for cur != nil || len(stk) > 0 {
		for cur != nil {
			stk = append(stk, cur)
			cur = cur.left
		}

		// 弹出一个元素
		cur = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		result = append(result, cur.val)

		// 处理其右子树
		cur = cur.right
	}
}
