package main

// 参考资料: https://zhuanlan.zhihu.com/p/96621396
// dijkstra
// 		1. 贪心思路：每次选择最优点进行扩展（此处最有点选择过程可以是一维的，也可以是二维的）
// 			https://blog.csdn.net/u013081425/article/details/26020401 例如这题便是同时考虑
// 			距离和花费两个条件，我们在选择松弛节点的时候也需要同时考虑这两个条件
/*
for(int i = 1; i <= n; i++)
    {
		pos := -1
        for(int j = 1; j <= n; j++)
        {
            if(vis[j]) continue;
			if pos == -1 || dis[pos] > dis[j] {
				pos = j
			}
        }
        vis[pos] = 1;
        for(int j = 1; j <= n; j++)
        {
            if(vis[j]) continue;
			if dis[j] > dis[pos] + g[pos][j] {  // 与 prim 算法不同点
				dis[j] = dis[pos] + g[pos][j]
			}
        }
    }
*/

// 堆优化的 dijkstra 算法，复杂度 O(m * logm)

/*
type edge struct{ to, t int }
type pair struct{ v, dis int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v pair)          { heap.Push(h, v) }
func (h *hp) pop() pair            { return heap.Pop(h).(pair) }

func dijkstra(g [][]edge, start int) []int {
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = 1e9
	}
	dis[start] = 0
	h := hp{{start, 0}}
	for len(h) > 0 {
		vd := h.pop()
		v := vd.v
		if dis[v] < vd.dis {
			continue
		}
		for _, e := range g[v] {
			w, wt := e.to, e.t
			if newD := dis[v] + wt; newD < dis[w] {
				dis[w] = newD
				h.push(pair{w, dis[w]})
			}
		}
	}
	return dis
}
*/

/*
// C++ 实现
struct Polar{
    int dist, id;
    Polar(int dist, int id) : dist(dist), id(id){}
};

struct cmp{
    bool operator ()(Polar a, Polar b){ // 重载()运算符，使其成为一个仿函数
        return a.dist > b.dist;    // 这里是大于，使得距离短的先出队
    }
};
priority_queue<Polar, vector<Polar>, cmp> Q;

void Dij(int s)
{
    dist[s] = 0;
    Q.push(Polar(0, s));
    while (!Q.empty())
    {
        int p = Q.top().id;
        Q.pop();
        if (vis[p])  // 保证每个顶点只被处理一次（最小的）
            continue;
        vis[p] = 1;
        for (int e = head[p]; e != 0; e = edges[e].next)
        {
            int to = edges[e].to;

            // dist[to] = min(dist[to], dist[p] + edges[e].w);
            // if (!vis[to])
            //     Q.push(Polar(dist[to], to));

			if dist[to] > dist[p] + edges[e].w {
				dist[to] = dist[p] + edges[e].w
				Q.push(Polar(dist[to], to))
			}
        }
    }
}
*/
