package main

import "container/heap"

// 题目链接：https://leetcode.cn/problems/find-median-from-data-stream/description/?envType=study-plan-v2&envId=top-100-liked

// 1. 双堆维护
// 		【大顶堆】维护较小的【一半】整数
//      【小顶堆】维护较大的【一半】整数

type MedianFinder struct {
	lh MaxHeap // 较小的
	rh MinHeap // 较大的
}

func Constructor() MedianFinder {
	return MedianFinder{
		lh: make(MaxHeap, 0),
		rh: make(MinHeap, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {

	if this.rh.Len() == 0 || this.rh.Top().(int) >= num {
		heap.Push(&this.lh, num)
	} else {
		heap.Push(&this.rh, num)
	}

	if this.lh.Len() > this.rh.Len()+1 {
		heap.Push(&this.rh, this.lh.Top().(int))
		heap.Pop(&this.lh)
		return
	}

	if this.rh.Len() > this.lh.Len() {
		heap.Push(&this.lh, this.rh.Top().(int))
		heap.Pop(&this.rh)
		return
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lh.Len() == this.rh.Len() {
		return float64(this.lh.Top().(int)+this.rh.Top().(int)) / 2.0
	}
	return float64(this.lh.Top().(int))
}

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
