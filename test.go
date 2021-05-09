package main

import (
	"container/heap"
	"sort"
)

type Item struct {
	l, r int
}

type MinPQ []*Item

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Less(i, j int) bool {
	return (pq[i].r - pq[i].l) <= (pq[j].r - pq[j].l)
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

type Query struct {
	pos int
	idx int
}

func minInterval(intervals [][]int, queries []int) []int {
	n, m := len(intervals), len(queries)
	qs := make([]*Query, 0, m)
	for i := 0; i < m; i++ {
		qs = append(qs, &Query{
			pos: queries[i],
			idx: i,
		})
	}
	sort.Slice(qs, func(i, j int) bool {
		return qs[i].pos < qs[j].pos
	})

	intervalList := make([]*Item, 0, n)
	for i := 0; i < n; i++ {
		intervalList = append(intervalList, &Item{
			l: intervals[i][0],
			r: intervals[i][1],
		})
	}
	sort.Slice(intervalList, func(i, j int) bool {
		if intervalList[i].l == intervalList[j].l {
			return intervalList[i].r < intervalList[j].r
		}
		return intervalList[i].l < intervalList[j].l
	})

	hp, idx, ans := make(MinPQ, 0, n), 0, make([]int, m)
	for i := 0; i < m; i++ {
		for idx < n && intervalList[idx].l <= qs[i].pos {
			heap.Push(&hp, intervalList[idx])
			idx++
		}
		for hp.Len() > 0 && hp.Top().(*Item).r < qs[i].pos {
			heap.Pop(&hp)
		}
		if hp.Len() > 0 {
			item := hp.Top().(*Item)
			ans[qs[i].idx] = item.r - item.l + 1
		} else {
			ans[qs[i].idx] = -1
		}
	}
	return ans
}

func main() {

}
