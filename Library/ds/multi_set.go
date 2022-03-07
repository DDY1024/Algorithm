package main

import "sort"

// 基于 avl 树 node 加 count 属性实现
// 测试题目: https://leetcode-cn.com/problems/maximum-number-of-tasks-you-can-assign/

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type AvlTree struct {
	root *AvlNode
	size int
}

func NewAvlTree() *AvlTree {
	return &AvlTree{}
}

func (self *AvlTree) Root() *AvlNode {
	return self.root
}

func (self *AvlTree) Size() int {
	return self.size
}

func (self *AvlTree) Exists(key int) bool {
	return self.root.Exists(key)
}

func (self *AvlTree) Put(key int) {
	self.root, _ = self.root.Put(key)
}

func (self *AvlTree) Get(key int) *AvlNode {
	return self.root.Get(key)
}

func (self *AvlTree) Remove(key int) {
	newRoot := self.root.Remove(key)
	self.root = newRoot
}

func (self *AvlTree) MaxNode() *AvlNode {
	if self.root == nil {
		return nil
	}

	cur := self.root
	for cur.right != nil {
		cur = cur.right
	}
	return cur
}

func (self *AvlTree) MinNode() *AvlNode {
	if self.root == nil {
		return nil
	}

	cur := self.root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func (self *AvlTree) LowerBound(key int) *AvlNode {
	return self.root.LowerBound(key)
}

func (self *AvlTree) UpperBound(key int) *AvlNode {
	return self.root.UpperBound(key)
}

type AvlNode struct {
	key    int // types.Hashable
	height int
	left   *AvlNode
	right  *AvlNode
	count  int
}

func (self *AvlNode) Exists(key int) bool {
	if self == nil {
		return false
	}

	if self.key == key {
		return true
	}

	if key < self.key {
		return self.left.Exists(key)
	}

	return self.right.Exists(key)
}

func (self *AvlNode) Get(key int) *AvlNode {
	if self == nil {
		return nil
	}

	if self.key == key {
		return self
	}

	if key < self.key {
		return self.left.Get(key)
	}

	return self.right.Get(key)
}

func (self *AvlNode) popNode(node *AvlNode) *AvlNode {
	if self == nil {
		return nil
	}

	if self == node {
		var n *AvlNode
		if node.left != nil {
			n = node.left
		} else if node.right != nil {
			n = node.right
		} else {
			n = nil
		}
		node.left = nil
		node.right = nil
		return n
	}

	if node.key < self.key {
		self.left = self.left.popNode(node)
	} else {
		self.right = self.right.popNode(node)
	}

	self.height = maxInt(self.left.Height(), self.right.Height()) + 1
	return self
}

func (self *AvlNode) pushNode(node *AvlNode) *AvlNode {
	if self == nil {
		node.height = 1
		return node
	}
	if node.key < self.key {
		self.left = self.left.pushNode(node)
	} else {
		self.right = self.right.pushNode(node)
	}

	self.height = maxInt(self.left.Height(), self.right.Height()) + 1
	return self
}

func (self *AvlNode) rotateRight() *AvlNode {
	if self == nil {
		return self
	}

	if self.left == nil {
		return self
	}

	newRoot := self.left.rmd()
	self = self.popNode(newRoot)
	newRoot.left = self.left
	newRoot.right = self.right
	self.left = nil
	self.right = nil
	return newRoot.pushNode(self)
}

func (self *AvlNode) rotateLeft() *AvlNode {
	if self == nil {
		return self
	}

	if self.right == nil {
		return self
	}

	newRoot := self.right.lmd()
	self = self.popNode(newRoot)
	newRoot.left = self.left
	newRoot.right = self.right
	self.left = nil
	self.right = nil
	return newRoot.pushNode(self)
}

func (self *AvlNode) balance() *AvlNode {
	if self == nil {
		return self
	}

	for absInt(self.left.Height()-self.right.Height()) > 2 {
		if self.left.Height() > self.right.Height() {
			self = self.rotateRight()
		} else {
			self = self.rotateLeft()
		}
	}

	return self
}

func (self *AvlNode) Put(key int) (*AvlNode, bool) {
	if self == nil {
		return &AvlNode{key: key, height: 1, count: 1}, false
	}

	if self.key == key {
		self.count++
		return self, true
	}

	var updated bool
	if key < self.key {
		self.left, updated = self.left.Put(key)
	} else {
		self.right, updated = self.right.Put(key)
	}

	if !updated { // 新增节点，节点高度 + 1
		self.height++
		return self.balance(), updated
	}

	return self, updated
}

func (self *AvlNode) Remove(key int) *AvlNode {
	if self == nil {
		return nil
	}

	if self.key == key {
		// multi
		self.count--
		if self.count > 0 {
			return self
		}

		if self.left != nil && self.right != nil {
			if self.left.Size() < self.right.Size() {
				lmd := self.right.lmd()
				lmd.left = self.left
				return self.right
			} else {
				rmd := self.left.rmd()
				rmd.right = self.right
				return self.left
			}
		} else if self.left == nil {
			return self.right
		} else if self.right == nil {
			return self.left
		} else {
			return nil
		}
	}
	if key < self.key {
		self.left = self.left.Remove(key)
	} else {
		self.right = self.right.Remove(key)
	}
	return self
}

func (self *AvlNode) Height() int {
	if self == nil {
		return 0
	}
	return self.height
}

func (self *AvlNode) Size() int {
	if self == nil {
		return 0
	}
	return 1 + self.left.Size() + self.right.Size()
}

func (self *AvlNode) Key() int {
	return self.key
}

func (self *AvlNode) Left() *AvlNode {
	if self.left == nil {
		return nil
	}
	return self.left
}

func (self *AvlNode) Right() *AvlNode {
	if self.right == nil {
		return nil
	}
	return self.right
}

func (self *AvlNode) _md(side func(*AvlNode) *AvlNode) *AvlNode {
	if self == nil {
		return nil
	} else if side(self) != nil {
		return side(self)._md(side)
	} else {
		return self
	}
}

func (self *AvlNode) lmd() *AvlNode {
	return self._md(func(node *AvlNode) *AvlNode { return node.left })
}

func (self *AvlNode) rmd() *AvlNode {
	return self._md(func(node *AvlNode) *AvlNode { return node.right })
}

func (self *AvlNode) LowerBound(key int) *AvlNode {
	if self == nil {
		return nil
	}

	if self.key == key {
		return self
	}

	if key < self.key {
		ll := self.left.LowerBound(key)
		if ll != nil {
			return ll
		}
		return self
	}

	return self.right.LowerBound(key)
}

func (self *AvlNode) UpperBound(key int) *AvlNode {
	if self == nil {
		return nil
	}

	if self.key <= key {
		return self.right.UpperBound(key)
	}

	ll := self.left.UpperBound(key)
	if ll != nil {
		return ll
	}

	return self
}

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	n, m := len(tasks), len(workers)
	sort.Ints(tasks)
	sort.Ints(workers)

	var check = func(mid int) bool {
		have := pills
		avl := NewAvlTree()
		for i := m - mid; i < m; i++ {
			avl.Put(workers[i])
		}

		for i := mid - 1; i >= 0; i-- {
			maxNode := avl.MaxNode()
			if maxNode.key >= tasks[i] {
				avl.Remove(maxNode.key)
				continue
			}

			if have == 0 {
				return false
			}

			have--
			canNode := avl.LowerBound(tasks[i] - strength)
			if canNode == nil {
				return false
			}
			avl.Remove(canNode.key)
		}
		return true
	}

	ans, l, r := 0, 1, minInt(n, m)
	for l <= r {
		mid := l + (r-l)/2
		if check(mid) {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
