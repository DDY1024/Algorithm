package main

type TrieNode struct {
	child  [26]*TrieNode
	isLeaf bool
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (this *Trie) Insert(word string) {
	n := len(word)
	cur := this.root
	for i := 0; i < n; i++ {
		cidx := int(word[i] - 'a')
		if cur.child[cidx] == nil {
			cur.child[cidx] = NewTrieNode()
		}
		cur = cur.child[cidx]
	}
	cur.isLeaf = true
}

func (this *Trie) Search(word string) bool {
	n := len(word)
	cur := this.root
	for i := 0; i < n; i++ {
		cidx := int(word[i] - 'a')
		if cur.child[cidx] == nil {
			return false
		}
		cur = cur.child[cidx]
	}
	return cur.isLeaf // 整个单词
}

func (this *Trie) StartsWith(prefix string) bool {
	n := len(prefix)
	cur := this.root
	for i := 0; i < n; i++ {
		cidx := int(prefix[i] - 'a')
		if cur.child[cidx] == nil {
			return false
		}
		cur = cur.child[cidx]
	}
	return true // 前缀
}