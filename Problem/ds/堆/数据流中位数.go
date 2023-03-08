package main

import "container/heap"

// 题目: 数据流不断输入过程中动态维护中位数
// 思路:
// 1. 双堆：小顶堆维护较大的一半数，大顶堆维护较小的一半数
// 2. 小顶堆的最小数 >= 大顶堆的最大数的
// 3. 小顶堆的元素数量 与 大顶堆元素数量差值不超过 1
// 4. 中位数: 大顶堆堆顶 或 (小顶堆堆顶 + 大顶堆堆顶) / 2

type MedianFinder struct {
	minHP MinHeap
	maxHP MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHP: make(MinHeap, 0),
		maxHP: make(MaxHeap, 0),
	}
}

// maxHeap 元素数 >= minHeap 元素数
func (this *MedianFinder) AddNum(num int) {
	if this.maxHP.Len() == 0 {
		heap.Push(&this.maxHP, num)
		return
	}

	// 双堆求解思路
	// 1. 左边堆：大顶堆
	// 2. 右边堆：小顶堆

	mv := this.maxHP.Top().(int)
	if mv >= num {
		heap.Push(&this.maxHP, num)
	} else {
		heap.Push(&this.minHP, num)
	}

	if this.maxHP.Len() > this.minHP.Len()+1 {
		heap.Push(&this.minHP, this.maxHP.Top().(int))
		heap.Pop(&this.maxHP)
	} else if this.minHP.Len() > this.maxHP.Len() {
		heap.Push(&this.maxHP, this.minHP.Top().(int))
		heap.Pop(&this.minHP)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHP.Len() == this.minHP.Len() {
		return float64(this.maxHP.Top().(int)+this.minHP.Top().(int)) / 2.0
	}
	return float64(this.maxHP.Top().(int))
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
