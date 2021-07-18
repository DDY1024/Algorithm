package main

// 0-1 字典树实现，通常用于求解最大异或和，直接提供一套模板
// 类似可扩展到通常意义上的字典树
// dfs 搜索当前的状态便是一条搜索路径
//
//

const (
	BitsLimit = 18
)

type TrieNode struct {
	child [2]*TrieNode
	isR   bool
	val   int
	count int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		isR:   false,
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

func InsertNode(cur *TrieNode, val int) {
	for i := BitsLimit - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if cur.child[bit] == nil {
			cur.child[bit] = NewTrieNode()
		}
		cur = cur.child[bit]
	}
	cur.val = val
	cur.count++
}

func SearchNode(cur *TrieNode, val int) *TrieNode {
	for i := BitsLimit - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if cur.child[bit] == nil {
			return nil
		}
		cur = cur.child[bit]
	}
	return cur
}

// 0/1 字典树求解最大异或和
func FindMaxXor(cur *TrieNode, val int) int {
	for i := BitsLimit - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if cur.child[bit^1] != nil {
			cur = cur.child[bit^1]
		} else {
			cur = cur.child[bit]
		}
	}
	return val ^ cur.val
}

// 彻底删除
func DeleteNode(cur *TrieNode, val, pos int) *TrieNode {
	if cur.val >= 0 { // 正常取值
		if cur.val == val {
			cur.count--
			if cur.count == 0 {
				return nil
			}
			return cur
		}
		return cur
	}

	bit := checkBit(val, pos)
	cur.child[bit] = DeleteNode(cur.child[bit], val, pos-1)
	if !cur.isR && cur.child[0] == nil && cur.child[1] == nil {
		return nil
	}
	return cur
}
