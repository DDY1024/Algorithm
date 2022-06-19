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

//
// level 取值范围 [0, maxLevel], 其中第一层存储全部的数据元素，高层只存储指针
// redis 中上升层数的概率为 1/4
// 		rand() & 0xffff < 1/4 * 0xffff
//
func randomLevel() int {
	level := 0
	for rand.Intn(100) < 50 { // 此例中以 1/2 的概率爬升
		level++
	}
	return minInt(level, maxLevel)
}

type Node struct {
	key  int
	val  int
	next []*Node
}

func NewNode(key, val, n int) *Node {
	return &Node{
		key:  key,
		val:  val,
		next: make([]*Node, n),
	}
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
		// 根据不同场景需要变更头节点初始化取值
		head: NewNode(-0x3f3f3f3f, -0x3f3f3f3f, maxLevel+1),
	}
}

func (sl *Skiplist) Search(key int) *Node {
	if sl.size == 0 {
		return nil
	}

	cur, pre := sl.head, sl.head
	for level := sl.level; level >= 0; level-- {
		for cur = cur.next[level]; cur != nil && cur.key < key; pre, cur = cur, cur.next[level] {
		}
		if cur != nil && cur.key == key { // 自定义比较
			return cur
		}
		cur = pre // 去下一层继续寻找
	}

	return nil
}

func (sl *Skiplist) SearchLessEqual(key int) *Node {
	if sl.size == 0 {
		return nil
	}

	cur, pre := sl.head, sl.head
	for level := sl.level; level >= 0; level-- {
		for cur = cur.next[level]; cur != nil && cur.key < key; pre, cur = cur, cur.next[level] {
		}
		if cur != nil && cur.key == key {
			return cur
		}
		cur = pre // 去下一层继续寻找
	}

	if cur == sl.head { // 排除 head 节点
		return nil
	}

	return cur
}

func (sl *Skiplist) Insert(key, val int) {

	rLevel := randomLevel()           // 获取该节点上升的层数
	level := maxInt(rLevel, sl.level) // 更新最大层数

	cur, pre := sl.head, sl.head
	updateNds := make([]*Node, level+1) // 先找到查找路径，然后进行插入操作
	for i := level; i >= 0; i-- {
		for cur = cur.next[i]; cur != nil && cur.key < key; pre, cur = cur, cur.next[i] {
		}

		updateNds[i] = pre // 按层数记录待更新的前驱节点
		cur = pre
	}

	nd := NewNode(key, val, maxLevel+1)
	for i := level; i >= 0; i-- {
		if updateNds[i] != nil {
			nd.next[i] = updateNds[i].next[i]
			updateNds[i].next[i] = nd
		}
	}

	sl.size++
	sl.level = maxInt(sl.level, level)
}

func (sl *Skiplist) Delete(key int) bool {

	var nd *Node                       // 保存待删除的节点
	upNds := make([]*Node, sl.level+1) // 删除操作同样需要保存更新路径
	cur, pre := sl.head, sl.head

	for level := sl.level; level >= 0; level-- {
		for cur = cur.next[level]; cur != nil && cur.key < key; pre, cur = cur, cur.next[level] {
		}

		// 查找到待删除的节点
		if cur != nil && cur.key == key {
			nd = cur
		}

		upNds[level] = pre
		cur = pre
	}

	// 待删除节点不存在
	if nd == nil {
		return false
	}

	// 更新路径上的节点指向信息
	for i := sl.level; i >= 0; i-- {
		if upNds[i] != nil {
			upNds[i].next[i] = nd.next[i]
		}
		if sl.head.next[i] == nil && i > 0 { // 更新最高层数
			sl.level--
		}
	}

	sl.size--
	return true
}

func (sl *Skiplist) FindAbsNearest(key int) *Node {
	if sl.size == 0 {
		return nil
	}

	nd := sl.SearchLessEqual(key)
	if nd == nil {
		return sl.head.next[0] // 没有 <=, 直接返回第一个元素
	}

	if nd.key == key || nd.next[0] == nil {
		return nd
	}

	if key-nd.val <= nd.next[0].val-key {
		return nd
	}

	return nd.next[0]
}
