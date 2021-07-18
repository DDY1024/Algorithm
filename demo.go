package main

import "fmt"

//
//
// 欠题
// TODO: 后缀数组
// https://leetcode-cn.com/problems/longest-common-subpath/
//

const (
	BitsLimit = 18
)

type TrieNode struct {
	child [2]*TrieNode
	isR   bool
	val   int
	count int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		isR:   false,
		val:   -0xff, // 中间节点
		count: 0,
	}
}

func checkBit(x, i int) int {
	if x&(1<<uint(i)) > 0 {
		return 1
	}
	return 0
}

func InsertNode(cur *TrieNode, val int) {
	for i := BitsLimit - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if cur.child[bit] == nil {
			cur.child[bit] = NewTrieNode()
		}
		cur = cur.child[bit]
	}
	cur.val = val
	cur.count++
}

func SearchNode(cur *TrieNode, val int) *TrieNode {
	for i := BitsLimit - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if cur.child[bit] == nil {
			return nil
		}
		cur = cur.child[bit]
	}
	return cur
}

// 0/1 字典树求解最大异或和
func FindMaxXor(cur *TrieNode, val int) int {
	for i := BitsLimit - 1; i >= 0; i-- {
		bit := checkBit(val, i)
		if cur.child[bit^1] != nil {
			cur = cur.child[bit^1]
		} else {
			cur = cur.child[bit]
		}
	}
	return val ^ cur.val
}

// 彻底删除
func DeleteNode(cur *TrieNode, val, pos int) *TrieNode {
	if cur.val >= 0 { // 正常取值
		if cur.val == val {
			cur.count--
			if cur.count == 0 {
				return nil
			}
			return cur
		}
		return cur
	}

	bit := checkBit(val, pos)
	cur.child[bit] = DeleteNode(cur.child[bit], val, pos-1)
	if !cur.isR && cur.child[0] == nil && cur.child[1] == nil {
		return nil
	}
	return cur
}

type Query struct {
	val int
	idx int
}

func maxGeneticDifference(parents []int, queries [][]int) []int {
	n := len(parents)
	adj := make([][]int, n)
	var r int
	for i := 0; i < n; i++ {
		if parents[i] != -1 {
			adj[parents[i]] = append(adj[parents[i]], i)
		} else {
			r = i
		}
	}

	qlist := make([][]Query, n)
	m := len(queries)
	for i := 0; i < m; i++ {
		u, val := queries[i][0], queries[i][1]
		qlist[u] = append(qlist[u], Query{val, i})
	}

	ans := make([]int, m)
	root := NewTrieNode()
	root.isR = true
	var dfs func(u int)
	dfs = func(u int) {
		InsertNode(root, u)
		// 查询
		for _, q := range qlist[u] {
			ans[q.idx] = FindMaxXor(root, q.val)
		}
		for _, v := range adj[u] {
			dfs(v)
		}
		DeleteNode(root, u, BitsLimit-1)
	}
	dfs(r)
	return ans
}

// dfs 搜索当前的状态是一条从根节点出发到当前节点的路径

func main() {
	fmt.Println("hello, world!")
}
