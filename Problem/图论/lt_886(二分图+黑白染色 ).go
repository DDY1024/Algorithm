package graph

// 题目链接：https://leetcode.cn/problems/possible-bipartition/
//
// 解题思路
// 		1. 根据题目中人与人之间的敌对关系，将所有人分成两组，每组人中不能存在敌对关系
//      2. 恰好符合二分图的性质，两组不同顶点间存在边，同一分组内不存在边
//      3. 问题转化成二分图判定，采用经典的黑白染色法进行判定
//      4. 判断奇环问题同样可以采用黑白染色法

// 经典的黑白染色判定二分图代码
func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for i := 0; i < len(dislikes); i++ {
		u, v := dislikes[i][0]-1, dislikes[i][1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	// 0: 未染色
	// 1: 黑色
	// 2: 白色
	color := make([]int, n)

	var dfs func(int, int) bool
	dfs = func(u, c int) bool {
		color[u] = c
		for _, v := range g[u] {
			if color[v] == c {
				return false // 存在奇环
			}

			// 该节点未染色，进行染色
			if color[v] == 0 && !dfs(v, c^3) {
				return false
			}
		}
		return true
	}

	// 遍历所有节点进行染色分组
	for i, c := range color {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}

	return true
}
