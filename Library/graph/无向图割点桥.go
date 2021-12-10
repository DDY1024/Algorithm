package main

// 参考：https://blog.csdn.net/lvshubao1314/article/details/43964889

type Edge struct {
	u, v, next int
	isBridge bool
}

// contest_154 4.go
func criticalConnections(n int, connections [][]int) [][]int {
	head := make([]int, n)
	for i := 0; i < n; i++ {
		head[i] = -1
	}
	edges, edgeNum := make([]Edge, 2*len(connections)), 0
	var addEdge = func(u, v int) {
		edges[edgeNum].u = u
		edges[edgeNum].v = v
		edges[edgeNum].next = head[u]
		head[u] = edgeNum
		edgeNum++
	}
	for i := 0; i < len(connections); i++ {
		addEdge(connections[i][0], connections[i][1])
		addEdge(connections[i][1], connections[i][0])
	}
	low, dfn, snum, isCurV, addBlock := make([]int, n), make([]int, n), 0, make([]bool, n), make([]int, n)
	var tarjan func(u, p int)
	tarjan = func(u, p int) {
		snum++
		low[u], dfn[u] = snum, snum   // low[u] 通过子树节点能够达到的最小dfs序号，dfn[u]即顶点 u 的 dfs 序号
		childNum := 0
		for i := head[u]; i != -1; i = edges[i].next {
			v := edges[i].v
			if v == p { continue }
			if dfn[v] == 0 {
				childNum++
				tarjan(v, u)
				if low[u] > low[v] {
					low[u] = low[v]
				}

				// 割边判断条件
				if low[v] > dfn[u] {
					edges[i].isBridge = true
					edges[i^1].isBridge = true
				}

				// 非根节点的割点判断条件
				if u != p && low[v] >= dfn[u] {
					isCurV[u] = true
					addBlock[u]++
				}

			} else if low[u] > dfn[v] {
				low[u] = dfn[v]
			}
		}
		// 根节点判断
		if u == p && childNum > 1 {
			isCurV[u] = true
		}
		if u == p {
			addBlock[u] = childNum - 1
		}
	}
	for i := 0; i < n; i++ {
		if dfn[i] == 0 {
			tarjan(i, i)
		}
	}
	var bridgeSet [][]int
	for i := 0; i < edgeNum; i += 2 {
		if edges[i].isBridge {
			bridgeSet = append(bridgeSet, []int{edges[i].u, edges[i].v})
		}
	}
	return bridgeSet
}

/*
// 提供一套 C++ 模板
void dfs(int u, int fa)
{
    int low[u] = pre[u] = ++dfs_block;
    int child = 0;
    for(int i = 0; i < G[u].size(); i++)
    {
        int v = G[u][i];
        if(!pre[v])
        {
            child++;
            dfs(v, u);
            low[u] = min(low[u], low[v]);
            if(low[v] >= pre[u])
                iscut[u] = true;
        }
        else if(v != fa)
            low[u] = min(low[u], pre[v]);
    }
    if(fa < 0 && child == 1) iscut[u] = false;
}
*/


// 边双联通分量
// 点双联通分量
// https://blog.csdn.net/fuyukai/article/details/51303292
// http://inceptions.me/2019/06/01/ACM-Graph/