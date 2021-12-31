package main

import "sort"

// 解题思路: https://leetcode-cn.com/problems/concatenated-words/

type Node struct {
	child [26]*Node
	flag  bool
}

func NewNode() *Node {
	return &Node{
		flag: false,
	}
}

func getChildIndex(ch byte) int {
	return int(ch - 'a')
}

func InsertNode(cur *Node, s string) {
	for _, ch := range s {
		idx := getChildIndex(byte(ch))
		if cur.child[idx] == nil {
			cur.child[idx] = NewNode()
		}
		cur = cur.child[idx]
	}
	cur.flag = true
}

func SearchNode(cur *Node, s string) bool {
	for _, ch := range s {
		idx := getChildIndex(byte(ch))
		if cur.child[idx] == nil {
			return false
		}
		cur = cur.child[idx]
	}
	return cur.flag
}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool { // 长度从小到大排序: 长度小的拼接成长度大的
		return len(words[i]) < len(words[j])
	})

	root := NewNode()
	var check func(nd *Node, word string) bool
	// 学会 trie 上如何进行 dfs 搜索
	check = func(nd *Node, word string) bool {
		if len(word) == 0 { // 小拼大肯定是大于 1 个的
			return true
		}

		for i, ch := range word {
			if nd.child[getChildIndex(byte(ch))] == nil {
				return false
			}
			nd = nd.child[getChildIndex(byte(ch))]
			if nd.flag && check(root, word[i+1:]) {
				return true
			}
		}
		return false
	}

	ret := make([]string, 0, len(words))
	for i := 0; i < len(words); i++ {
		if len(words[i]) == 0 {
			continue
		}
		if check(root, words[i]) { // 优化: 对于可以被其它更短字符串拼接而成的字符串，不需要加入字典树，只会增加 dfs 搜索的复杂度
			ret = append(ret, words[i])
		} else {
			InsertNode(root, words[i]) // 原子单词加入 trie
		}
	}
	return ret
}
