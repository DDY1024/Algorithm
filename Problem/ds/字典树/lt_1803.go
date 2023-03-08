package main

// 题目链接：https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/description/
// 0/1 字典树用于统计 x ^ num <= limit 中 x 的个数
//
// 利用字典树进行统计时，我们下述方法进行讨论
// num 的第 i bit 位为 a(i)
// high 的第 i bit 位为 b(i)
// x 的第 i bit 位为 c(i)
//	1. a(i)=1, b(i)=1; 如果 c(i) = 0, 则需要继续比较；如果 c(i) = 1，则肯定 < high
//  2. a(i)=1, b(i)=0; c(i) 必然为 1
//  3. a(i)=0, b(i)=1; 分情况
//  4. a(i)=0, b(i)=0; c(i) 必然为 0
//
// 最后，将区间和转化为求解前缀和

type TrieNode struct {
	cnt   int
	child [2]*TrieNode
}

func getBit(x, i int) int {
	if x&(1<<uint(i)) > 0 {
		return 1
	}
	return 0
}

func insertTrie(root *TrieNode, val int) {
	for i := 14; i >= 0; i-- { // 2^15 - 1
		bit := getBit(val, i)
		if root.child[bit] == nil {
			root.child[bit] = &TrieNode{0, [2]*TrieNode{}}
		}
		root.cnt++ // 每个节点要汇总其子树中所有叶子节点出现频次之和
		root = root.child[bit]
	}
	root.cnt++
}

func findTrie(root *TrieNode, val int) int {
	for i := 14; i >= 0; i-- { // 2^15 - 1
		bit := getBit(val, i)
		if root.child[bit] == nil {
			return 0
		}
		root = root.child[bit]
	}
	return root.cnt
}

func findLess(root *TrieNode, val, limit int) int {
	total := 0
	for i := 14; i >= 0; i-- {
		b1, b2 := getBit(val, i), getBit(limit, i)
		// 0, 1 --> 0, 1
		// 0, 0 --> 0
		// 1, 0 --> 1
		// 1, 1 --> 1, 0
		if b1 == 0 && b2 == 1 { // 可以为 0 或 1
			if root.child[0] != nil {
				total += root.child[0].cnt
			}
			if root.child[1] == nil {
				return total
			}
			root = root.child[1]
		} else if b1 == 0 && b2 == 0 {
			if root.child[0] == nil { // 必须为 0
				return total
			}
			root = root.child[0]
		} else if b1 == 1 && b2 == 0 { // 必须为 1
			if root.child[1] == nil {
				return total
			}
			root = root.child[1]
		} else { // 可以为 0 或 1
			if root.child[1] != nil {
				total += root.child[1].cnt
			}
			if root.child[0] == nil {
				return total
			}
			root = root.child[0]
		}
	}
	total += root.cnt // 相等关系
	return total
}

func countPairs(nums []int, low int, high int) int {
	n, ret := len(nums), 0
	root := &TrieNode{}
	for i := 0; i < n; i++ {
		ret += findLess(root, nums[i], high) - findLess(root, nums[i], low-1)
		insertTrie(root, nums[i])
	}
	return ret
}
