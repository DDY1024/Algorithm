package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// 二叉树非递归前序遍历: 根节点、左子树、右子树
func preOrder(root *TreeNode) {
	if root == nil {
		return
	}

	result := make([]int, 0)
	stk := make([]*TreeNode, 0)
	stk = append(stk, root)

	for len(stk) > 0 {
		nd := stk[len(stk)-1]
		result = append(result, nd.val)

		// 先入右子树
		if nd.right != nil {
			stk = append(stk, nd.right)
		}

		// 再入左子树（出栈顺序正好相反）
		if nd.left != nil {
			stk = append(stk, nd.left)
		}
	}

	fmt.Println(result)
	return
}

// 中序非递归遍历：左子树、根节点、右子树
func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	stk := make([]*TreeNode, 0)
	result := make([]int, 0)
	cur := root
	for cur != nil || len(stk) > 0 {
		for cur != nil {
			stk = append(stk, cur)
			cur = cur.left
		}

		cur = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		result = append(result, cur.val)
		cur = cur.right
	}

	fmt.Println(result)
	return
}
