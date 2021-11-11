package main

import (
	"strconv"
	"strings"
)

// 提供一种二叉树序列化/反序列化的方式
type Node struct {
	left  *Node
	right *Node
	val   int
}

// type Tree struct {
// 	root *Node
// }

func Serialization(root *Node) string {
	var result strings.Builder
	var encode func(root *Node)
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
		node := &Node{}
		val, _ := strconv.ParseInt(param[pos], 10, 64)
		node.val = int(val)
		node.left = decode()
		node.right = decode()
		return node
	}
	return decode()
}
