package avl

// 参考：https://github.com/timtadh/data-structures/tree/master/tree/avl

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
	return &AvlTree{
		size: 0,
	}
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

func (self *AvlTree) Put(key int, value interface{}) {
	self.root, _ = self.root.Put(key, value)
}

func (self *AvlTree) Get(key int) *AvlNode {
	return self.root.Get(key)
}

func (self *AvlTree) Remove(key int) {
	newRoot, _ := self.root.Remove(key)
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

// LowerBound：寻找第一个 >= key 的 key
func (self *AvlTree) LowerBound(key int) *AvlNode {
	return self.root.LowerBound(key)
}

// UpperBound：寻找第一个 > key 的 key
func (self *AvlTree) UpperBound(key int) *AvlNode {
	return self.root.UpperBound(key)
}

type AvlNode struct {
	key    int
	value  interface{}
	height int
	left   *AvlNode
	right  *AvlNode
}

// 左子树 < 根节点 < 右子树
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
	if self == nil || self.left == nil {
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

	// 平衡条件：左右子树高度差的绝对值 <= 1
	for absInt(self.left.Height()-self.right.Height()) > 1 {
		// 左子树高：向右旋转
		if self.left.Height() > self.right.Height() {
			self = self.rotateRight()
		} else { // 右子树高：向左旋转
			self = self.rotateLeft()
		}
	}
	return self
}

func (self *AvlNode) Put(key int, value interface{}) (_ *AvlNode, updated bool) {
	if self == nil {
		return &AvlNode{key: key, value: value, height: 1}, false
	}

	if self.key == key {
		self.value = value
		return self, true
	}

	if key < self.key {
		self.left, updated = self.left.Put(key, value)
	} else {
		self.right, updated = self.right.Put(key, value)
	}

	if !updated {
		self.height += 1
		return self.balance(), updated
	}

	return self, updated
}

func (self *AvlNode) Remove(key int) (_ *AvlNode, value interface{}) {
	if self == nil {
		return nil, nil
	}

	if self.key == key {
		if self.left != nil && self.right != nil {
			if self.left.Size() < self.right.Size() {
				lmd := self.right.lmd()
				lmd.left = self.left
				return self.right, self.value
			} else {
				rmd := self.left.rmd()
				rmd.right = self.right
				return self.left, self.value
			}
		} else if self.left == nil {
			return self.right, self.value
		} else if self.right == nil {
			return self.left, self.value
		} else {
			return nil, self.value
		}
	}
	if key < self.key {
		self.left, value = self.left.Remove(key)
	} else {
		self.right, value = self.right.Remove(key)
	}
	return self, value
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

func (self *AvlNode) Value() interface{} {
	return self.value
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
