package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// 通常存在两类树的 dfs 序列
// 1. 在节点进和出都加进序列
// 2. 欧拉序: 只要到达每个顶点都将其加进序列中
// 参考：https://www.jianshu.com/p/55037ae618ce

type Edge struct {
	v, next int
}

var edges []Edge
var head []int
var enum int
var n, m int

func initGraph() {
	head = make([]int, n)
	for i := 0; i < n; i++ {
		head[i] = -1
	}
	enum = 0
	edges = make([]Edge, m)
}

// 无向图 <u, v> 和 <v, u>
func addEdge(u, v int) {
	edges[enum].v = v
	edges[enum].next = head[u]
	head[u] = enum
	enum++

	edges[enum].v = u
	edges[enum].next = head[v]
	head[v] = enum
	enum++
}

// 在顶点进入和出来均打点标记，能够解决下面的问题
// 1. 根节点至当前节点的路径权值和: 进入的点采用正权值，出来的点采用负权值，根节点 --> 目标节点，转化成前缀和求解即可。
// 2. 子树权值和: in[u] ---> out[u]，节点权值总和 / 2，由于进出每个节点被计算了两次。
// 正常的 dfs 序最终会输出 2*n 个顶点
var inN []int
var outN []int
var vn []int
var dfn int

func dfsOne(u, p int) {
	dfn++
	inN[u] = dfn  // 进入点
	vn[dfn] = u
	for i := head[u]; i != -1; i = edges[i].next {
		v := edges[i].v
		if v == p {
			continue
		}
		dfsOne(v, u)
	}
	dfn++
	outN[u] = dfn  // 出去点
	vn[dfn] = u
	return
}

// 欧拉序输出
// 1. 子树权值和: 从欧拉序记录的方法不难看出，所有子树节点均存在于 first[u] --> last[u] 之间，不重复计算节点权值和即可
// 2. LCA 求解: first[u] ... fist[v] 区间之间深度最小的顶点即为最近公共祖先
// 欧拉序列最终会输出 2*n - 1 个顶点
var first []int	// 记录该顶点在欧拉序中第一次出现的位置
var rmq []int // 记录每个顶点的深度
func dfsTwo(u, p, depth int) {
	dfn++
	vn[dfn] = u	// 进入点
	rmq[dfn] = depth
	first[u] = dfn
	for i := head[u]; i != -1; i = edges[i].next {
		v := edges[i].v
		if v == p {
			continue
		}
		dfsTwo(v, u, depth+1)
		dfn++
		vn[dfn] = u  // 回溯点
		rmq[dfn] = depth
	}
	return
}


// 1. 欧拉序 + rmq 在线求解 lca
var dp [][]int

