package main

import (
	"container/heap"
	"sort"
)

// 题目链接: https://leetcode-cn.com/problems/course-schedule-iii/
// 解题思路
// 参考: https://leetcode-cn.com/problems/course-schedule-iii/solution/ke-cheng-biao-iii-by-leetcode-solution-yoyz/
// 1. 截止时间越早的课，理论上应该越先上，因为截止时间晚的课容忍度更高
//		x + c1 <= d1
//      x + c1 + c2 <= d2
//
//      x + c2 <= d2
//      x + c2 + c1 <= d1
// 假设两门课程都要争取上，当 d1 <= d2 时，显然先上 1 再上 2 更优，因此我们按照截止时间从小到大的顺序进行筛选。
// 同时，我们在选课的时候也要进行微调，假设在截止时间 x 之前，我们最多修 y 门课，那我们需要保证这 y 门课的总耗时最小，则为最优。
// 此处我们用优先级队列维护已经选择的课程，然后不断地进行淘汰和替换，保证选择的课程是最优的。
//
// 这种先选择，随后后悔替换的思路，是一种常见的贪心思路

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

func scheduleCourse(courses [][]int) int {
	// sort.Slice 同样也可以用于多维数组
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] == courses[j][1] {
			return courses[i][0] < courses[j][0]
		}
		return courses[i][1] < courses[j][1]
	})

	n, cost := len(courses), 0
	pq := make(MaxHeap, 0, n)
	for i := 0; i < n; i++ {
		if cost+courses[i][0] <= courses[i][1] {
			heap.Push(&pq, courses[i][0])
			cost += courses[i][0] // duration 累加
		} else if pq.Len() > 0 && pq.Top().(int) > courses[i][0] {
			cost += courses[i][0] - pq.Top().(int) // duration 累加
			heap.Pop(&pq)
			heap.Push(&pq, courses[i][0])
		}
	}
	return pq.Len()
}
