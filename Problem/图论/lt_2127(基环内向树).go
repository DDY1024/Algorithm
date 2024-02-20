package main

// 基环树的一些定义: https://blog.csdn.net/Q1410136042/article/details/81152191
// 1. 基环树: 树中带环
// 2. 基环内向树: 所有顶点有且只有一条出边，外在表现: 环外节点指向环内
// 3. 基环外向树: 所有顶点有且只有一条入边，外在表现: 环内节点指向环外
//
// 官方解题报告: https://leetcode-cn.com/problems/maximum-employees-to-be-invited-to-a-meeting/solution/can-jia-hui-yi-de-zui-duo-yuan-gong-shu-u8e8u/
// 1. 大小 >= 3 的环只能安排一个，因此必须选择最大的环
// 2. 大小 = 2 的环，安排两条最长的指向二元环的链路；且所有二元环结构可以共同参加会议

// 1. dfs 搜索进行求解
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

// 2. 拓扑排序方式 + 动态规划方法
// dp[i] = max{dp[j] + 1}, <j,i> 为一条有向边
//
// 基环内向树: 通过拓扑排序统一处理后，所有树枝均会被砍掉，剩余的节点构成一个个不相交的环
func maximumInvitationsTwo(favorite []int) int {
	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(favorite)
	vis := make([]bool, n)
	dis := make([]int, n)
	ind := make([]int, n)
	que := make([]int, 0, n)
	for i := 0; i < n; i++ {
		ind[favorite[i]]++
	}
	for i := 0; i < n; i++ {
		if ind[i] == 0 {
			que = append(que, i)
		}
	}

	// 拓扑序动态规划
	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		vis[u] = true
		ind[favorite[u]]--
		dis[favorite[u]] = maxInt(dis[favorite[u]], dis[u]+1)
		if ind[favorite[u]] == 0 {
			que = append(que, favorite[u])
		}
	}

	ret1 := 0
	for i := 0; i < n; i++ {
		if favorite[favorite[i]] == i {
			// 环上的节点在拓扑排序过程中可能访问不到
			// 因此 dis 存储的是环外节点的个数，需要额外 +1
			ret1 += dis[i] + 1
		}
	}

	// 由基环内向树的性质可知 --> 经过这一轮筛选之后，剩余节点构成了一个个环
	// 因此从环上任一顶点出发，计算环的长度即可
	ret2 := 0
	for i := 0; i < n; i++ {
		if !vis[i] {
			vis[i] = true
			cycleCnt, u := 1, i
			for !vis[favorite[u]] {
				vis[favorite[u]] = true
				u = favorite[u]
				cycleCnt++
			}
			ret2 = maxInt(ret2, cycleCnt)
		}
	}
	return maxInt(ret1, ret2)
}
