package main

// 参考：http://www.cainiaoxueyuan.com/suanfa/5110.html
type Node struct {
	child  [26]*Node
	isRoot bool
	val    string
	cnt    int
}

func NewNode() *Node {
	return &Node{
		isRoot: false,
		val:    "", // 空串代表非叶子节点
		cnt:    0,
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
	cur.val = s
	cur.cnt++
}

func SearchNode(cur *Node, s string) bool {
	for _, ch := range s {
		idx := getChildIndex(byte(ch))
		if cur.child[idx] == nil {
			return false
		}
		cur = cur.child[idx]
	}
	return cur.val == s // 必须进行相等判断，防止只是作为前缀中间节点的存在
}

func isLeafNode(cur *Node) bool {
	for i := 0; i < 26; i++ {
		if cur.child[i] != nil {
			return false
		}
	}
	return true
}

// 常规字典树在删除叶子节点后，针对 trie 树做节点裁剪时需要考虑的情况更多，需要注意下
func DeleteNode(cur *Node, s string, pos int) *Node {
	if cur == nil {
		return nil
	}

	if pos >= len(s) {
		if cur.val == s {
			cur.cnt--
			if cur.cnt == 0 {
				cur.val = "" // 恢复节点初始状态
			}
			if isLeafNode(cur) && cur.cnt == 0 {
				return nil
			}
			return cur
		}
		return cur
	}

	idx := getChildIndex(s[pos])
	cur.child[idx] = DeleteNode(cur.child[idx], s, pos+1)
	if !cur.isRoot && cur.cnt == 0 && isLeafNode(cur) {
		return nil
	}
	return cur
}
