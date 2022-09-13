package main

// 题目链接：https://leetcode.cn/problems/word-search-ii/
//
// 1. 将 words 列表构建字典树
// 2. 利用字典树在回溯过程中进行剪枝优化

type TrieNode struct {
	child  [26]*TrieNode
	isLeaf bool
	word   string
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
	cur.word = word
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
	return cur.isLeaf
}

func findWords(board [][]byte, words []string) []string {
	trie := NewTrie()
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}

	m, n := len(board), len(board[0])
	dx, dy := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	occur := make(map[string]bool, len(words))

	var search func(x, y int, cur *TrieNode)
	search = func(x, y int, cur *TrieNode) {
		if cur.isLeaf {
			occur[cur.word] = true
		}

		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < m && yy >= 0 && yy < n && !vis[xx][yy] && cur.child[int(board[xx][yy]-'a')] != nil {
				vis[xx][yy] = true
				search(xx, yy, cur.child[int(board[xx][yy]-'a')])
				vis[xx][yy] = false
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur := trie.root
			if cur.child[int(board[i][j]-'a')] != nil {
				vis[i][j] = true
				search(i, j, cur.child[int(board[i][j]-'a')])
				vis[i][j] = false
			}
		}
	}

	ret := make([]string, 0)
	for i := 0; i < len(words); i++ {
		if occur[words[i]] {
			ret = append(ret, words[i])
		}
	}
	return ret
}
