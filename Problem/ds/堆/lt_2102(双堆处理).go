package main

import "container/heap"

// 题目链接: https://leetcode-cn.com/problems/sequentially-ordinal-rank-tracker/
// 解题思路: 双堆、小顶堆、大顶堆
// 1. 查询第 i 大元素，其中 i 是单调递增的，且每次 + 1
// 2. 我们完全可以维护更新前 i 大元素，另外维护剩余元素的最大值，第 i + 1 次查询，我们直接从剩余最大元素中取即可
// 3. 最终我们采用 小顶堆 + 大顶堆 双堆的思想来进行求解

// 扩展到一般性问题: 数据流第 k 大元素，其中查询 k 是单调递增的，我们均可以采用双堆的思路进行求解

type Item struct {
	name  string
	score int
}

type MaxPQ []*Item

func (pq MaxPQ) Len() int { return len(pq) }

func (pq MaxPQ) Less(i, j int) bool {
	if pq[i].score == pq[j].score {
		return pq[i].name <= pq[j].name
	}
	return pq[i].score > pq[j].score
}

func (pq MaxPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MaxPQ) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *MaxPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *MaxPQ) Top() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	return (*pq)[0]
}

type MinPQ []*Item

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Less(i, j int) bool {
	if pq[i].score == pq[j].score {
		return pq[i].name > pq[j].name
	}
	return pq[i].score < pq[j].score
}

func (pq MinPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinPQ) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *MinPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *MinPQ) Top() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	return (*pq)[0]
}

type SORTracker struct {
	maxPQ MaxPQ
	minPQ MinPQ
}

func Constructor() SORTracker {
	return SORTracker{
		maxPQ: make(MaxPQ, 0, 20000),
		minPQ: make(MinPQ, 0, 20000),
	}
}

// minPQ: 维护前 i 大的元素
// maxPQ: 维护剩余的元素
// 操作过程中保证: minPQ 最小元素 > maxPQ 最大元素
// 对于每次 add 操作，我们首先应该加入 minPQ 更新前 i 大元素，并将堆顶元素移动至 maxPQ
func (this *SORTracker) Add(name string, score int) {
	heap.Push(&this.minPQ, &Item{name, score})
	item := this.minPQ.Top().(*Item)
	heap.Pop(&this.minPQ)
	heap.Push(&this.maxPQ, item)
}

// 对于每次 get 操作，由于 minPQ 中已经存在前 i 大元素，因此我们只需要将 maxPQ 的最大元素移动至 minPQ 即可
func (this *SORTracker) Get() string {
	item := this.maxPQ.Top().(*Item)
	heap.Pop(&this.maxPQ)
	heap.Push(&this.minPQ, item)
	return item.name
}
