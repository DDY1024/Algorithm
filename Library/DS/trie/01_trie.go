package main

// 0-1 字典树: 每个节点的子节点只包含 0 和 1 两个子节点
// 利用 0-1 字典树求解的问题:
// 1. 最大异或和：利用字典树的前缀索引能力，总是朝着相反 bit 位搜索

const (
	MaxBits = 18
)

type TrieNode struct {
	child  [2]*TrieNode
	isRoot bool // 是否根节点，根节点不进行删除
	// isLeaf bool // 是否叶子节点
	val int // 取值
	cnt int // 出现次数
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		isRoot: false,
		val:    -0xff, // 中间节点，赋值一个负数，方便和叶子节点区分
		cnt:    0,
	}
}

func checkBit(x, i int) int {
	return (x >> uint(i)) & 1
	// if x&(1<<uint(i)) > 0 {
	// 	return 1
	// }
	// return 0
}

func InsertNode(cur *TrieNode, val int) {
	for i := MaxBits - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if cur.child[bit] == nil {
			cur.child[bit] = NewTrieNode()
		}
		cur = cur.child[bit]
	}
	cur.val = val
	cur.cnt++
}

func SearchNode(cur *TrieNode, val int) *TrieNode {
	for i := MaxBits - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if cur.child[bit] == nil {
			return nil
		}
		cur = cur.child[bit]
	}
	return cur
}

// 0/1 字典树搜索求解最大异或和
func FindMaxXor(cur *TrieNode, val int) int {
	for i := MaxBits - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if cur.child[bit^1] != nil { // 优先朝着相反 bit 方向搜索
			cur = cur.child[bit^1]
		} else {
			cur = cur.child[bit]
		}
	}
	return val ^ cur.val
}

// 0/1 字典树中由于输入字符串定长，因此对于叶子节点的判断可以直接根据指针判断
func isLeaf(cur *TrieNode) bool {
	if cur == nil {
		return false
	}
	return cur.child[0] == nil && cur.child[1] == nil
}

// 由于在 0/1 字典树场景下，我们给出的字符串一般都是定长的，因此对于每一次删除操作，肯定是可以触达叶子节点的
// 因此在删除后针对 trie 树裁剪操作是容易的
func DeleteNode(cur *TrieNode, val, pos int) *TrieNode {
	if cur == nil {
		return nil
	}

	if isLeaf(cur) {
		cur.cnt--
		if cur.cnt == 0 {
			return nil
		}
		return cur
	}

	bit := checkBit(val, pos)
	cur.child[bit] = DeleteNode(cur.child[bit], val, pos-1)
	if !cur.isRoot && isLeaf(cur) {
		return nil
	}
	return cur
}
