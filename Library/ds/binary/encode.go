package main

import (
	"strconv"
	"strings"
)

// 1. 方式一：针对 nil 节点进行编码，解码可以确保唯一性
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	var result strings.Builder
	var encode func(root *TreeNode)
	encode = func(root *TreeNode) {
		if root == nil { // nil 节点的编码也加入其中
			if result.Len() > 0 {
				result.WriteString(",")
			}
			result.WriteString("null")
			return
		}

		if result.Len() > 0 {
			result.WriteString(",")
		}
		result.WriteString(strconv.Itoa(root.Val))
		// result.WriteString(strconv.FormatInt(int64(root.Val), 10))
		encode(root.Left)
		encode(root.Right)
	}
	encode(root)
	return result.String()
}

func (this *Codec) deserialize(data string) *TreeNode {
	param, pos := strings.Split(data, ","), -1

	var decode func() *TreeNode
	decode = func() *TreeNode {
		pos++
		if param[pos] == "null" {
			return nil
		}

		node := new(TreeNode)
		val, _ := strconv.ParseInt(param[pos], 10, 64)
		node.Val = int(val)
		node.Left = decode()
		node.Right = decode()
		return node
	}

	return decode()
}

// 2. 方式二：
//     a. 中序 + 前序  --> 左子树、根节点、右子树      根节点、左子树、右子树
//     b. 中序 + 后序  --> 左子树、根节点、右子树      左子树、右子树、根节点
