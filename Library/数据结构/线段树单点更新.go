package main

const maxn = 10010

// 以区间求和为例构建对应的线段树模板
// notonlysuccess 线段树模板
var lchild = func(idx int) int { return idx << 1 }
var rchild = func(idx int) int { return (idx << 1) | 1 }

func pushUp(idx int, sum []int) {
	sum[idx] = sum[lchild(idx)] + sum[rchild(idx)]
}

// 根节点 1
// 左孩子: 2*i
// 右孩子: 2*i+1
func buildTree(l, r, idx int, arr []int, sum []int) {
	if l == r {
		sum[idx] = arr[l]
		return
	}
	mid := (l + r) >> 1
	buildTree(l, mid, lchild(idx), arr, sum)
	buildTree(mid+1, r, rchild(idx), arr, sum)
	pushUp(idx, sum)
}

// [l, r]: 查询区间
// [s, t]: 当前线段树节点代表的区间
// idx: 节点索引编号
func query(l, r, s, t, idx int, sum []int) int {
	if l <= s && t <= r {
		return sum[idx]
	}
	mid, ret := (s+t)>>1, 0
	if l <= mid {
		ret += query(l, r, s, mid, lchild(idx), sum)
	}
	if r > mid {
		ret += query(l, r, mid+1, t, rchild(idx), sum)
	}
	return ret
}

func update(pos, delta, s, t, idx int, sum []int) {
	if s == t {
		sum[idx] += delta
		return
	}
	mid := (s + t) >> 1
	if pos <= mid {
		update(pos, delta, s, mid, lchild(idx), sum)
	} else {
		update(pos, delta, mid+1, t, rchild(idx), sum)
	}
	pushUp(idx, sum)
}

/*
// 优雅的 C++ 单点更新模板
#include <cstdio>

// 优雅点 1：参数宏定义
#define lson l , m , rt << 1
#define rson m + 1 , r , rt << 1 | 1

const int maxn = 55555;
int sum[maxn << 2];

// 优雅点 2：PushUp 抽离
void PushUP(int rt) {
    sum[rt] = sum[rt << 1] + sum[rt << 1 | 1];
}

void build(int l, int r, int rt) {
    if (l == r) {
        scanf("%d", &sum[rt]);
        return ;
    }
    // 优雅点 3：能用位运算就用位运算
    int m = (l + r) >> 1;
    build(lson);
    build(rson);
    PushUP(rt);
}

void update(int p, int add, int l, int r, int rt) {
    if (l == r) {
        sum[rt] += add;
        return ;
    }
    int m = (l + r) >> 1;
    if (p <= m)
        update(p , add , lson);
    else
        update(p , add , rson);
    PushUP(rt);
}

int query(int L, int R, int l, int r, int rt) {
    if (L <= l && r <= R) {
        return sum[rt];
    }
    int m = (l + r) >> 1;
    int ret = 0;
    if (L <= m) ret += query(L , R , lson);
    if (R > m)  ret += query(L , R , rson);
    return ret;
}
*/
