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
	child  [2]*TrieNode
	isRoot bool // root 节点标志，在递归删除时该节点是不允许被销毁的
	val    int  // 节点取值，可自定义
	cnt    int  // 出现次数
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		isRoot: false,
		val:    -0xff, // 中间节点，赋值一个负数，方便和叶子节点区分
		cnt:    0,
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
	cur.cnt++
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

func isLeaf(cur *TrieNode) bool {
	if cur == nil {
		return false
	}
	return cur.child[0] == nil && cur.child[1] == nil
}

// 由于在 0/1 字典树场景下，我们给出的字符串一般都是定长的，因此对于每一次删除操作，肯定是可以触达叶子节点的
// 因此在删除后针对 trie 树裁剪操作是容易的
func DeleteNode(cur *TrieNode, val, pos int) *TrieNode {
	if cur == nil {
		return nil
	}

	if isLeaf(cur) {
		cur.cnt--
		if cur.cnt == 0 {
			return nil
		}
		return cur
	}

	bit := checkBit(val, pos)
	cur.child[bit] = DeleteNode(cur.child[bit], val, pos-1)
	if !cur.isRoot && isLeaf(cur) {
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
	root.isRoot = true
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

func main() {
	fmt.Println("hello, world!")
}
