package main

import "container/heap"

// 题目链接：https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid/description/
// 解题思路：https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid/solutions/2134066/zui-duan-lu-by-tsreaper-otha/

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Item struct {
	x, y, t int
}

type MinPQ []*Item

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Less(i, j int) bool {
	return pq[i].t < pq[j].t
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

// 1. 通过反复横跳来增加等待时间，不改变时间的奇偶性
// 2. 优先级队列实现最短路径的计算
func minimumTime(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	// 起始点的相邻点不可达，则终点不可达
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}

	dx, dy := []int{1, -1, 0, 0}, []int{0, 0, 1, -1}
	mark := make([][]bool, n)
	for i := 0; i < n; i++ {
		mark[i] = make([]bool, m)
	}

	pq := make(MinPQ, 0)
	heap.Push(&pq, &Item{0, 0, 0})
	mark[0][0] = true
	for pq.Len() > 0 {
		nd := heap.Pop(&pq).(*Item)
		if nd.x == n-1 && nd.y == m-1 {
			return nd.t
		}

		for i := 0; i < 4; i++ {
			xx, yy := nd.x+dx[i], nd.y+dy[i]
			if xx >= 0 && xx < n && yy >= 0 && yy < m {
				if !mark[xx][yy] && grid[xx][yy] <= nd.t+1 {
					mark[xx][yy] = true
					heap.Push(&pq, &Item{xx, yy, nd.t + 1}) // 可以直接到达（不需要等待）的格子只计算一次就好，多余的重复计算没有意义
					continue
				}

				// 通过和上一个格子反复横跳，等待一定时间后，可以到达当前格子
				if !mark[xx][yy] && grid[xx][yy] > nd.t+1 {
					// 下一个格子和当前格子时间奇偶性必然不同，根据 grid[xx][yy] - nd.t 的奇偶性来决定是否 +1
					heap.Push(&pq, &Item{xx, yy, grid[xx][yy] + ((grid[xx][yy]-nd.t)&1 ^ 1)})
				}
			}
		}
	}
	return -1
}
