package main

// 0-1 字典树，顾名思义每个节点最多包含两个字节节点（0 或 1）
//
// 利用 0-1 字典树可以求解下面这几类题目
// 1. 求解数组中两个元素的最大异或和
//
//
// 相关题目
// 1. 最大异或和: https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/

const (
	MaxBits = 32
)

type TrieNode struct {
	child [2]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

func getBit(val, pos int) int {
	return (val >> uint(pos)) & 1
}

func Insert(root *TrieNode, val int) {
	cur := root
	for i := MaxBits - 1; i >= 0; i-- {
		bit := getBit(val, i)
		if cur.child[bit] == nil {
			cur.child[bit] = NewTrieNode()
		}
		cur = cur.child[bit]
	}
}

func Find(root *TrieNode, val int) *TrieNode {
	cur := root
	for i := MaxBits - 1; i >= 0; i-- {
		bit := getBit(val, i)
		if cur.child[bit] == nil {
			return nil
		}
		cur = cur.child[bit]
	}
	return cur
}

func MaxXor(root *TrieNode, val int) int {
	ret := 0
	cur := root
	for i := MaxBits - 1; i >= 0; i-- {
		bit := getBit(val, i)
		if cur.child[bit^1] != nil {
			cur = cur.child[bit^1]
			ret |= 1 << uint(i)
		} else {
			cur = cur.child[bit]
		}
	}
	return ret
}
