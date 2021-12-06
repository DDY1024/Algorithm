package main

import "sort"

// 基于 RBTree 实现 multiset
// 测试: https://leetcode-cn.com/problems/maximum-number-of-tasks-you-can-assign/
type Color uint

const (
	Red Color = iota
	Black
)

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	color  Color
	key    int // 可拓展其它类型，此处只是为了刷题方便
	value  interface{}
	cnt    int
}

func (n *Node) Key() interface{} {
	return n.key
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) SetValue(val interface{}) {
	n.value = val
}

// 二叉搜索树中序遍历的后继节点
func (n *Node) Next() *Node {
	return successor(n)
}

// 二叉搜索树中序遍历的前驱节点
func (n *Node) Prev() *Node {
	return presuccessor(n)
}

// 寻找后继节点
func successor(x *Node) *Node {
	if x.right != nil {
		return minimum(x.right)
	}

	// 当前分支作为一棵左子树存在
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = x.parent
	}
	return y
}

// 前驱节点
func presuccessor(x *Node) *Node {
	if x.left != nil {
		return maximum(x.left)
	}

	if x.parent != nil {
		if x.parent.right == x {
			return x.parent
		}
		for x.parent != nil && x.parent.left == x {
			x = x.parent
		}
		return x.parent
	}
	return nil
}

func minimum(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func maximum(n *Node) *Node {
	for n.right != nil {
		n = n.right
	}
	return n
}

type RbTree struct {
	root *Node
	size int
}

func NewRBTree() *RbTree {
	return &RbTree{}
}

func (t *RbTree) Clear() {
	t.root = nil
	t.size = 0
}

func (t *RbTree) Begin() *Node {
	return t.First()
}

func (t *RbTree) First() *Node {
	if t.root == nil {
		return nil
	}
	return minimum(t.root)
}

func (t *RbTree) RBegin() *Node {
	return t.Last()
}

func (t *RbTree) Last() *Node {
	if t.root == nil {
		return nil
	}
	return maximum(t.root)
}

func (t *RbTree) Empty() bool {
	if t.size == 0 {
		return true
	}
	return false
}

func (t *RbTree) Size() int {
	return t.size
}

func (t *RbTree) Find(key int) *Node {
	x := t.root
	for x != nil {
		if key < x.key {
			x = x.left
		} else if key == x.key {
			return x
		} else {
			x = x.right
		}
	}
	return nil
}

func (t *RbTree) Insert(key int, value interface{}) {
	x := t.root
	var y *Node

	for x != nil {
		y = x
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &Node{parent: y, color: Red, key: key, value: value, cnt: 1}
	t.size++

	if y == nil {
		z.color = Black
		t.root = z
		return
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	t.rbInsertFixup(z)
}

func (t *RbTree) rbInsertFixup(z *Node) {
	var y *Node
	for z.parent != nil && z.parent.color == Red {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = Black
				z.parent.parent.color = Red
				t.rightRotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y != nil && y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = Black
				z.parent.parent.color = Red
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = Black
}

func (t *RbTree) Delete(node *Node) {
	z := node
	if z == nil {
		return
	}

	var x, y *Node
	if z.left != nil && z.right != nil {
		y = successor(z)
	} else {
		y = z
	}

	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	xparent := y.parent
	if x != nil {
		x.parent = xparent
	}
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != z {
		z.key = y.key
		z.value = y.value
	}

	if y.color == Black {
		t.rbDeleteFixup(x, xparent)
	}
	t.size--
}

func (t *RbTree) rbDeleteFixup(x, parent *Node) {
	var w *Node
	for x != t.root && getColor(x) == Black {
		if x != nil {
			parent = x.parent
		}
		if x == parent.left {
			x, w = t.rbFixupLeft(x, parent, w)
		} else {
			x, w = t.rbFixupRight(x, parent, w)
		}
	}
	if x != nil {
		x.color = Black
	}
}

func (t *RbTree) rbFixupLeft(x, parent, w *Node) (*Node, *Node) {
	w = parent.right
	if w.color == Red {
		w.color = Black
		parent.color = Red
		t.leftRotate(parent)
		w = parent.right
	}
	if getColor(w.left) == Black && getColor(w.right) == Black {
		w.color = Red
		x = parent
	} else {
		if getColor(w.right) == Black {
			if w.left != nil {
				w.left.color = Black
			}
			w.color = Red
			t.rightRotate(w)
			w = parent.right
		}
		w.color = parent.color
		parent.color = Black
		if w.right != nil {
			w.right.color = Black
		}
		t.leftRotate(parent)
		x = t.root
	}
	return x, w
}

func (t *RbTree) rbFixupRight(x, parent, w *Node) (*Node, *Node) {
	w = parent.left
	if w.color == Red {
		w.color = Black
		parent.color = Red
		t.rightRotate(parent)
		w = parent.left
	}
	if getColor(w.left) == Black && getColor(w.right) == Black {
		w.color = Red
		x = parent
	} else {
		if getColor(w.left) == Black {
			if w.right != nil {
				w.right.color = Black
			}
			w.color = Red
			t.leftRotate(w)
			w = parent.left
		}
		w.color = parent.color
		parent.color = Black
		if w.left != nil {
			w.left.color = Black
		}
		t.rightRotate(parent)
		x = t.root
	}
	return x, w
}

func (t *RbTree) leftRotate(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RbTree) rightRotate(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

func (t *RbTree) LowerBound(key int) *Node {
	return t.findLowerBoundNode(t.root, key)
}

func (t *RbTree) findLowerBoundNode(x *Node, key int) *Node {
	if x == nil {
		return nil
	}

	if key <= x.key {
		ret := t.findLowerBoundNode(x.left, key)
		if ret == nil {
			return x
		}
		return ret
	}

	return t.findLowerBoundNode(x.right, key)
}

func (t *RbTree) UpperBound(key int) *Node {
	return t.findUpperBoundNode(t.root, key)
}

func (t *RbTree) findUpperBoundNode(x *Node, key int) *Node {
	if x == nil {
		return nil
	}

	if x.key <= key {
		return t.findUpperBoundNode(x.right, key)
	}

	ret := t.findUpperBoundNode(x.left, key)
	if ret != nil {
		return ret
	}
	return x
}

func (t *RbTree) Traversal(do func(key int, val interface{})) {
	for node := t.First(); node != nil; node = node.Next() {
		do(node.key, node.value)
	}
}

func getColor(n *Node) Color {
	if n == nil {
		return Black
	}
	return n.color
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	n, m := len(tasks), len(workers)
	sort.Ints(tasks)
	sort.Ints(workers)

	var check = func(mid int) bool {
		have := pills
		rbt := NewRBTree()
		for i := m - mid; i < m; i++ {
			nd := rbt.Find(workers[i])
			if nd != nil {
				nd.cnt++
				continue
			}
			rbt.Insert(workers[i], workers[i])
		}

		for i := mid - 1; i >= 0; i-- {
			maxNode := rbt.Last()
			if maxNode.key >= tasks[i] {
				maxNode.cnt--
				if maxNode.cnt == 0 {
					rbt.Delete(maxNode)
				}
				continue
			}

			if have == 0 {
				return false
			}

			have--
			canNode := rbt.LowerBound(tasks[i] - strength)
			if canNode == nil {
				return false
			}

			canNode.cnt--
			if canNode.cnt == 0 {
				rbt.Delete(canNode)
			}
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
