package main

import "fmt"

/*
* 简单Radix-Tree
* 只需要简单的插入、查询、删除、子树遍历操作即可
*
* 针对单子节点进行压缩，减少不必要的空间浪费
 */

type Entry struct {
	key string
	val string
}

type Node struct {
	leaf   *Entry // 是否包含叶子节点
	prefix string
	childs map[byte]*Node
}

func NewNode(prefix string, entry *Entry) *Node {
	return &Node{
		leaf:   entry,
		prefix: prefix,
		childs: make(map[byte]*Node),
	}
}

func findCommonPrefix(s1, s2 string) int {
	idx := 0
	for ; idx < len(s1) && idx < len(s2) && s1[idx] == s2[idx]; idx++ {
	}
	return idx
}

// InsertNode 根据s插入节点entry
func InsertNode(root *Node, s string, entry *Entry) {

	var p, nd *Node
	var idx int
	nd = root
	for {

		if _, ok := nd.childs[s[0]]; !ok {
			nd.childs[s[0]] = NewNode(s, entry)
			return
		}

		p = nd
		nd = nd.childs[s[0]]
		idx = findCommonPrefix(nd.prefix, s)

		// 1. 最长公共前缀长度小于任何一个字符串，则节点需要分裂
		if idx < len(nd.prefix) && idx < len(s) {

			nd1 := NewNode(s[:idx], nil)
			p.childs[s[0]] = nd1

			nd1.childs[nd.prefix[idx]] = nd
			nd.prefix = nd.prefix[idx:]

			nd2 := NewNode(s[idx:], entry)
			nd1.childs[s[idx]] = nd2
			return
		}

		if idx >= len(s) && idx >= len(nd.prefix) {
			nd.leaf = entry
			return
		}

		if idx >= len(s) {
			nd1 := NewNode(s[:idx], entry)
			p.childs[s[0]] = nd1
			nd1.childs[nd.prefix[idx]] = nd
			nd.prefix = nd.prefix[idx:]
			return
		}

		s = s[idx:]
	}
}

// GetNode 根据s查找对应的节点
func GetNode(root *Node, s string) (bool, *Entry) {
	for {

		if _, ok := root.childs[s[0]]; !ok {
			return false, nil
		}

		root = root.childs[s[0]]
		idx := findCommonPrefix(s, root.prefix)

		if idx < len(root.prefix) {
			return false, nil
		}

		if idx == len(root.prefix) && idx == len(s) {
			return root.leaf != nil, root.leaf
		}

		s = s[idx:]
	}
}

// DeleteNode 返回该节点是否存在以及该旧节点的数据
// 思路: 先定位节点然后, 再决定是否删除该节点, 删除过程中牵扯到节点合并
// 1. 直接干掉该节点：该节点无任何孩子节点
// 2. 合并该节点和其孩子节点：该节点不是根节点且该节点只有一个孩子节点
// 3. 合并父亲节点和该节点: 父节点不为根节点，父节点不是叶子节点，父节点的孩子节点只有一个，父节点不为空
func DeleteNode(root *Node, s string) {

	var isExist bool
	var p, nd *Node
	nd = root
	for {

		if _, ok := nd.childs[s[0]]; !ok {
			isExist = false
			break
		}

		p = nd
		nd = nd.childs[s[0]]
		idx := findCommonPrefix(s, nd.prefix)

		if idx < len(nd.prefix) {
			isExist = false
			break
		}

		if idx == len(nd.prefix) && idx == len(s) {
			if nd.leaf != nil {
				isExist = true
			} else {
				isExist = false
			}
			break
		}

		s = s[idx:]
	}

	if !isExist {
		return
	}

	// 置空该节点的值
	nd.leaf = nil

	if p != nil && len(nd.childs) == 0 {
		delete(p.childs, nd.prefix[0])
	}

	if nd != root && len(nd.childs) == 1 {
		for lb := range nd.childs {
			nd.leaf = nd.childs[lb].leaf
			nd.prefix = nd.prefix + nd.childs[lb].prefix
			nd.childs = nd.childs[lb].childs
			break
		}
	}

	if p != nil && p != root && len(p.childs) == 1 && p.leaf == nil {
		for lb := range p.childs {
			p.leaf = p.childs[lb].leaf
			p.prefix = p.prefix + p.childs[lb].prefix
			p.childs = p.childs[lb].childs
			break
		}
	}
}

/*
func main() {
	root := NewNode("", nil)
	InsertNode(root, "abc", &Entry{key: "abc", val: "abc"})
	//ok1, nd1 := GetNode(root, "abc")
	//fmt.Println(ok1, nd1.key, nd1.val)
	//ok2, _ := GetNode(root, "ab")
	//fmt.Println(ok2)
	InsertNode(root, "abe", &Entry{key: "abe", val: "abe"})
	//_, ok := root.childs['a']
	//fmt.Println(ok)
	ok3, nd3 := GetNode(root, "abe")
	fmt.Println(ok3, nd3.key, nd3.val)

	InsertNode(root, "acd", &Entry{key: "acd", val: "acd"})
	ok4, nd4 := GetNode(root, "acd")
	fmt.Println(ok4, nd4.key, nd4.val)
	ok5, nd5 := GetNode(root, "abc")
	fmt.Println(ok5, nd5.key, nd5.val)
	ok6, nd6 := GetNode(root, "abe")
	fmt.Println(ok6, nd6.key, nd6.val)

	InsertNode(root, "b", &Entry{key: "b", val: "b"})
	ok7, nd7 := GetNode(root, "b")
	fmt.Println(ok7, nd7.key, nd7.val)
	//ok4, nd4 := GetNode(root, "abc")
	//fmt.Println(ok4, nd4.key, nd4.val)
	DeleteNode(root, "b")
	ok8, nd8 := GetNode(root, "b")
	fmt.Println(ok8, nd8)
}
*/

func main() {
	root := NewNode("", nil)
	InsertNode(root, "a", &Entry{key: "a", val: "a"})
	InsertNode(root, "ab", &Entry{key: "ab", val: "ab"})
	DeleteNode(root, "ab")
	ok1, nd1 := GetNode(root, "ab")
	fmt.Println(ok1, nd1)
	ok2, nd2 := GetNode(root, "a")
	fmt.Println(ok2, nd2)
	InsertNode(root, "abcd", &Entry{key: "abcd", val: "abcd"})
	InsertNode(root, "abc", &Entry{key: "abc", val: "abc"})
	DeleteNode(root, "abc")
	ok3, nd3 := GetNode(root, "abcd")
	fmt.Println(ok3, nd3)
	InsertNode(root, "abcde", &Entry{key: "abcde", val: "abcde"})
	InsertNode(root, "abcdf", &Entry{key: "abcdf", val: "abcdf"})
	DeleteNode(root, "abcd")
	ok4, nd4 := GetNode(root, "a")
	fmt.Println(ok4, nd4)
	ok5, nd5 := GetNode(root, "abcde")
	fmt.Println(ok5, nd5)
	ok6, nd6 := GetNode(root, "abcdf")
	fmt.Println(ok6, nd6)
	InsertNode(root, "abc", &Entry{key: "abc", val: "abc"})
	InsertNode(root, "abcd", &Entry{key: "abcd", val: "abcd"})
	DeleteNode(root, "abc")
	ok7, nd7 := GetNode(root, "abcde")
	fmt.Println(ok7, nd7)
	DeleteNode(root, "abcd")
	ok8, nd8 := GetNode(root, "a")
	fmt.Println(ok8, nd8)
}
