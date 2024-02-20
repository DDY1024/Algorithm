package main

// https://leetcode.cn/problems/sum-of-prefix-scores-of-strings/

// 解题思路：简单的利用字典树进行统计的题目

type TrieNode struct {
	child [26]*TrieNode
	cnt   int
}

func Insert(root *TrieNode, word string) {
	for _, c := range word {
		if root.child[int(c-'a')] == nil {
			root.child[int(c-'a')] = &TrieNode{}
		}
		root.child[int(c-'a')].cnt++
		root = root.child[int(c-'a')]
	}
}

func Calc(root *TrieNode, word string) int {
	ret := 0
	for _, c := range word {
		ret += root.child[int(c-'a')].cnt
		root = root.child[int(c-'a')]
	}
	return ret
}

func sumPrefixScores(words []string) []int {
	n := len(words)
	root := &TrieNode{}
	for i := 0; i < n; i++ {
		Insert(root, words[i])
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = Calc(root, words[i])
	}
	return ans
}
