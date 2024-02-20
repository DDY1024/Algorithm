package main

// 题目链接：https://leetcode.cn/problems/path-sum-iii/
//
// 解题思路
//		利用二叉树上到根节点的路径前缀和进行求解

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	pathSum := make(map[int]int, 0)
	pathSum[0] = 1

	var dfs func(root *TreeNode, sum int) int
	dfs = func(root *TreeNode, sum int) int {
		if root == nil {
			return 0
		}

		curSum := sum + root.Val
		cnt := pathSum[curSum-targetSum] // 相应前缀和路径数

		pathSum[curSum]++
		cnt += dfs(root.Left, curSum)  // 左子树
		cnt += dfs(root.Right, curSum) // 右子树
		pathSum[curSum]--
		return cnt
	}
	return dfs(root, 0)
}
