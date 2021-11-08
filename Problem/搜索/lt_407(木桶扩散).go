package main

import "container/heap"

// 木桶扩散原理: 每次总是选择最低的位置进行扩散
// 解题思路参考: https://leetcode-cn.com/problems/trapping-rain-water-ii/solution/jie-yu-shui-ii-by-leetcode-solution-vlj3/
type Item struct {
	h int
	x int
	y int
}

type MinPQ []*Item

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Less(i, j int) bool {
	return pq[i].h <= pq[j].h
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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trapRainWater(heightMap [][]int) int {
	n, m, ans := len(heightMap), len(heightMap[0]), 0
	dx, dy := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}
	visit := make([][]bool, n)
	for i := 0; i < n; i++ {
		visit[i] = make([]bool, m)
	}

	pq := make(MinPQ, 0, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || i == n-1 || j == 0 || j == m-1 {
				heap.Push(&pq, &Item{heightMap[i][j], i, j})
				visit[i][j] = true
			}
		}
	}

	// 我们按照单调不减的方向进行扩展 --> 因为最终内部的高度肯定是 >= 外围最小高度的 --> 想想木桶盛水的情形
	// 木桶扩散方法 --> 总是从当前外围最低的位置进行扩散
	// 每次选择高度最低的点进行扩展，这样得到的扩展点是正确的
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		for i := 0; i < 4; i++ {
			x, y := item.x+dx[i], item.y+dy[i]
			if x >= 0 && x < n && y >= 0 && y < m && !visit[x][y] {
				ans += maxInt(0, item.h-heightMap[x][y]) // 实际高度 < 接水后的高度
				visit[x][y] = true
				// max 操作
				heap.Push(&pq, &Item{maxInt(item.h, heightMap[x][y]), x, y})
			}
		}
	}
	return ans
}

/*
class Solution {
public:
    priority_queue<pair<int, pair<int, int> >, vector<pair<int, pair<int, int> > >, greater<pair<int, pair<int, int> > > > pq;
    bool vis[201][201];
    int dirx[4] = {0, 0, 1, -1};
    int diry[4] = {1, -1, 0, 0};
    int trapRainWater(vector<vector<int>>& heightMap) {
        int n = heightMap.size(), m = heightMap[0].size(), ans = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (i == 0 || i == n - 1 || j == 0 || j == m - 1) {
                    pq.push(make_pair(heightMap[i][j], make_pair(i, j)));
                    vis[i][j] = true;
                }
            }
        }

        while (!pq.empty()) {
            int h = pq.top().first, x = pq.top().second.first, y = pq.top().second.second;
            pq.pop();
            for (int i = 0; i < 4; i++) {
                int nx = dirx[i] + x, ny = diry[i] + y;
                if (nx >= 0 && nx < n && ny >= 0 && ny < m && (!vis[nx][ny])) {
                    if (heightMap[nx][ny] < h) {
                        ans += h - heightMap[nx][ny];
                    }
                    vis[nx][ny] = 1;
                    pq.push(make_pair(max(heightMap[nx][ny], h), make_pair(nx, ny)));
                }
            }
        }
        return ans;
    }
};
*/
