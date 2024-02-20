package main

// 如何恢复 BST?
// 	https://leetcode.cn/problems/recover-binary-search-tree/description/
//
// 1. BST【中序遍历】按照从小到大的顺序排列
// 2. BST 中【两个节点】的值发生交换，则在有序排列中存在【一处】或【两处】使得 arr[i] > arr[i+1]
// 3. 我们只需要找到这【一处】或【两处】位置的相应元素进行交换即可

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var (
		nd1  *TreeNode
		nd2  *TreeNode
		prev *TreeNode
	)

	var dfs func(cur *TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}

		dfs(cur.Left)
		if prev != nil && cur.Val < prev.Val {
			if nd1 == nil {
				nd1, nd2 = prev, cur
			} else {
				nd2 = cur
			}
		}
		prev = cur
		dfs(cur.Right)
	}
	dfs(root)
	nd1.Val, nd2.Val = nd2.Val, nd1.Val
}
