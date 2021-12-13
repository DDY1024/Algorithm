package main

import "container/heap"

// 题目: 数据流不断输入过程中动态维护中位数
// 思路:
// 1. 双堆：小顶堆维护较大的一半数，大顶堆维护较小的一半数
// 2. 小顶堆的最小数 >= 大顶堆的最大数的
// 3. 小顶堆的元素数量 与 大顶堆元素数量差值不超过 1
// 4. 中位数: 大顶堆堆顶 或 (小顶堆堆顶 + 大顶堆堆顶) / 2

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Top() interface{} {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[0]
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Top() interface{} {
	if h.Len() == 0 {
		return nil
	}
	return (*h)[0]
}

type MedianManager struct {
	minH MinHeap
	maxH MaxHeap
}

func NewMedianManager() *MedianManager {
	return &MedianManager{
		minH: make(MinHeap, 0, 1000),
		maxH: make(MaxHeap, 0, 1000),
	}
}

func (self *MedianManager) Add(val int) {
	heap.Push(&self.minH, val)
	if self.minH.Len() > self.maxH.Len() {
		val = self.minH.Top().(int)
		heap.Pop(&self.minH)
		heap.Push(&self.maxH, val)
	}
}

func (self *MedianManager) Get() float64 {
	if self.minH.Len() == self.maxH.Len() {
		v1, v2 := self.minH.Top().(int), self.maxH.Top().(int)
		return float64(v1+v2) / 2.0
	}
	return float64(self.maxH.Top().(int))
}
