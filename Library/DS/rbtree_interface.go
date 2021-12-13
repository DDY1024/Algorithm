package main

// 最终参考: https://github.com/liyue201/gostl/blob/master/ds/rbtree/rbtree.go
// Red-Black tree properties:  http://en.wikipedia.org/wiki/Rbtree
//
//  1) A node is either red or black
//  2) The root is black
//  3) All leaves (NULL) are black
//  4) Both children of every red node are black
//  5) Every simple path from root to leaves contains the same number
//     of black nodes.

type Color uint

const (
	Red Color = iota
	Black
)

// k1 == k2: 0
// k1 < k2: -1
// k1 > k2: 1
type comparator func(k1, k2 interface{}) int

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	color  Color
	key    interface{}
	value  interface{}
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

func getColor(n *Node) Color {
	if n == nil {
		return Black
	}
	return n.color
}

type RbTree struct {
	root *Node
	size int
	cmp  comparator
}

func NewRBTree(cmp comparator) *RbTree {
	return &RbTree{
		root: nil,
		size: 0,
		cmp:  cmp,
	}
}

func (t *RbTree) Clear() {
	t.root = nil
	t.size = 0
	t.cmp = nil
}

func (t *RbTree) First() *Node {
	if t.root == nil {
		return nil
	}
	return minimum(t.root)
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

func (t *RbTree) Find(key interface{}) *Node {
	x := t.root
	for x != nil {
		if t.cmp(key, x.key) < 0 {
			x = x.left
		} else if t.cmp(key, x.key) == 0 {
			return x
		} else {
			x = x.right
		}
	}
	return nil
}

func (t *RbTree) Insert(key interface{}, value interface{}) {
	x := t.root
	var y *Node

	for x != nil {
		y = x
		if t.cmp(key, x.key) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &Node{parent: y, color: Red, key: key, value: value}
	t.size++

	if y == nil {
		z.color = Black
		t.root = z
		return
	} else if t.cmp(z.key, y.key) < 0 {
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

func (t *RbTree) LowerBound(key interface{}) *Node {
	return t.findLowerBoundNode(t.root, key)
}

func (t *RbTree) findLowerBoundNode(x *Node, key interface{}) *Node {
	if x == nil {
		return nil
	}

	if t.cmp(key, x.key) <= 0 {
		ret := t.findLowerBoundNode(x.left, key)
		if ret == nil {
			return x
		}
		return ret
	}

	return t.findLowerBoundNode(x.right, key)
}

func (t *RbTree) UpperBound(key interface{}) *Node {
	return t.findUpperBoundNode(t.root, key)
}

func (t *RbTree) findUpperBoundNode(x *Node, key interface{}) *Node {
	if x == nil {
		return nil
	}

	if t.cmp(x.key, key) <= 0 {
		return t.findUpperBoundNode(x.right, key)
	}

	ret := t.findUpperBoundNode(x.left, key)
	if ret != nil {
		return ret
	}
	return x
}

// 测试题目: https://leetcode-cn.com/problems/sequentially-ordinal-rank-tracker/submissions/
type SORTracker struct {
	tree *RbTree
	cur  *Node
}

type Item struct {
	name  string
	score int
}

func Constructor() SORTracker {
	tree := NewRBTree(func(v1, v2 interface{}) int {
		it1 := v1.(*Item)
		it2 := v2.(*Item)
		if it1.score < it2.score {
			return -1
		}
		if it1.score > it2.score {
			return 1
		}
		if it1.name == it2.name {
			return 0
		}
		if it1.name > it2.name {
			return -1
		}
		return 1
	})
	tree.Insert(&Item{"", 0x3f3f3f3f}, nil)
	return SORTracker{
		tree: tree,
		cur:  tree.First(),
	}
}

func (this *SORTracker) Add(name string, score int) {
	item := &Item{name, score}
	this.tree.Insert(item, nil)
	if this.tree.cmp(this.cur.key, item) < 0 {
		this.cur = this.cur.Next()
	}
}

func (this *SORTracker) Get() string {
	this.cur = this.cur.Prev()
	item := this.cur.key.(*Item)
	return item.name
}

func main() {
	Constructor()
}
