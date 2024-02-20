package main

import (
	"math"
	"strconv"
	"strings"
)

// 左子树 < 根节点 < 右子树

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
	res := make([]string, 0)

	var encode func(root *TreeNode)
	encode = func(root *TreeNode) {
		if root == nil {
			return
		}

		encode(root.Left)
		encode(root.Right)
		res = append(res, strconv.Itoa(root.Val))
	}
	encode(root)

	return strings.Join(res, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")

	var decode func(low, high int) *TreeNode
	decode = func(low, high int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}

		val, _ := strconv.Atoi(arr[len(arr)-1])
		if val < low || val > high {
			return nil
		}

		arr = arr[:len(arr)-1]
		// 注意：需要先递归构建右子树，然后构建左子树
		return &TreeNode{Val: val, Right: decode(val, high), Left: decode(low, val)}
	}

	return decode(math.MinInt32, math.MaxInt32)
}
