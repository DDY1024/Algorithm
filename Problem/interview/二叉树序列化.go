package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	var (
		result strings.Builder
		encode func(root *TreeNode)
	)

	encode = func(root *TreeNode) {
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
		result.WriteString(strconv.FormatInt(int64(root.Val), 10))

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
