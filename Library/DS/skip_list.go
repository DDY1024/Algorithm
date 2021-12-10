package Library

import (
	"math/rand"
	"time"
)

var (
	maxLevel = 20
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// level取值范围 0 ~ maxLevel
// redis 具体实现中 maxLevel = 32, p = 1/4
func randomLevel() int {
	level := 0
	for rand.Intn(100) < 50 {
		level++
	}
	return minInt(level, maxLevel)
}

func NewNode(key, val, n int) *Node {
	return &Node{
		key:  key,
		val:  val,
		next: make([]*Node, n),
	}
}

type Node struct {
	key  int
	val  int
	next []*Node
}

type Skiplist struct {
	level int
	size  int
	head  *Node
}

func NewSkiplist() Skiplist {
	return Skiplist{
		level: 0,
		size:  0,
		head:  NewNode(-0x3f3f3f3f, -0x3f3f3f3f, maxLevel+1), // 初始赋值一个不可能存在的值
	}
}

func (sl *Skiplist) IsEmpty() bool {
	return sl.size == 0
}

func (sl *Skiplist) Search(target int) bool {
	if sl.IsEmpty() {
		return false
	}
	cur, pre := sl.head, sl.head
	for level := sl.level; level >= 0; level-- {
		for cur = cur.next[level]; cur != nil && cur.key < target; pre, cur = cur, cur.next[level] {
		}
		if cur != nil && cur.key == target {
			return true
		}
		cur = pre
	}
	return false
}

func (sl *Skiplist) SearchLessEqual(target int) *Node {
	if sl.IsEmpty() {
		return nil
	}
	cur, pre := sl.head, sl.head
	for level := sl.level; level >= 0; level-- {
		for cur = cur.next[level]; cur != nil && cur.key < target; pre, cur = cur, cur.next[level] {
		}
		if cur != nil && cur.key == target {
			return cur
		}
		cur = pre
	}
	if cur == sl.head { // <= 此处需要特判下
		return nil
	}
	return cur
}

func (sl *Skiplist) Insert(num int) {
	rLevel := randomLevel()
	level := maxInt(rLevel, sl.level)
	cur, pre := sl.head, sl.head
	updateNds := make([]*Node, level+1)
	for i := level; i >= 0; i-- {
		for cur = cur.next[i]; cur != nil && cur.key < num; pre, cur = cur, cur.next[i] {

		}
		updateNds[i] = pre
		cur = pre
	}
	nd := NewNode(num, num, maxLevel+1)
	for i := level; i >= 0; i-- {
		if updateNds[i] != nil {
			nd.next[i] = updateNds[i].next[i]
			updateNds[i].next[i] = nd
		}
	}
	sl.size++
	sl.level = maxInt(sl.level, level)
}

func (sl *Skiplist) Delete(num int) bool {
	upNds := make([]*Node, sl.level+1)
	cur, pre := sl.head, sl.head
	var nd *Node
	for level := sl.level; level >= 0; level-- {
		for cur = cur.next[level]; cur != nil && cur.key < num; pre, cur = cur, cur.next[level] {

		}
		if cur != nil && cur.key == num {
			upNds[level] = pre
			nd = cur
		}
		cur = pre
	}
	if nd == nil {
		return false
	}

	sl.size--
	for i := sl.level; i >= 0; i-- {
		if upNds[i] != nil {
			upNds[i].next[i] = nd.next[i]
		}
		if sl.head.next[i] == nil && i > 0 {
			sl.level--
		}
	}
	return true
}

func (sl *Skiplist) FindAbsNearest(target int) *Node {
	if sl.size == 0 {
		return nil
	}
	nd := sl.SearchLessEqual(target)
	if nd == nil {
		return sl.head.next[0]
	}
	if nd.key == target || nd.next[0] == nil {
		return nd
	}
	if target-nd.val <= nd.next[0].val-target {
		return nd
	}
	return nd.next[0]
}
