package main

// 常规字典树介绍参考: https://segmentfault.com/a/1190000040801084

type Node struct {
	child [26]*Node
	val   string // 可以为任意类型 interface{}
	cnt   int    // 叶子节点出现次数

	// isRoot bool
	// isLeaf bool   // 叶子节点标志（针对长度不同的字符串插入情况）

}

func NewNode() *Node {
	return &Node{}
}

func getChild(ch byte) int {
	return int(ch - 'a')
}

func Insert(root *Node, s string) {
	cur := root
	for _, ch := range s {
		idx := getChild(byte(ch))
		if cur.child[idx] == nil {
			cur.child[idx] = NewNode()
		}
		cur = cur.child[idx]
	}
	cur.val = s
	cur.cnt++
}

func Exist(root *Node, s string) bool {
	cur := root
	for _, ch := range s {
		idx := getChild(byte(ch))
		if cur.child[idx] == nil {
			return false
		}
		cur = cur.child[idx]
	}
	return cur.cnt > 0
}
