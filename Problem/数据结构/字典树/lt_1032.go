package main

// https://leetcode.cn/problems/stream-of-characters/
//
// 解题思路: 一道典型的利用字典树来优化查找的题目
// 		由于搜索的为后缀，因此针对 words 数组反转后的字符串插入 trie 即可

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	child  [26]*Node
	isWord bool
}

func NewNode() *Node {
	return &Node{}
}

func getChild(ch byte) int {
	return int(ch - 'a')
}

func Insert(root *Node, s string) {
	cur := root
	for i := len(s) - 1; i >= 0; i-- { // 反向插入
		idx := getChild(byte(s[i]))
		if cur.child[idx] == nil {
			cur.child[idx] = NewNode()
		}
		cur = cur.child[idx]
	}
	cur.isWord = true
}

func Find(root *Node, s string) bool {
	cur := root
	for i := len(s) - 1; i >= 0; i-- {
		idx := getChild(byte(s[i]))
		if cur.child[idx] == nil {
			return false
		}
		cur = cur.child[idx]
		if cur.isWord {
			return true
		}
	}
	return false
}

type StreamChecker struct {
	root   *Node
	maxLen int
	data   []byte
}

func Constructor(words []string) StreamChecker {
	root := NewNode()
	maxLen := 0
	for _, s := range words {
		Insert(root, s)
		maxLen = maxInt(maxLen, len(s))
	}
	return StreamChecker{
		root:   root,
		data:   make([]byte, 0),
		maxLen: maxLen,
	}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.data = append(this.data, letter)
	if len(this.data) < this.maxLen {
		return Find(this.root, string(this.data))
	}
	return Find(this.root, string(this.data[len(this.data)-this.maxLen:]))
}