func initRMQ() {
	dp = make([][]int, 2*n)
	for i := 0; i < 2*n; i++ {
		dp[i] = make([]int, 20)
	}
	for i := 1; i < 2*n; i++ {
		dp[i][0] = i
	}
	for l := 1; 1<<uint(l) < 2*n; l++ {
		for i := 1; i + (1<<uint(l)) < 2*n; i++ {
			id1 := dp[i][l-1]
			id2 := dp[i+(1<<uint(l-1))][l-1]
			if rmq[id1] <= rmq[id2] {
				dp[i][l] = id1
			} else {
				dp[i][l] = id2
			}
		}
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func queryRMQ(u, v int) int {
	l := minInt(first[u], first[v])
	r := maxInt(first[u], first[v])
	st := int(math.Log2(float64(r-l+1)))
	// fmt.Println(l, r, st)
	id1 := dp[l][st]
	id2 := dp[r-(1<<uint(st))+1][st]
	// fmt.Println("Query:", u, v, id1, id2)
	if rmq[id1] <= rmq[id2] {
		return vn[id1]
	}
	return vn[id2]
}


// 2. tarjan 算法求解 lca
/*
// 并查集 + dfs
Tarjan(u)//marge和find为并查集合并函数和查找函数
{
    for each(u,v)    //访问所有u子节点v
    {
        Tarjan(v);        //继续往下遍历
        marge(u,v);    //合并v到u上
        标记v被访问过;
    }
    for each(u,e)    //访问所有和u有询问关系的e
    {
        如果e被访问过;
        u,e的最近公共祖先为find(e);
    }
}
*/

// 并查集维护父节点的关系
var parent []int

func initSet() {
	parent = make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
}

func find(u int) int {
	if parent[u] < 0 {
		return u
	}
	parent[u] = find(parent[u])
	return parent[u]
}

// 注意 u, v 间的顺序
// 此处的 merge 顺序也是很重要的
func merge(u, v int) {
	r1, r2 := find(u), find(v)
	if r1 != r2 {  // 注意 merge 的方向
		parent[r1] = r2
	}
}

type Query struct {
	v, idx int   // 存储查询结果
}

var vis []bool
var ances []int
var query [][]Query

func TarjanLCA(u int) {
	ances[u] = u
	vis[u] = true
	for i := head[u]; i != -1; i = edges[i].next {
		v := edges[i].v
		if vis[v] {
			continue
		}
		TarjanLCA(v)
		merge(u, v)
		ances[find(u)] = u
	}
	for i := range query[u] {
		v := query[u][i].v
		if vis[v] {  // 顶点被访问过
			fmt.Println("LCA: ", u, v, ances[find(v)])
		}
	}
}

var fa [][]int  // fa[i][j] 节点 i 的第 2^j 个祖先节点
var vdepth []int // 节点 i 的深度

func BFS(root, n int) {
	fa = make([][]int, n)
	vdepth = make([]int, n)
	vis = make([]bool, n)
	for i := 0; i < n; i++ {
		fa[i] = make([]int, 10)
	}
	que := make([]int, 0, n)
	que = append(que, root)
	fa[root][0] = root
	vdepth[root] = 0
	vis[root] = true
	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		for i := 1; i < 10; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := head[u]; i != -1; i = edges[i].next {
			v := edges[i].v
			if vis[v] {
				continue
			}
			vis[v] = true  // 此处其实可以直接用父节点做判断
			vdepth[v] = vdepth[u] + 1
			fa[v][0] = u
			que = append(que, v)
		}
	}
	return
}

// 采用倍增法求解 lca
// 自底向上一步步逼近
func BLCA(u, v int) int {
	if vdepth[u] > vdepth[v] {
		vdepth[u], vdepth[v] = vdepth[v], vdepth[u]
	}
	for dh, i := vdepth[v]-vdepth[u], 0; dh > 0; dh, i = dh/2, i+1 {
		if dh&1 > 0 {
			v = fa[v][i]
		}
	}
	if u == v {
		return u
	}
	// fa[u][i] 到顶了，则 fa[u][i+1] 也是到顶的
	for i := 9; i >= 0; i-- {  // 此处求解的方法比较巧妙
		if fa[u][i] == fa[v][i] {
			continue
		}
		u, v = fa[u][i], fa[v][i]
	}
	return fa[u][0]
}


func main() {
	n, m = 8, 20
	initGraph()
	addEdge(0, 1)
	addEdge(0, 5)
	addEdge(1, 2)
	addEdge(1, 3)
	addEdge(1, 4)
	addEdge(5, 6)
	addEdge(5, 7)
	inN = make([]int, 8)
	outN = make([]int, 8)
	vn = make([]int, 100)
	rmq = make([]int, 100)
	first = make([]int, 100)
	dfn = 0
	dfsOne(0, -1)
	fmt.Println(vn[:dfn+1])
	fmt.Println(inN)
	fmt.Println(outN)
	// [0 0 5 7 7 6 6 5 1 4 4 3 3 2 2 1 0]
	// [1 8 13 11 9 2 5 3]
	// [16 15 14 12 10 7 6 4]

	dfn = 0
	dfsTwo(0, -1, 0)
	fmt.Println(vn[:dfn+1])
	// [0 0 5 7 5 6 5 0 1 4 1 3 1 2 1 0]
	fmt.Println(rmq[:dfn+1])
	fmt.Println(first[:n])
	initRMQ()
	// fmt.Println(len(动态规划), len(动态规划[0]))
	fmt.Println(0, 5, queryRMQ(0, 5))
	fmt.Println(6, 7, queryRMQ(6, 7))
	fmt.Println(0, 7, queryRMQ(0, 7))
	fmt.Println(3, 7, queryRMQ(3, 7))

	// tarjan 离线算法求解 lca
	initSet()
	vis = make([]bool, n)
	ances = make([]int, n)
	query = make([][]Query, n)
	query[0] = append(query[0], Query{5, -1})
	query[6] = append(query[6], Query{7, -1})
	query[0] = append(query[0], Query{7, -1})
	query[3] = append(query[3], Query{7, -1})
	TarjanLCA(0)  // tarjan 算法求 lca 只需要过一遍即可，还是很高效的

	// 倍增法求解 lca
	BFS(0, 8)
	fmt.Println(0, 5, BLCA(0, 5))
	fmt.Println(6, 7, BLCA(6, 7))
	fmt.Println(0, 7, BLCA(0, 7))
	fmt.Println(3, 7, BLCA(3, 7))

	xx := map[int][]string{}
	xx[1] = []string{"a", "b"}
	ssx, _ := json.Marshal(xx)
	fmt.Println(string(ssx))
	var yy map[int][]string
	_ = json.Unmarshal(ssx, &yy)
	fmt.Println(yy)
}

