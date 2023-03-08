package main

// 题目链接：https://leetcode.cn/problems/binary-tree-coloring-game/description/
// 解题思路：https://leetcode.cn/problems/binary-tree-coloring-game/solutions/2089813/mei-you-si-lu-yi-zhang-tu-miao-dong-pyth-btav/

// 个人思考
// 		1. 首先，对于 x 和 y 可以随便选，只需要保证 x != y 即可
//  	2. 后续操作只能基于 x 和 y 扩展选择邻接点（重要）
//      3. 在选择 x 点后，整棵树被划分成三部分：左子树，右子树，x 父节点为根的子树（树结构划分）
//      4. y 的最优选择为这三棵子树中的最大子树的根节点
//			分析可知，如果 y 不选择子树根节点，只会导致被 x 占据更多的节点，因为基于性质（2）, 另外两棵子树 y 是无法染色的
//			因此，选择子树根节点更优；进一步我们应该选择最大子树的根节点
//
//
// 扩展问题（二叉树 or 一般树）
// 		1. 如果起初 x 点可以随便选择，按照上述游戏规则，应该如何选择 x 点，保证第一位玩家染色节点数最多？
//			结论：x 节点应该选择【树的重心】
// 		2. 树的重心
// 			对于一棵 n 个节点的无根树，找到一个点，将无根树变为以该节点为根的有根树时，【最大子树的结点数最小】

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

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	lnum, rnum := 0, 0
	var dfs func(nd *TreeNode) int
	dfs = func(nd *TreeNode) int {
		if nd == nil {
			return 0
		}
		n1 := dfs(nd.Left)
		n2 := dfs(nd.Right)
		if nd.Val == x {
			lnum, rnum = n1, n2
		}
		return n1 + n2 + 1
	}
	dfs(root)

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	return 2*maxInt(n-lnum-rnum-1, maxInt(lnum, rnum)) > n
}
