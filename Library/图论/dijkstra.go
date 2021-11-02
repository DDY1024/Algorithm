package main

import (
	"container/heap"
	"fmt"
)

// TODO: 待修改成更优雅的方式

type Node struct {
	d int // 距离 d
	v int // 顶点 v
}

// NodeHeap 顶点堆
type NodeHeap []Node

// Len 长度
func (sh *NodeHeap) Len() int { return len(*sh) }

// Less 比较函数
// < : 小顶堆
// > : 大顶堆
func (sh *NodeHeap) Less(i, j int) bool { return (*sh)[i].d < (*sh)[j].d }

// Swap 交换函数
func (sh *NodeHeap) Swap(i, j int) { (*sh)[i], (*sh)[j] = (*sh)[j], (*sh)[i] }

// Pop heap pop
func (sh *NodeHeap) Pop() interface{} {
	n := len(*sh)
	x := (*sh)[n-1]
	*sh = (*sh)[:n-1]
	return x
}

// Push heap push
func (sh *NodeHeap) Push(x interface{}) {
	*sh = append(*sh, x.(Node))
}

// heap.Init()
// heap.Pop()
// heap.Push()

const (
	maxN = 1010
	maxE = 1010 * 1010 / 2
	inf  = 0x3f3f3f3f
)

type Edge struct {
	u    int
	v    int
	w    int
	next int
}

var (
	dis     [maxN][2]int
	edges   = [maxE]Edge{}
	head    = [maxN]int{}
	edgeNum int
)

// n : 顶点个数，编号 0 ~ n-1
func initGraph(n int) {
	edgeNum = 0
	for i := 0; i < n; i++ {
		head[i] = -1
	}
}

// 无向图加边
// 有向图加边
func addEdge(u, v, w int) {

	edges[edgeNum] = Edge{u: u, v: v, w: w}
	edges[edgeNum].next = head[u]
	head[u] = edgeNum
	edgeNum++

	edges[edgeNum] = Edge{u: v, v: u, w: w}
	edges[edgeNum].next = head[v]
	head[v] = edgeNum
	edgeNum++
}

// n: 顶点编号 0 ~ n-1
// src: 源点
// 注意: 次短路径的求解过程中，路径可以重复走
func solve(n int, src int) {

	var (
		ndHeap = &NodeHeap{}
		nd     Node
		u      int
		v      int
		d      int
	)

	for i := 0; i < n; i++ {
		dis[i][0] = inf // 最短路径
		dis[i][1] = inf // 次短路径
	}

	// dis[src][0], dis[src][1] = 0, 0
	dis[src][0] = 0
	*ndHeap = append(*ndHeap, Node{0, src})
	heap.Init(ndHeap)
	for len(*ndHeap) > 0 {
		nd = heap.Pop(ndHeap).(Node)
		if dis[nd.v][1] < nd.d { // 注意: 此处只能取 < ，而不能取 <=
			continue
		}
		u = nd.v
		for i := head[u]; i != -1; i = edges[i].next {
			v = edges[i].v
			d = nd.d + edges[i].w
			if dis[v][0] > d {
				// TODO: 维护最短路径上前继节点
				dis[v][0], d = d, dis[v][0]
				heap.Push(ndHeap, Node{dis[v][0], v})
			}
			if dis[v][0] <= d && dis[v][1] > d { // 注意: dis[v][0] <= d 而不是 dis[v][0] < d
				// TODO: 维护次短路径上前继节点
				dis[v][1] = d
				heap.Push(ndHeap, Node{dis[v][1], v})
			}
		}
	}
	return
}

func main() {

	// 1. 最短路径唯一
	initGraph(3)
	addEdge(0, 1, 1)
	addEdge(1, 2, 2)
	solve(3, 0)
	for i := 0; i < 3; i++ {
		fmt.Println("Test: ", i, dis[i][0], dis[i][1])
	}

	fmt.Println("--------------------------")
	// 2. 最短路径和次短路径长度不相同
	initGraph(4)
	addEdge(0, 1, 1)
	addEdge(0, 2, 2)
	addEdge(1, 3, 1)
	addEdge(2, 3, 1)
	solve(4, 0)
	for i := 0; i < 4; i++ {
		fmt.Println("Test: ", i, dis[i][0], dis[i][1])
	}

	fmt.Println("--------------------------")
	// 3. 最短路径和次短路径长度相同
	initGraph(4)
	addEdge(0, 1, 1)
	addEdge(0, 2, 1)
	addEdge(1, 3, 1)
	addEdge(2, 3, 1)
	solve(4, 0)
	for i := 0; i < 4; i++ {
		fmt.Println("Test: ", i, dis[i][0], dis[i][1])
	}
}

// dijkstra 算法本质上是一种贪心算法，每次总是选择当前有效顶点集合的最优点进行扩展。
// 每次贪心选择时候要考虑的条件可以是一维的，也可以而二维的，需要根据题目要求。
// https://blog.csdn.net/u013081425/article/details/26020401 例如这题便是同时考虑
// 距离和花费两个条件，我们在选择松弛节点的时候也需要同时考虑这两个条件
/*
for(int i = 1; i <= n; i++)
    {
        int M1= INF, M2 = INF, pos;
        for(int j = 1; j <= n; j++)
        {
            if(vis[j]) continue;
			// 综合考虑两个条件
            if(dis1[j] < M1 || (dis1[j] == M1 && dis2[j] < M2))
            {
                M1 = dis1[j];
                M2 = dis2[j];
                pos = j;
            }
        }
        vis[pos] = 1;
        for(int j = 1; j <= n; j++)
        {
            if(vis[j]) continue;
            int tmp1 = dis1[pos] + Map[pos][j];
            int tmp2 = dis2[pos] + cost[pos][j];
            if(tmp1 < dis1[j] || (tmp1 == dis1[j] && tmp2 < dis2[j]))
            {
                dis1[j] = tmp1;
                dis2[j] = tmp2;
            }
        }
    }
*/
