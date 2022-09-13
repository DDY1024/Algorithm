package main

// 1. 解法一: 纯暴力搜索
func findWords(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	maxLen := 0

	dict := make(map[string]bool, 0)

	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(words); i++ {
		maxLen = maxInt(maxLen, len(words[i]))
	}

	tmp := make([]byte, maxLen)
	var generate func(x, y, c int)
	generate = func(x, y, c int) {
		dict[string(tmp[:c])] = true

		if c >= maxLen {
			return
		}

		for i := 0; i < 4; i++ {
			xx := x + dx[i]
			yy := y + dy[i]
			if xx >= 0 && xx < m && yy >= 0 && yy < n && !vis[xx][yy] {
				vis[xx][yy] = true
				tmp[c] = board[xx][yy]
				generate(xx, yy, c+1)
				vis[xx][yy] = false
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			vis[i][j] = true
			tmp[0] = board[i][j]
			generate(i, j, 1)
			vis[i][j] = false
		}
	}

	ret := make([]string, 0)
	for i := 0; i < len(words); i++ {
		if dict[words[i]] {
			ret = append(ret, words[i])
		}
	}
	return ret
}

// 2. 借助字典树辅助优化搜索（剪枝）

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

// func Constructor() Trie {
// 	return Trie{
// 		root: NewTrieNode(),
// 	}
// }

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

func findWords2(board [][]byte, words []string) []string {
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
		if cur == nil {
			return
		}
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
