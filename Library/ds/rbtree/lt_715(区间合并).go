package main

// import "github.com/emirpasic/gods/trees/redblacktree"

// // 解题报告: https://leetcode.cn/problems/range-module/solution/range-mo-kuai-by-leetcode-solution-4utf/
// // 经典的区间 合并/拆分 问题
// //
// // 利用 rbtree（红黑树）[按序维护一系列不相交的区间]
// // 在插入和删除某一个区间时，进行区间的合并和拆分操作
// //
// // 核心点在于: 我们要按顺序维护一系列不相交的区间，这样方便我们针对一个给定的查询区间判断是否完全覆盖

// type RangeModule struct {
// 	*redblacktree.Tree
// }

// func Constructor() RangeModule {
// 	return RangeModule{
// 		redblacktree.NewWithIntComparator(),
// 	}
// }

// // 针对 Add 和 Remove 操作需要多画画图，理清楚大小关系

// func (t RangeModule) AddRange(left, right int) {

// 	// 1. 存在 <= left 的最右边的区间
// 	if node, ok := t.Floor(left); ok {
// 		r := node.Value.(int)
// 		if r >= right { // 存在 [l, r) 区间包含 [left, right)，则直接返回
// 			return
// 		}

// 		// 如果 r >= left，即 l <= left <= r <= right，那我们直接以 [l, right) 区间参与合并
// 		if r >= left {
// 			left = node.Key.(int)
// 			t.Remove(left) // 剔除区间 [l,r)
// 		}

// 		// 如果 r < left，则直接以 [left, right) 参与合并
// 	}

// 	// 2. 如果不存在，则直接以 [left, right) 参与合并

// 	// 3. 以 [left, right) 为起始区间，不断地寻找 l' >= left 且 l' <= right 的区间进行合并，并剔除原先的区间节点
// 	for node, ok := t.Ceiling(left); ok && node.Key.(int) <= right; node, ok = t.Ceiling(left) {
// 		right = max(right, node.Value.(int))
// 		t.Remove(node.Key)
// 	}

// 	// 插入最终合并后的区间
// 	t.Put(left, right)
// }

// // 删除操作同样类似于插入操作，存在拆分区间的可能性
// func (t RangeModule) RemoveRange(left, right int) {

// 	// 1. 存在 <= left 的最右边的区间
// 	if node, ok := t.Floor(left); ok {
// 		l, r := node.Key.(int), node.Value.(int)
// 		// 1. 区间包含，直接区间内部消化即可，最多余留一个区间
// 		if r >= right {
// 			if l == left {
// 				t.Remove(l)
// 			} else {
// 				node.Value = left // 相当于剩余区间变为 [l,left)
// 			}

// 			// 右边 [right, r) 不为空区间，则直接插入
// 			if right != r {
// 				t.Put(right, r)
// 			}
// 			return
// 		}

// 		// 2. 区间不包含：先将原先区间变为 [l, left)，剩余 [left, right) 继续参与后续区间的删除操作
// 		if r > left {
// 			node.Value = left
// 		}
// 	}

// 	// 以 [left, right) 参与后续区间的删除操作
// 	for node, ok := t.Ceiling(left); ok && node.Key.(int) < right; node, ok = t.Ceiling(left) {
// 		r := node.Value.(int)
// 		// 如果 [l, r) 区间仍然被 [left, right) 包含，则直接删除即可
// 		t.Remove(node.Key)
// 		// [l, r) 区间与 [left, right) 部分重叠，此时便为待处理的最后一个区间
// 		if r > right {
// 			t.Put(right, r) // [right, r) 不为空，则直接插入一个新区间，处理结束
// 			break
// 		}
// 	}
// }

// // 查找比较容易：直接定位到某个区间，查看是否包含即可
// func (t RangeModule) QueryRange(left, right int) bool {
// 	node, ok := t.Floor(left) // <= left 最大值，然后查看区间右端点是否包含 right
// 	return ok && node.Value.(int) >= right
// }

// func max(a, b int) int {
// 	if b > a {
// 		return b
// 	}
// 	return a
// }
