package main

// https://leetcode.cn/problems/path-sum-iii/
//
// 利用二叉树上根节点到当前节点的路径前缀和，实现 O(n) 复杂度求解区间路径和方案数
//
// 树上前缀和 vs 数组前缀和
// 		当前节点到根节点路径上的所有祖先节点才能作为前缀

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	prefixSum := make(map[int]int, 0)
	prefixSum[0] = 1

	var dfs func(cur *TreeNode, sum int) int
	dfs = func(cur *TreeNode, sum int) int {
		if cur == nil {
			return 0
		}

		curSum := sum + cur.Val // cur.Val not root.Val
		ret := prefixSum[curSum-targetSum]

		prefixSum[curSum]++
		ret += dfs(cur.Left, curSum)  // 左子树
		ret += dfs(cur.Right, curSum) // 右子树
		prefixSum[curSum]--           // 只会包含当前节点到根节点路径上的所有祖先节点

		return ret
	}
	return dfs(root, 0)
}
