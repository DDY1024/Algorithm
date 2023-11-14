package main

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	n, m := len(spells), len(potions)
	sort.Ints(potions)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		idx := sort.Search(m, func(j int) bool {
			return spells[i]*potions[j] >= int(success) // 利用单调递增特性来做
		})
		ans[i] = m - idx
	}
	return ans
}

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

func recoverTree(root *TreeNode) {
	var (
		nd1  *TreeNode
		nd2  *TreeNode
		prev *TreeNode
	)

	var dfs func(cur *TreeNode)
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}

		dfs(cur.Left)
		if prev != nil && cur.Val < prev.Val {
			if nd1 == nil {
				nd1, nd2 = prev, cur
			} else {
				nd2 = cur
			}
		}
		prev = cur
		dfs(cur.Right)
	}
	dfs(root)
	nd1.Val, nd2.Val = nd2.Val, nd1.Val
}

type NumArray struct {
	stats []int
	nums  []int
	n     int
}

func lowBit(x int) int {
	return x & (-x)
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	stats := make([]int, n+1)
	for i := 1; i <= n; i++ {
		stats[i] += nums[i-1]
		j := i + lowBit(i)
		if j <= n {
			stats[j] += stats[i]
		}
	}
	return NumArray{
		n:     n,
		stats: stats,
		nums:  nums,
	}
}

func (this *NumArray) Update(index int, val int) {
	delta := val - this.nums[index]
	this.nums[index] = val
	for i := index + 1; i <= this.n; i += lowBit(i) {
		this.stats[i] += delta
	}
}

func (this *NumArray) Sum(pos int) int {
	ret := 0
	for i := pos + 1; i > 0; i -= lowBit(i) {
		ret += this.stats[i]
	}
	return ret
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Sum(right) - this.Sum(left-1)
}

func countPoints(rings string) int {
	n := len(rings)

	var get = func(b byte) int {
		switch b {
		case 'R':
			return 1
		case 'G':
			return 2
		case 'B':
			return 4
		default:
			return 1
		}
	}

	// "B0B6G0R6R0R6G9"
	count := make([]int, 10)
	for i := 0; i < n; i += 2 {
		count[int(rings[i^1]-'0')] |= get(rings[i])
	}

	ret := 0
	for i := 0; i < 10; i++ {
		if count[i] == 7 {
			ret++
		}
	}
	return ret
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := []*Node{root}
	for len(q) > 0 {
		q1 := make([]*Node, 0)
		for i := 0; i < len(q); i++ {
			if i+1 < len(q) {
				q[i].Next = q[i+1]
			}
			if q[i].Left != nil {
				q1 = append(q1, q[i].Left)
			}
			if q[i].Right != nil {
				q1 = append(q1, q[i].Right)
			}
		}
		q = q1
	}
	return root
}

func findMaximumXOR(nums []int) int {
	var getBit = func(x, p int) int {
		return (x >> p) & 1
	}

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n, res := len(nums), 0
	for i := 0; i < n; i++ {
		arr1, tmp := nums[:i+1], 0
		for j := 30; j >= 0; j-- {
			arr2 := make([]int, 0)
			for k := 0; k < len(arr1); k++ {
				if getBit(arr1[k], j)^getBit(nums[i], j) == 1 {
					tmp |= 1 << j
					arr2 = append(arr2, arr1[k])
				}
			}
			if len(arr2) > 0 {
				arr1 = arr2
			}
		}
		res = maxInt(res, tmp)
	}
	return res
}

// 0-1 字典树求解的题目
// 0 ~ 30 			(1<<31) - 1
//  (1<<31) - 1   bit_0、bit_1、...、bit_30
//  max_xor  0^0 = 0, 0^1 = 1, 1^0 = 1, 1^1 = 0
