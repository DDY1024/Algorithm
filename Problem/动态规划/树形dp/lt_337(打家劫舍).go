package main

// 题目链接：https://leetcode.cn/problems/house-robber-iii/
//
// 解题思路
// 		由于题目要求相邻节点同一天晚上不能被同时打劫，因此我们定义两个状态
// f(root): 表示当前子树根节点被选中的情况下，能够获得的最大收益
// g(root): 表示当前子树根节点不被选中的情况下，能够获得的最大收益

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// f[nil] 和 g[nil] 是 ok 的，默认值为 0
	f := make(map[*TreeNode]int, 0)
	g := make(map[*TreeNode]int, 0)

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
