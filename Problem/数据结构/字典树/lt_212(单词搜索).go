package main

// 题目链接（中等）：https://leetcode.cn/problems/word-search/
// 题目链接（困难）：https://leetcode.cn/problems/word-search-ii/

// 1. 利用 Trie 存储整个单词词典
// 2. 在 dfs 搜索时，借助 Trie 进行判断，剪枝优化

type TrieNode struct {
	child [26]*TrieNode
	flag  bool
	word  string
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

func insert(root *TrieNode, word string) {
	n := len(word)
	for i := 0; i < n; i++ {
		cidx := int(word[i] - 'a')
		if root.child[cidx] == nil {
			root.child[cidx] = NewTrieNode()
		}
		root = root.child[cidx]
	}
	root.flag = true
	root.word = word
}

func exists(root *TrieNode, word string) bool {
	n := len(word)
	for i := 0; i < n; i++ {
		cidx := int(word[i] - 'a')
		if root.child[cidx] == nil {
			return false
		}
		root = root.child[cidx]
	}
	return root.flag
}

func findWords(board [][]byte, words []string) []string {
	root := NewTrieNode()
	for i := 0; i < len(words); i++ {
		insert(root, words[i])
	}

	m, n := len(board), len(board[0])
	occur := make(map[string]bool, len(words))
	dx, dy := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	var dfs func(x, y int, cur *TrieNode)
	dfs = func(x, y int, cur *TrieNode) {
		if cur.flag {
			occur[cur.word] = true
		}

		for i := 0; i < 4; i++ {
			xx, yy := x+dx[i], y+dy[i]
			if xx >= 0 && xx < m && yy >= 0 && yy < n && !vis[xx][yy] && cur.child[int(board[xx][yy]-'a')] != nil {
				vis[xx][yy] = true
				dfs(xx, yy, cur.child[int(board[xx][yy]-'a')])
				vis[xx][yy] = false
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur := root
			if cur.child[int(board[i][j]-'a')] != nil {
				vis[i][j] = true
				dfs(i, j, cur.child[int(board[i][j]-'a')])
				vis[i][j] = false
			}
		}
	}

	// 存在重复单词，需要去重
	ret := make([]string, 0)
	for word := range occur {
		ret = append(ret, word)
	}
	return ret
}
