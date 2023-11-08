package main

// 1. 割点
// 2. 割边
// 3. 点双连通分量
// 4. 边双连通分量

type Edge struct {
	u, v, next int
	isBridge   bool
}

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

	low := make([]int, n) // low[u] 通过子树节点能够达到的最小 dfn 序号
	dfn := make([]int, n)
	vNum, isCurV, addBlock := 0, make([]bool, n), make([]int, n)

	var tarjan func(u, p int)
	tarjan = func(u, p int) {
		vNum++
		low[u], dfn[u] = vNum, vNum

		child := 0
		for i := head[u]; i != -1; i = edges[i].next {
			v := edges[i].v
			if v == p {
				continue
			}

			if dfn[v] == 0 {
				child++
				tarjan(v, u)
				// 通过子节点能够到达的最小 dfn 序号
				if low[u] > low[v] {
					low[u] = low[v]
				}

				// 割边判断
				if low[v] > dfn[u] {
					edges[i].isBridge = true
					edges[i^1].isBridge = true
				}

				// 非根节点的割点判断
				if u != p && low[v] >= dfn[u] {
					isCurV[u] = true
					addBlock[u]++
				}

			} else if low[u] > dfn[v] {
				low[u] = dfn[v]
			}
		}

		// 根节点为割点当且仅当其孩子节点数 > 1
		if u == p && child > 1 {
			isCurV[u] = true
		}
		if u == p {
			addBlock[u] = child - 1
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
