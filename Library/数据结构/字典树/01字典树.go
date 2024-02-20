package main

// 0-1 字典树：每个非叶子节点只包括两个节点 0 和 1
// 数组最大异或和: https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/

const (
	maxBits = 30
)

// 0 ^ 0 = 0
// 1 ^ 0 = 1
// 0 ^ 1 = 1
// 1 ^ 1 = 0

type Node struct {
	child [2]*Node
	value int
}

func child(val, pos int) int {
	return (val >> pos) & 1
}

func insertNode(root *Node, val int) {
	// 按照二进制高位 --> 低位顺序插入字典树
	for i := maxBits; i >= 0; i-- {
		idx := child(val, i)
		if root.child[idx] == nil {
			root.child[idx] = &Node{}
		}
		root = root.child[idx]
	}
	root.value = val
}

// 由于 0-1 字典树中我们将每个数已经转化成固定长度的二进制数，因此针对查找方法，能够走到最终的必然为叶子节点，直接返回 true 即可
func findNode(root *Node, val int) bool {
	for i := maxBits; i >= 0; i-- {
		idx := child(val, i)
		if root.child[idx] == nil {
			return false
		}
		root = root.child[idx]
	}
	return true
}

func maxXor(root *Node, x int) int {
	res := 0
	for i := maxBits; i >= 0; i-- {
		idx := child(x, i)
		if root.child[idx^1] != nil {
			res |= 1 << i
			root = root.child[idx^1]
		} else {
			root = root.child[idx]
		}
	}
	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaximumXOR(nums []int) int {
	var (
		root = &Node{}
		n    = len(nums)
		res  = 0
	)

	for i := 0; i < n; i++ {
		insertNode(root, nums[i])
		res = maxInt(res, maxXor(root, nums[i]))
	}
	return res
}
