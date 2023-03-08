package other

import (
	"math/rand"
	"time"
)

const (
	inf = 0x3f3f3f3f
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

// level 取值范围 [0, maxLevel], 其中第一层存储全部的数据元素，上层只存储指针
//
//  1. redis 中上升层数的概率为 1/4
//     rand() & 0xffff < 1/4 * 0xffff
func randomLevel() int {
	level := 0
	for rand.Intn(100) < 50 { // 1/2 概率爬升
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
		next: make([]*Node, n), // skiplist 最大高度
	}
}

type Skiplist struct {
	level int
	size  int
	head  *Node // 增加 head 节点，方便处理
}

func NewSkiplist() Skiplist {
	return Skiplist{
		level: 0,
		size:  0,
		head:  NewNode(-inf, -inf, maxLevel+1), // [0, maxLevel]
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
		if cur != nil && cur.key == key {
			return cur
		}
		// cur = nil 或 cur.key > key，在下层进行搜索时需要前移一下，赋值 pre
		cur = pre
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
		cur = pre
	}

	if cur == sl.head { // head 节点为了方便处理为保留节点，非真实节点
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
		if sl.head.next[i] == nil && i > 0 { // 维护更新最大层数
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

	// 最底层是整个完整的有序链表
	if key-nd.val <= nd.next[0].val-key {
		return nd
	}
	return nd.next[0]
}

// rank 扩展实现参考：https://github.com/gansidui/skiplist/blob/master/skiplist.go
// 扩展：对于 GetRank 扩展操作，我们采用如下方法
// 对于 skiplist 中每一层的每个节点维护其与同层下一个节点之间的距离；即节点 A --> 节点 B 中间需要经历多少个节点
// skiplist 中对于 key 的查找会形成一条搜索路径，如下图所示
//
//  h0                    -->     x6'    -->
//
//  h1 --> x1->x2         -->     x6'    -->
//             |
//  h2         x2-->x3    -->     x6'    -->
//                  |
//  h3              x3->x4 	  -->     x6'  -->
//                      |
//  h4                  x4 --> x5 --> x6' -->
//
// 由于我们本身维护了同层两个相邻节点之间的 span，因此我们按照搜索路径 h1 --> x1 --> x2 --> x3 --> x4 --> x5，将这些路径上相邻节点
// 之间的 span 求 sum，即可得到 x5 的 rank
//
// 1. 插入操作维护 span ?
// 	   画图可知，span 更新从第 0 层逐步到第 maxLevel 层
//	   在每层需要累计寻找该层 pre 节点所经过的 span 和
//
// 2. 删除操作如何维护 span ?
//     画图可知，span 更新从第 maxLevel 层逐步到第 0 层
//     相比于插入操作，删除操作只需要将 pre -> pre.next[i] -> pre.next[i].next[i] 的 span 值合并即可，并减去 1（删除节点）
