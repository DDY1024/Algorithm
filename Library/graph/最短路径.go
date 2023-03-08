package main

import (
	"container/heap"
)

// https://leetcode-cn.com/contest/weekly-contest-236/problems/minimum-sideway-jumps/
// 利用 priority_queue 进行扩展，求解最短路径

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Item struct {
	x int
	y int
	d int
}

type MinPQ []*Item

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Less(i, j int) bool {
	return pq[i].d < pq[j].d
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

func minSideJumpsOne(obstacles []int) int {
	n := len(obstacles)
	inf := 0x3f3f3f3f
	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			dis[i][j] = inf
		}
	}

	dis[0][2] = 0
	pq := make(MinPQ, 0, n*3+10)
	heap.Push(&pq, &Item{0, 2, 0})
	for pq.Len() > 0 {
		nd := heap.Pop(&pq).(*Item)
		if nd.x == n-1 {
			return nd.d
		}

		if nd.x+1 < n && obstacles[nd.x+1] != nd.y && dis[nd.x+1][nd.y] > nd.d {
			heap.Push(&pq, &Item{nd.x + 1, nd.y, nd.d})
			dis[nd.x+1][nd.y] = nd.d
		}

		for yy := 1; yy <= 3; yy++ {
			if yy != nd.y && obstacles[nd.x] != yy && dis[nd.x][yy] > nd.d+1 {
				heap.Push(&pq, &Item{nd.x, yy, nd.d + 1})
				dis[nd.x][yy] = nd.d + 1
			}
		}
	}
	return -1
}
