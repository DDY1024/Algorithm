package main

// 0-1 字典树实现，通常用于求解最大异或和，直接提供一套模板
// 类似可扩展到通常意义上的字典树
// dfs 搜索当前的状态便是一条搜索路径
//
//

const (
	BitsLimit = 20
)

type TrieNode struct {
	child [2]*TrieNode
	val   int
	count int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		val:   -0xff, // 中间节点
		count: 0,
	}
}

func checkBit(x, i int) int {
	if x&(1<<uint(i)) > 0 {
		return 1
	}
	return 0
}

func InsertNode(root *TrieNode, val int) {
	for i := BitsLimit - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if root.child[bit] == nil {
			root.child[bit] = NewTrieNode()
		}
		root = root.child[bit]
	}
	root.val = val
	root.count++
}

func SearchNode(root *TrieNode, val int) *TrieNode {
	for i := BitsLimit - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if root.child[bit] == nil {
			return nil
		}
		root = root.child[bit]
	}
	return root
}

// 0/1 字典树求解最大异或和
func FindMaxXor(root *TrieNode, val int) int {
	for i := BitsLimit - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if root.child[bit^1] != nil {
			root = root.child[bit^1]
		} else {
			root = root.child[bit]
		}
	}
	return val ^ root.val
}

// 不常用
func DeleteNode(root *TrieNode, val, pos int) *TrieNode {
	if root.val >= 0 { // 正常取值
		if root.val == val {
			return nil
		}
		return root
	}

	bit := checkBit(val, pos)
	root.child[bit] = DeleteNode(root.child[bit], val, pos+1)
	if root.child[0] == nil && root.child[1] == nil {
		return nil
	}
	return root
}
