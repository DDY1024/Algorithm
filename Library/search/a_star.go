package main

// 参考资料: https://oi-wiki.org/search/astar/#fn:note1
// 利用 A-star 求解最优路径（次优路径）的难点在于如何设计有效的估价函数
// f(x) = g(x) + h(x)
// g(x): 从起始点到当前节点 x 已经走过的路径长度
// h(x): 当前节点 x 到终点的距离估价函数
// h(x) <= h'(x) 时，A* 算法是可以求解出最优路径的，其中 h'(x) 为 x 节点到终点实际意义上的最短距离

// 1. 八数码 A* 算法
// 估价函数 h(s): 当前不在对应位置上的数字个数
// 八数码状态压缩表示: 康托展开、逆康托展开

// 2. 从顶点 s 到顶点 t 的第 K 短路径
// 估价函数 h(u): 当前节点 u 到终点 t 的最短路径
// 剪枝方法: 由于求解的是第 k 短路径，因此这 k 条路径中每个顶点经过的次数最多不超过 k 次，可用于搜索时的剪枝
/*
#include <algorithm>
#include <cstdio>
#include <cstring>
#include <queue>
using namespace std;
const int maxn = 5010;
const int maxm = 400010;
const double inf = 2e9;
int n, m, k, u, v, cur, h[maxn], nxt[maxm], p[maxm], cnt[maxn], ans;
int cur1, h1[maxn], nxt1[maxm], p1[maxm];
double e, ww, w[maxm], f[maxn];
double w1[maxm];
bool tf[maxn];
void add_edge(int x, int y, double z) {  //正向建图函数
  cur++;
  nxt[cur] = h[x];
  h[x] = cur;
  p[cur] = y;
  w[cur] = z;
}
void add_edge1(int x, int y, double z) {  //反向建图函数
  cur1++;
  nxt1[cur1] = h1[x];
  h1[x] = cur1;
  p1[cur1] = y;
  w1[cur1] = z;
}
struct node {  //使用A*时所需的结构体
  int x;
  double v;
  bool operator<(node a) const { return v + f[x] > a.v + f[a.x]; }
};
priority_queue<node> q;
struct node2 {  //计算t到所有结点最短路时所需的结构体
  int x;
  double v;
  bool operator<(node2 a) const { return v > a.v; }
} x;
priority_queue<node2> Q;
int main() {
  scanf("%d%d%lf", &n, &m, &e);
  while (m--) {
    scanf("%d%d%lf", &u, &v, &ww);
    add_edge(u, v, ww);   //正向建图
    add_edge1(v, u, ww);  //反向建图
  }
  for (int i = 1; i < n; i++) f[i] = inf;
  // 创建逆向图求解从终点 t 到其它顶点的最短路径作为估价函数 h(i)
  Q.push({n, 0});
  while (!Q.empty()) {  //计算t到所有结点的最短路
    x = Q.top();
    Q.pop();
    if (tf[x.x]) continue;
    tf[x.x] = true;
    f[x.x] = x.v;
    for (int j = h1[x.x]; j; j = nxt1[j]) Q.push({p1[j], x.v + w1[j]});
  }
  k = (int)e / f[1];
  q.push({1, 0});
  // f(i) = g(i) + h(i) 按照优先级顺序出队
  while (!q.empty()) {  //使用A*算法
    node x = q.top();
    q.pop();
    cnt[x.x]++;  // 记录某个顶点出现的次数，我们控制每个顶点的到达次数最多不超过 k 次
    if (x.x == n) {
      e -= x.v;
      if (e < 0) { // 猜测这个 e 是总共允许的最大花费
        printf("%d\n", ans);
        return 0;
      }
      ans++;
    }
    for (int j = h[x.x]; j; j = nxt[j])
	  // 超过 k 次进行剪枝，不再遍历
	  // cnt[p[j]] < k 吧？已经遍历 k 次，剩下没必要再进行遍历了
      if (cnt[p[j]] <= k && x.v + w[j] <= e) q.push({p[j], x.v + w[j]});
  }
  printf("%d\n", ans);
  return 0;
}
*/
