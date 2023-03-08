package main

// 0-1 字典树
//		每个节点只包括两个孩子节点，即 0 和 1
//
// 1. 最大异或和: https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/

const (
	MaxBits = 31
)

type TrieNode struct {
	child [2]*TrieNode
	val   int
}

func getBit(val, pos int) int {
	return (val >> pos) & 1
}

func insert(cur *TrieNode, val int) {
	for i := MaxBits - 1; i >= 0; i-- {
		bit := getBit(val, i)
		if cur.child[bit] == nil {
			cur.child[bit] = &TrieNode{}
		}
		cur = cur.child[bit]
	}
	cur.val = val
}

func search(cur *TrieNode, val int) bool {
	for i := MaxBits - 1; i >= 0; i-- {
		bit := getBit(val, i)
		if cur.child[bit] == nil {
			return false
		}
		cur = cur.child[bit]
	}
	return true
}

func maxXor(cur *TrieNode, val int) int {
	ret := 0
	for i := MaxBits - 1; i >= 0; i-- {
		bit := getBit(val, i)
		if cur.child[bit^1] != nil {
			cur = cur.child[bit^1]
			ret |= 1 << i
		} else {
			cur = cur.child[bit]
		}
	}
	return ret
}
