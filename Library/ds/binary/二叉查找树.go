package main

import (
	"math"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

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

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	arr := make([]string, 0)
	var encode func(root *TreeNode)
	encode = func(root *TreeNode) {
		if root == nil {
			return
		}
		encode(root.Left)
		encode(root.Right)
		arr = append(arr, strconv.Itoa(root.Val))
	}
	encode(root)
	return strings.Join(arr, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
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

		arr = arr[:len(arr)-1] // 裁剪规则决定了需要先构建右子树再构建左子树
		// 注意：decode 先构建右子树，后构建左子树
		return &TreeNode{Val: val, Right: decode(val, high), Left: decode(low, val)}
	}

	return decode(math.MinInt32, math.MaxInt32)
}
