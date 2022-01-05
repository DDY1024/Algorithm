package main

// 题目链接: https://leetcode-cn.com/problems/binary-tree-right-side-view/
// 解题思路
// 二叉树右视图：即每一层最右边的那个节点
// 1. 层次遍历
// 2. 根节点、右子树、左子树的处理顺序
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := make([]int, 0)
	var dfs func(r *TreeNode, dep int)
	dfs = func(r *TreeNode, dep int) {
		if r == nil {
			return
		}

		if len(ret)-1 < dep { // 最多差一
			ret = append(ret, r.Val)
		}
		dfs(r.Right, dep+1)
		dfs(r.Left, dep+1)
	}

	dfs(root, 0)
	return ret
}
