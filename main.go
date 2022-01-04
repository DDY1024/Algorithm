package main

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

func maximumInvitations(favorite []int) int {
	n, ret := len(favorite), 0
	dis := make([]int, n)
	label := make([]int, n)
	time := 0
	for i := 0; i < n; i++ {
		dis[i] = -1
		label[i] = -1
	}

	// 1. 第一种情况: 求解长度最大的有向环
	var maxCycle func(x, d int)
	maxCycle = func(x, d int) {
		dis[x] = d
		label[x] = time
		if dis[favorite[x]] != -1 { // 该节点已经被遍历过
			if label[x] == label[favorite[x]] { // 相邻两个顶点必须属于同一次遍历过程，否则会计算错误
				ret = maxInt(ret, dis[x]-dis[favorite[x]]+1)
			}
			return
		}
		maxCycle(favorite[x], d+1)
	}

	for i := 0; i < n; i++ {
		if dis[i] == -1 {
			time++ // 标识第几次遍历
			maxCycle(i, 0)
		}
	}

	// 2. 所有 "链条 --> A --> B <-- 链条"
	//                  A <-- B
	// 总长度
	// 构建逆向图进行求解
	// 注意图性质: 每个顶点有且仅存在一条出边，因此该顶点仅可能出现在唯一的环或链条中
	// 所有的环之间是相互独立的，不存在任何可达路径；即同一个连通分量不会存在两个或更多的环
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		if favorite[favorite[i]] != i { // 二元环不需要建边
			adj[favorite[i]] = append(adj[favorite[i]], i)
		}
	}

	var maxPath func(x int) int
	maxPath = func(x int) int {
		ret := 1
		for _, v := range adj[x] {
			ret = maxInt(ret, maxPath(v)+1) // 保证有向无环
		}
		return ret
	}

	total := 0
	for i := 0; i < n; i++ {
		if favorite[favorite[i]] == i {
			total += maxPath(i)
		}
	}
	return maxInt(ret, total)
}

// 分析图的性质，然后分类进行讨论计算，求取最终的结果
func main() {

}
