package main

import (
	"strconv"
	"strings"
)

// 实现一种二叉树的序列化/反序列化方法
// 1. 针对叶子节点的孩子节点仍然进行编码 "null"
// 2. 采用前序编码方式即可（根节点、左子树、右子树）

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

// 另外一种二叉树 序列化/反序列化 的方式
// 1. 按照二叉树（前序遍历+中序遍历）/（后序遍历+中序遍历），得到二叉树的遍历序列
// 2. 根据 （前序遍历+中序遍历）/（后序遍历+中序遍历）结果，可以唯一的构造一棵二叉树
