package main

// https://leetcode.cn/problems/house-robber-iii/
//
//
// 二叉树上的动态规划
//
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	f := make(map[*TreeNode]int, 0) // 该节点被选中
	g := make(map[*TreeNode]int, 0) // 该节点不被选中

	var dfs func(cur *TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}

		dfs(cur.Left)
		dfs(cur.Right)

		f[cur] = cur.Val + g[cur.Left] + g[cur.Right]
		g[cur] = maxInt(f[cur.Left], g[cur.Left]) + maxInt(f[cur.Right], g[cur.Right])
	}
	dfs(root)

	return maxInt(f[root], g[root])
}
