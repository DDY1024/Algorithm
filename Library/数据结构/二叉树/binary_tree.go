package main

import (
	"strconv"
	"strings"
)

// 针对 nil 节点同样进行序列化，采用先序遍历方式
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
	res := make([]string, 0)

	var encode func(root *TreeNode)
	encode = func(root *TreeNode) {
		if root == nil {
			res = append(res, "null")
			return
		}

		res = append(res, strconv.Itoa(root.Val))
		encode(root.Left)
		encode(root.Right)
	}
	encode(root)

	return strings.Join(res, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	vals, pos := strings.Split(data, ","), -1

	var decode func() *TreeNode
	decode = func() *TreeNode {
		pos++
		if vals[pos] == "null" {
			return nil
		}

		nd := &TreeNode{}
		val, _ := strconv.ParseInt(vals[pos], 10, 64)
		nd.Val = int(val)
		nd.Left = decode()
		nd.Right = decode()
		return nd
	}

	return decode()
}
