package graph

// 参考文章: https://oi-wiki.org/graph/tree-diameter/
//
// 树的直径：树上任意两点之间最长的简单路径
// 相关结论
//		一棵树可以存在多条直径且长度相等
//		如果树的【边权为正整数】，则所有直径的重点重合
//
// 1. 两次搜索算法（非负权边）
// 		a. 从任意一点 u 出发，按照 dfs 或 bfs 搜索离 u 最远的顶点 v
//      b. 从顶点 v 出发，按照 dfs 或 bfs 搜索离 v 最远的顶点 w
//      c. 顶点 v 到 w 的简单路径长度，即为树的直径
//      d. 顶点 v 到 w 路径的中点（一个或两个）即为树的中心
//	注意：如果树中存在负权边，则不能采用此算法求解
/*
int n, c, d[N];
vector<int> E[N];

void dfs(int u, int fa) {
  for (int v : E[u]) {
    if (v == fa) continue;  // 防止走回父节点
    d[v] = d[u] + 1;
    if (d[v] > d[c]) c = v;  // 更新距离最远的顶点
    dfs(v, u);
  }
}

int main() {
  scanf("%d", &n);
  for (int i = 1; i < n; i++) {
    int u, v;
    scanf("%d %d", &u, &v);
    E[u].push_back(v), E[v].push_back(u);
  }

  dfs(1, 0); // 第一次搜索
  d[c] = 0;
  dfs(c, 0); // 第二次搜索
}
*/

// 2. 树形 DP
// 		最长路径的两个端点必然是叶子节点或根节点
//		针对负权边场景，仍然能够正确求解
//
//	状态定义
// 		d1(i): 以 i 为根节点的子树，向下延伸到达的【最远】节点
// 		d2(i): 以 i 为根节点的子树，向下延伸到达的【次远】节点
//	状态转移
// 		d1(u) = max{d1(v)} + 1
// 		d2(u) = second_max{d1(v)} + 1
//	树的直径 = max{d1(u)+d2(u)}
//
// 	如何求解树的直径上的节点
//		1. 找到节点 u
//      2. 找到 d1(u) 的最大子节点
//		3. 找到 d2(u) 的最大子节点
//      4. 根据 2 和 3 中子节点的【向下最长路径】的节点构成
/*
int n, d = 0;
int d1[N], d2[N];
vector<int> E[N];

void dfs(int u, int fa) {
  d1[u] = d2[u] = 0;
  for (int v : E[u]) {
    if (v == fa) continue;
    dfs(v, u);
    int t = d1[v] + 1;
    if (t > d1[u])
      d2[u] = d1[u], d1[u] = t;
    else if (t > d2[u])
      d2[u] = t;
  }
  d = max(d, d1[u] + d2[u]);
}
int main() {
  scanf("%d", &n);
  for (int i = 1; i < n; i++) {
    int u, v;
    scanf("%d %d", &u, &v);
    E[u].push_back(v), E[v].push_back(u);
  }
  dfs(1, 0);
}
*/

/************************************  相关题目  *************************************************/
// 1. https://leetcode-cn.com/problems/minimum-height-trees/
// 树的直径的中间节点构成的树高度最小（一个或两个）

func findMinHeightTrees(n int, edges [][]int) []int {
	m := len(edges)
	adj := make([][]int, n)
	for i := 0; i < m; i++ {
		u, v := edges[i][0], edges[i][1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	d1 := make([]int, n)
	d2 := make([]int, n)
	c1 := make([]int, n) // 最远路径的孩子节点
	c2 := make([]int, n) // 次远路径的孩子节点
	maxD := 0

	var dfs func(u, p int)
	dfs = func(u, p int) {
		d1[u], d2[u], c1[u], c2[u] = 0, 0, -1, -1
		for _, v := range adj[u] {
			if v == p {
				continue
			}
			dfs(v, u)
			tmp := d1[v] + 1
			if tmp > d1[u] {
				d2[u] = d1[u]
				c2[u] = c1[u]
				d1[u] = tmp
				c1[u] = v
			} else if tmp > d2[u] {
				d2[u] = tmp
				c2[u] = v
			}
		}
		maxD = maxInt(maxD, d1[u]+d2[u]+1)
	}
	dfs(0, -1)

	// 找出树直径上的节点
	path := make([]int, maxD)
	for i := 0; i < n; i++ {
		if d1[i]+d2[i]+1 == maxD {
			path[d1[i]] = i
			// 沿着最远路径进行寻找
			u, idx := c1[i], d1[i]-1
			for u != -1 {
				path[idx] = u
				idx--
				u = c1[u]
			}

			// 沿着次远路径进行寻找
			u, idx = c2[i], d1[i]+1
			for u != -1 {
				path[idx] = u
				idx++
				u = c1[u]
			}
			break
		}
	}

	// 寻找树的直径的中间节点
	if maxD&1 > 0 {
		return []int{path[maxD>>1]}
	}

	return []int{path[(maxD>>1)-1], path[maxD>>1]}
}
