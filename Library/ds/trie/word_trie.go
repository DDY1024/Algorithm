package main

// 字典树（词典）: https://segmentfault.com/a/1190000040801084

type Node struct {
	child [26]*Node
	val   string
	cnt   int
}

func getChild(ch byte) int {
	return int(ch - 'a')
}

func insert(cur *Node, s string) {
	for i := 0; i < len(s); i++ {
		idx := getChild(s[i])
		if cur.child[idx] == nil {
			cur.child[idx] = &Node{}
		}
		cur = cur.child[idx]
	}
	cur.val = s
	cur.cnt++
}

func search(cur *Node, s string) bool {
	for i := 0; i < len(s); i++ {
		idx := getChild(s[i])
		if cur.child[idx] == nil {
			return false
		}
		cur = cur.child[idx]
	}
	return cur.cnt > 0
}
