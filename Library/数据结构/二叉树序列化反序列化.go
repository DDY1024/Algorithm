package main

import (
	"strconv"
	"strings"
)

// 实现一种二叉树序列化/反序列化方法
// 算法能够正确执行的原因：针对 nil 节点做了序列化处理 "null"，反序列化时我们在遇到 "null" 时及时终止递归处理，这样我们其实并不需要知道左子树、右子树划分点，我们仍然可以
// 正确划分我们的左右子树
type Node struct {
	left  *Node
	right *Node
	val   int
}

func Serialization(root *Node) string {
	var (
		result strings.Builder
		encode func(root *Node)
	)

	encode = func(root *Node) {
		if root == nil {
			if result.Len() > 0 {
				result.WriteString(",")
			}
			result.WriteString("null")
			return
		}
		if result.Len() > 0 {
			result.WriteString(",")
		}
		result.WriteString(strconv.FormatInt(int64(root.val), 10))
		encode(root.left)
		encode(root.right)
	}

	encode(root)
	return result.String()
}

func Deserialization(data string) *Node {
	param, pos := strings.Split(data, ","), -1

	var decode func() *Node
	decode = func() *Node {
		pos++
		if param[pos] == "null" {
			return nil
		}

		node := new(Node)
		val, _ := strconv.ParseInt(param[pos], 10, 64)
		node.val = int(val)
		node.left = decode()
		node.right = decode()
		return node
	}

	return decode()
}
