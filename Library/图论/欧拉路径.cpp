/*
* 无向图判定: 连通、最多两个奇度顶点
* 有向图判定: 底图连通(单向连通)、两个奇度顶点(入度-出度=1,出度-入度=1)
*/

#include <set>
#include <stack>
#include <queue>
#include <cmath>
#include <cstdio>
#include <string>
#include <cstring>
#include <iostream>
#include <algorithm>
using namespace std;
//#pragma comment(linker, "/STACK:1024000000,1024000000")
#define FIN             freopen("input.txt","r",stdin)
#define FOUT            freopen("output.txt","w",stdout)
#define fst             first
#define snd             second
//typedef __int64 LL;
typedef long long LL;
typedef pair<int, int> PII;
const int MAXN = 1e4 + 5;
const int MAXE = 5e4 + 5;
int T, N, M, K;
struct Edge {
    int v, next;
    Edge() {}
    Edge (int v, int next) : v (v), next (next) {}
} edges[MAXE << 1];
int head[MAXN], tot;
int path[MAXN], cnt;
int deg[MAXN];
bool used[MAXE];
void init() {
    tot = 0;
    cnt = 0;
    memset (head, -1, sizeof (head) );
    memset (deg, 0, sizeof (deg) );
    memset (used, false, sizeof (used) );
}
void add_edge (int u, int v) {
    edges[tot] = Edge (v, head[u]);
    head[u] = tot++;
}
void dfs (int u) {
    // int v;
    // 删边遍历
    for (int i = head[u]; ~i; i = head[u]) {
        // v = edges[i].v;
        // head[u] = edges[i].next;
        if (!used[i]) {
            used[i] = used[i ^ 1] = true;  // 判边重而不是判点重
            dfs (edges[i].v);
        }
    }
    path[++ cnt] = u;  // 注意: 打印欧拉路径时应该逆序打印
}

/*
* 邻接矩阵表示法
* 顶点数 n
* 无向图欧拉回路求解
void euler(int u) {
    for(int v = 0; v < n; v++) {
        if (G[u][v] && !vis[u][v]) {
            vis[u][v] = vis[v][u] = 1;
            euler(v);
            printf("%d %d\n", u, v)  // 打印顺序应该逆序，实际使用时将边 (u, v) 压入栈中
        }
    }
}
* 有向图欧拉回路求解，则将上述代码改成 vis[u][v] = 1 即可
*/

int main() {
//#ifdef __WONZY_LOCAL__
//    FIN;
//#endif // __WONZY_LOCAL__
    int u, v, x;
    while (~scanf ("%d %d", &N, &M)) {
        init();
        for (int i = 1; i <= M; i++) {
            scanf ("%d %d", &u, &v);
            add_edge (u, v);
            add_edge (v, u);
            deg[u] ++;
            deg[v] ++;
        }
        int s = -1, t = -1, num = 0;
        for (int i = 1; i <= N; i++) {
            if (deg[i] & 1) {
                num ++;
                if (num == 1) s = i;
                else t = i;
            }
        }
        if (num == 0) s = 1;
        dfs (s);
        for (int i = 1; i <= M + 1; i++) {
            printf ("%d%c", path[i], i <= M ? ' ' : '\n');
        }
    }
    return 0;
}

