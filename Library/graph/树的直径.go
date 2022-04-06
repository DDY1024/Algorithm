package graph

// 参考文章: https://oi-wiki.org/graph/tree-diameter/
//
// 定义: 树上任意两节点之间最长的简单路径即为树的直径
// 结论一: 显然，一棵树可以有多条直径，他们的长度相等
// 结论二: 如果树上所有边权均为正，则树的所有直径中点重合
//
//
// 求解方法
// 一、两次搜索
// 1. 从任意一点 u 出发，按照 dfs 或 bfs 搜索离 u 最远的顶点 v
// 2. 从顶点 v 出发，按照 dfs 或 bfs 搜索离 v 最远的顶点 w
// 3. 顶点 v --> w 路径长度即为树的直径
// 4. 顶点 v --> w 路径上中间的一个顶点或两个顶点即为树的中心
//
// 负权边
// 上述证明过程建立在所有路径均不为负的前提下。
// 如果树上存在负权边，则上述证明不成立。故若存在负权边，则无法使用两次 DFS 的方式求解直径。
//
// 伪代码如下
/*
const int N = 10000 + 10;

int n, c, d[N];
vector<int> E[N];

void dfs(int u, int fa) {
  for (int v : E[u]) {
    if (v == fa) continue;
    d[v] = d[u] + 1;
    if (d[v] > d[c]) c = v;
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
  dfs(1, 0);
  d[c] = 0, dfs(c, 0);
  printf("%d\n", d[c]);
  return 0;
}
*/

// 二、树形 dp
// 结论一: 显然最长路径的两个端点必然是叶子节点或根节点
// 结论二: 树形 DP 可以在存在负权边的情况下求解出树的直径
//
// d1(i): 以 i 为根节点的子树，向下延伸到达的最远节点
// d2(i): 以 i 为根节点的子树，向下延伸到达的次远节点
// 注意: 最远节点 和 次远节点不能在同一条路径上，因此可以得到下述状态转移方程
//
// d1(u) = max{d1(v)} + 1
// d2(u) = second_max{d1(v)} + 1
//
// 如何求解树直径上的节点？
// 1. 首先找到 d1(u) + d2(u) 最大的节点 u
// 2. 找到 d1(u) 的子节点，沿着该子节点向下寻找
// 3. 找到 d2(u) 的子节点，沿着该子节点向下寻找
// 4. 最终合并两部分的节点，即为一条树直径上的所有点
//
//
//
// 伪代码如下
/*
const int N = 10000 + 10;

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
  printf("%d\n", d);
  return 0;
}
*/

/************************************  相关题目  *************************************************/
// 1. https://leetcode-cn.com/problems/minimum-height-trees/ 求解树的中心（直径中间节点），1 个节点或两个节点
// 采用 dp 方式求解

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
	path := make([]int, maxD) // 树的直径上共有 maxD 个节点
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

	if maxD&1 > 0 {
		return []int{path[maxD>>1]}
	}
	return []int{path[(maxD>>1)-1], path[maxD>>1]}
}
