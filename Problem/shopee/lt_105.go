package main

// 题目链接: https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 题目大意
// 前序、中序序列恢复二叉树

// 前序: 根、左、右
// 中序: 左、根、右
// 后序: 左、右、根

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	idx := 0
	for idx < len(inorder) {
		if inorder[idx] == val {
			break
		}
		idx++
	}

	root := &TreeNode{
		Val: val,
	}
	root.Left = buildTree(preorder[1:1+idx], inorder[:idx])
	root.Right = buildTree(preorder[1+idx:], inorder[idx+1:])
	return root
}
