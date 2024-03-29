package main

import (
	"math"
	"sort"
	"strconv"
)

// 欧拉回路：通过图中每条边恰好一次的回路
// 欧拉通路：通过图中每条边恰好一次的通路
// 欧拉图：  具有欧拉回路的图
// 半欧拉图：具有欧拉通路，但不具有欧拉回路的图

// 性质
// 1. 欧拉图中所有顶点的度数是【偶数】
// 2. 若 G 是欧拉图，则它为若干个环的并，且每条边被包含在【奇数】个环中

// 判定方法
// 1. 无向图是欧拉图当且仅当
//       a. 连通图
//       b. 顶点的度数是偶数
// 2. 无向图是半欧拉图当且仅当
//      a. 连通图
//      b. 恰好有 0 或 2 个奇度顶点
// 3. 有向图是欧拉图当且仅当
//      a. 非零度顶点【强连通】
//      b. 顶点的入度和出度相等
// 4. 有向图是半欧拉图当且仅当
//      a. 非零度顶点【弱连通】
//      b. 至多一个顶点的入度与出度差 1，入度 - 出度 = 1
//      c. 至多一个顶点的出度与入度差 1，出度 - 入度 = 1
//      d. 其它顶点的入度和出度相等

// Fleury 算法（不考虑）
//  每次选择下一条边的时候，优先选择不是桥的边
//      由于对于桥边的处理比较复杂，暂且不考虑该算法
//
//
//  Hierholzer 算法（逐步插入回路算法）
/*
struct Edge {
    int v, next;
    Edge() {}
    Edge (int v, int next) : v (v), next (next) {}
} edges[MAXE << 1];

int head[MAXN], tot;  // 链式前向星存储结构
int path[MAXN], cnt; // 记录欧拉回路经过的节点
int deg[MAXN]; // 存储顶点度数
bool used[MAXE]; // 标记边是否被访问过

// 初始化
void init() {
    tot = 0;
    cnt = 0;
    memset (head, -1, sizeof (head) );
    memset (deg, 0, sizeof (deg) );
    memset (used, false, sizeof (used) );
}

// 加边
void add_edge (int u, int v) {
    edges[tot] = Edge(v, head[u]);
    head[u] = tot++;
}

// 欧拉回路求解
// 如果需要输出字典序最小的欧拉路径，需要首先对边进行排序
void dfs (int u) {
    for (int i = head[u]; ~i; i = edges[i].next) {
        if (!used[i]) {
            used[i] = used[i ^ 1] = true;  // 无向图：正向边/反向边
            // used[i] = true // 有向图：单向边
            dfs (edges[i].v);
        }
    }
    path[++cnt] = u;  // 注意: 逆序打印路径上的顶点
}
*/

// https://leetcode.cn/problems/cracking-the-safe/description/ (欧拉回路经典问题)
//   如何建图？如何转化为求解欧拉回路问题？参考 https://leetcode.cn/problems/cracking-the-safe/solutions/393529/po-jie-bao-xian-xiang-by-leetcode-solution/
//       k^(n-1) 个顶点，k^n 条边，寻找一条欧拉回路

func crackSafe(n int, k int) string {
	visit := make(map[int]bool, 0) // 标记边是否被访问过
	mod := int(math.Pow10(n - 1))
	var ret string

	var dfs func(u int)
	dfs = func(u int) {
		for x := 0; x < k; x++ {
			v := u*10 + x // 掌握构图技巧
			if !visit[v] {
				visit[v] = true
				dfs(v % mod)
				ret += strconv.Itoa(x) // 注意加边顺序
			}
		}
	}
	dfs(0)
	for i := 1; i < n; i++ { // 由于默认从顶点 0 出发寻找欧拉回路，因此要补齐剩余的 0
		ret += "0"
	}
	return ret
}

// https://leetcode.cn/problems/reconstruct-itinerary/description/
//      输出字典序最小的欧拉路径 or 欧拉回路？
//      针对每个顶点的邻边进行排序处理？

type Edge struct {
	// from string
	to  string
	idx int
}

func findItinerary(tickets [][]string) []string {
	m := len(tickets)
	adj := make(map[string][]Edge, 0)
	for i := 0; i < m; i++ {
		from, to := tickets[i][0], tickets[i][1]
		adj[from] = append(adj[from], Edge{to, i})
	}

	// 邻接边按照顶点编号进行排序
	for _, neighs := range adj {
		sort.Slice(neighs, func(i, j int) bool {
			return neighs[i].to < neighs[j].to
		})
	}

	visit := make([]bool, m) // edge mark
	path := make([]string, 0)
	var dfs func(from string)
	dfs = func(from string) {
		for _, e := range adj[from] {
			if !visit[e.idx] {
				visit[e.idx] = true
				dfs(e.to)
			}
		}
		path = append(path, from)
	}
	dfs("JFK")

	// 欧拉路径/回路 --> 逆序反转
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
