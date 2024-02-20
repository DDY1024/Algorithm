package interview

import (
	"math"
	"strconv"
	"strings"
)

//  1. 二叉搜索树序列化/反序列化
//     a. 左子树 < 根节点 < 右子树
//     b. 假设不存在重复值节点
//     c. 可直接利用【先序遍历】、【后序遍历】进行序列化、反序列化
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
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

		arr = arr[:len(arr)-1]
		// 注意：slice 裁剪，需要先构建右子树再构建左子树
		return &TreeNode{Val: val, Right: decode(val, high), Left: decode(low, val)}
	}
	return decode(math.MinInt32, math.MaxInt32)
}
