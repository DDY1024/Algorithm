package main

import "container/heap"

// 最大 K 个元素: 小顶堆
// 最小 K 个元素：大顶堆

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

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	hp := make(MinHeap, 0, n)
	for i := 0; i < n; i++ {
		if hp.Len() < k {
			heap.Push(&hp, nums[i])
		} else {
			if hp.Top().(int) < nums[i] {
				heap.Pop(&hp)
				heap.Push(&hp, nums[i])
			}
		}
	}
	return hp.Top().(int)
}
