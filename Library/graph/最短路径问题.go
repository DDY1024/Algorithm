package main

import (
	"container/heap"
)

// https://leetcode-cn.com/contest/weekly-contest-236/problems/minimum-sideway-jumps/
// 利用 priority_queue + 搜索求解最短路径（尤其在顶点比较多，边比较少的稀疏图中会有奇效）

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
	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			dis[i][j] = 0x3f3f3f3f
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

// 划分阶段，dp 进行求解
func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			dp[i][j] = 0x3f3f3f3f
		}
	}
	dp[0][2] = 0
	dp[0][1] = 1
	dp[0][3] = 1
	for i := 1; i < n; i++ {
		for j := 1; j <= 3; j++ {
			if obstacles[i] != j {
				dp[i][j] = minInt(dp[i][j], dp[i-1][j])
			}
		}
		min := minInt(dp[i][1], minInt(dp[i][2], dp[i][3]))
		if obstacles[i] != 1 {
			dp[i][1] = minInt(dp[i][1], min+1)
		}
		if obstacles[i] != 2 {
			dp[i][2] = minInt(dp[i][2], min+1)
		}
		if obstacles[i] != 3 {
			dp[i][3] = minInt(dp[i][3], min+1)
		}
	}
	return minInt(dp[n-1][1], minInt(dp[n-1][2], dp[n-1][3]))
}
