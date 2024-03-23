package other

import (
	"math"
	"math/rand"
	"time"
)

const (
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

// Redis 元素上升概率：rand() & 0xffff < 0xffff/4
func randomLevel() int {
	level := 0
	for level < maxLevel && rand.Intn(100) < 50 { // 注意需要设置下 maxLevel 上限限制，避免一些无效计算
		level++
	}
	return minInt(level, maxLevel)
}

type Node struct {
	key  int
	val  int
	next []*Node
}

func NewNode(key int, val int, level int) *Node {
	return &Node{
		key:  key,
		val:  val,
		next: make([]*Node, level+1),
	}
}

type Skiplist struct {
	level int
	size  int
	head  *Node // head 节点
}

func NewSkiplist() Skiplist {
	return Skiplist{
		level: 0,
		size:  0,
		head:  NewNode(math.MinInt, math.MinInt, maxLevel), // head 节点存储 key 的最小值
	}
}

func (sl *Skiplist) Search(key int) *Node {
	if sl.size == 0 {
		return nil
	}

	cur, prev := sl.head, sl.head
	for level := sl.level; level >= 0; level-- {
		for cur = cur.next[level]; cur != nil && cur.key < key; prev, cur = cur, cur.next[level] {
		}

		if cur != nil && cur.key == key {
			return cur
		}

		// 在下一层搜索时，优先赋值 prev
		cur = prev
	}

	return nil
}

func (sl *Skiplist) SearchLessEqual(key int) *Node {
	if sl.size == 0 {
		return nil
	}

	cur, prev := sl.head, sl.head
	for level := sl.level; level >= 0; level-- {
		for cur = cur.next[level]; cur != nil && cur.key < key; prev, cur = cur, cur.next[level] {
		}

		if cur != nil && cur.key == key {
			return cur
		}

		cur = prev
	}

	if cur == sl.head {
		return nil
	}
	return cur
}

func (sl *Skiplist) Insert(key, val int) {
	level := maxInt(randomLevel(), sl.level)
	cur, prev := sl.head, sl.head

	// 1. 寻找每层的前驱节点
	updateNds := make([]*Node, level+1)
	for i := level; i >= 0; i-- {
		for cur = cur.next[i]; cur != nil && cur.key < key; prev, cur = cur, cur.next[i] {
		}
		updateNds[i] = prev
		cur = prev
	}

	// 2. 按照层次依次插入新节点
	nd := NewNode(key, val, maxLevel)
	for i := level; i >= 0; i-- {
		// Tips：这个判断是没有必要的，因为 sl.head 永远可以兜底，updateNds[i] 永远不为 nil
		if updateNds[i] != nil {
			nd.next[i] = updateNds[i].next[i]
			updateNds[i].next[i] = nd
		}
	}

	sl.size++
	sl.level = maxInt(sl.level, level)
}

func (sl *Skiplist) Delete(key int) bool {
	var (
		nd        *Node
		updateNds = make([]*Node, sl.level+1)
		cur       = sl.head
		prev      = sl.head
	)

	// 1. 查找待删除的节点，保存前驱节点路径
	for i := sl.level; i >= 0; i-- {
		for cur = cur.next[i]; cur != nil && cur.key < key; prev, cur = cur, cur.next[i] {
		}

		// 找到待删除的节点
		if cur != nil && cur.key == key {
			nd = cur
		}

		updateNds[i] = prev
		cur = prev
	}

	if nd == nil {
		return false
	}

	// 2. 删除每层的节点
	for i := sl.level; i >= 0; i-- {
		if updateNds[i] != nil {
			updateNds[i].next[i] = nd.next[i]
		}
		// 当前层已经没有节点，更新总层数
		if sl.head.next[i] == nil && i > 0 {
			sl.level--
		}
	}

	sl.size--
	return true
}

// 1. 只有 <=
// 2.  只有 >
// 3.  <= 和 > 均有
func (sl *Skiplist) FindAbsNearest(key int) *Node {
	if sl.size == 0 {
		return nil
	}

	nd := sl.SearchLessEqual(key)
	// 不存在 <= key 的元素，则第一个元素最接近
	if nd == nil {
		return sl.head.next[0] // 没有 <=, 直接返回第一个元素
	}

	// 查找到相同的 key 或 不存在 > key 的节点
	if nd.key == key || nd.next[0] == nil {
		return nd
	}

	// 比较 < key 和 > key 的节点距离，选择距离最近的节点
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
