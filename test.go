package main

import (
	"sort"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	key    int // 类型可变
	val    int // 类型可变
	height int
}

type Tree struct {
	root  *TreeNode
	count int
}

func (root *TreeNode) getHeight() int {
	if root != nil {
		return root.height
	}
	return -1
}

func (root *TreeNode) minValueNode() *TreeNode {
	nd := root
	for nd.left != nil {
		nd = nd.left
	}
	return nd
}

func (root *TreeNode) search(key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.key == key {
		return root
	}
	if root.key < key {
		return root.left.search(key)
	}
	return root.right.search(key)
}

func (root *TreeNode) getBalance() int {
	if root == nil {
		return 0
	}
	return root.left.getHeight() - root.right.getHeight()
}

func (root *TreeNode) leftRotate() *TreeNode {
	node := root.right
	root.right = node.left
	node.left = root

	root.height = maxInt(root.left.getHeight(), root.right.getHeight()) + 1
	node.height = maxInt(node.right.getHeight(), node.left.getHeight()) + 1
	return node
}

func (root *TreeNode) leftRightRotate() *TreeNode {
	root.left = root.left.leftRotate()
	root = root.rightRotate()
	return root
}

func (root *TreeNode) rightRotate() *TreeNode {
	node := root.left
	root.left = node.right
	node.right = root
	root.height = maxInt(root.left.getHeight(), root.right.getHeight()) + 1
	node.height = maxInt(node.left.getHeight(), node.right.getHeight()) + 1
	return node
}

func (root *TreeNode) rightLeftRotate() *TreeNode {
	root.right = root.right.rightRotate()
	root = root.leftRotate()
	return root
}

func (root *TreeNode) delete(key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = root.left.delete(key)
	} else if key == root.key {
		if root.left == nil && root.right == nil {
			return nil
		} else if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}

		tmp := root.minValueNode()
		root.key = tmp.key
		root.val = tmp.val
		root.right = root.right.delete(tmp.key)
		root.height = maxInt(root.left.getHeight(), root.right.getHeight()) + 1
		balance := root.getBalance()
		if balance > 1 {
			if root.left.getBalance() >= 0 {
				return root.rightRotate()
			}
			return root.leftRightRotate()
		} else if balance < -1 {
			if root.right.getBalance() <= 0 {
				return root.leftRotate()
			}
			return root.rightLeftRotate()
		}
	} else {
		root.right = root.right.delete(key)
	}

	return root
}

func (root *TreeNode) insert(key int, value int) (*TreeNode, *TreeNode) {
	if root == nil {
		root = &TreeNode{
			left:   nil,
			right:  nil,
			key:    key,
			val:    value,
			height: 0,
		}
		return root, root
	}

	if key == root.key {
		root.val = value
		return root, root
	}

	var new *TreeNode
	if key < root.key {
		root.left, new = root.left.insert(key, value)
		if root.left.getHeight()-root.right.getHeight() == 2 {
			if key < root.left.key {
				root = root.rightRotate()
			} else {
				root = root.leftRightRotate()
			}
		}
	} else {
		root.right, new = root.right.insert(key, value)
		if root.right.getHeight()-root.left.getHeight() == 2 {
			if key > root.right.key {
				root = root.leftRotate()
			} else {
				root = root.rightLeftRotate()
			}
		}
	}

	root.height = maxInt(root.left.getHeight(), root.right.getHeight()) + 1
	return root, new
}

func (root *TreeNode) exists(key int) bool {
	if root == nil {
		return false
	}

	if key == root.key {
		return true
	}

	if key < root.key {
		return root.left.exists(key)
	}
	return root.right.exists(key)
}

func (root *TreeNode) lowerBound(key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.key == key {
		return root
	}

	if key < root.key {
		return root.left.lowerBound(key)
	}

	rr := root.right.lowerBound(key)
	if rr != nil {
		return rr
	}
	return root
}

func (root *TreeNode) upperBound(key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.key == key {
		return root
	}

	if key < root.key {
		ll := root.left.upperBound(key)
		if ll != nil {
			return ll
		}
		return root
	}

	return root.right.upperBound(key)
}

func NewAVLTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(key int, value int) (new *TreeNode) {
	t.root, new = t.root.insert(key, value)
	t.count++
	return new
}

func (t *Tree) Exists(key int) bool {
	return t.root.exists(key)
}

type Query struct {
	size   int
	roomID int
	idx    int
}

type RoomInfo struct {
	size   int
	roomID int
}

func closestRoom(rooms [][]int, queries [][]int) []int {
	n, m := len(rooms), len(queries)
	roomList := make([]RoomInfo, 0, n)
	queryList := make([]Query, 0, m)
	result := make([]int, m)
	for i := 0; i < n; i++ {
		roomList = append(roomList, RoomInfo{
			size:   rooms[i][1],
			roomID: rooms[i][0],
		})
	}
	for i := 0; i < m; i++ {
		queryList = append(queryList, Query{
			size:   queries[i][1],
			roomID: queries[i][0],
			idx:    i,
		})
	}

	sort.Slice(roomList, func(i, j int) bool {
		return roomList[i].size > roomList[j].size
	})
	sort.Slice(queryList, func(i, j int) bool {
		return queryList[i].size > queryList[j].size
	})

	roomIdx, queryIdx := 0, 0
	avl := NewAVLTree()
	for queryIdx < m {
		for roomIdx < n && roomList[roomIdx].size >= queryList[queryIdx].size {
			avl.Insert(roomList[roomIdx].roomID, roomList[roomIdx].roomID)
			roomIdx++
		}
		nd1 := avl.root.lowerBound(queryList[queryIdx].roomID)
		nd2 := avl.root.upperBound(queryList[queryIdx].roomID)
		if nd1 == nil && nd2 == nil {
			result[queryList[queryIdx].idx] = -1
		} else if nd1 == nil {
			result[queryList[queryIdx].idx] = nd2.val
		} else if nd2 == nil {
			result[queryList[queryIdx].idx] = nd1.val
		} else {
			if absInt(nd1.val-queryList[queryIdx].roomID) <= absInt(nd2.val-queryList[queryIdx].roomID) {
				result[queryList[queryIdx].idx] = nd1.val
			} else {
				result[queryList[queryIdx].idx] = nd2.val
			}
		}
		queryIdx++
	}
	return result
}

func main() {

}
