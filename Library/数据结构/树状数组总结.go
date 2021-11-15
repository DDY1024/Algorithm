package main

import "math"

// 树状数组三类模型
// 参考资料
// https://blog.csdn.net/b735098742/article/details/52198579
// https://oi-wiki.org/ds/fenwick/
// 1. 改点求段
// 2. 改段求点
// 3. 改段求段

func lowBit(x int) int {
	return x & (-x)
}

// 1. 一维: 改点求段
// O(nlogn) 建树不再讨论
// 存在 O(n) 方式建树：即每个节点 i 只需要更新其"直接祖先节点" i + lowbit(i) 即可（数组下标索引从 1 开始）
func build(n int, arr []int, tree []int) {
	for i := 1; i <= n; i++ {
		tree[i] += arr[i]
		j := i + lowBit(i)
		if j <= n {
			tree[j] += tree[i]
		}
	}
}

func addSum(idx, n, c int, tree []int) {
	for i := idx; i <= n; i += lowBit(i) {
		tree[i] += c
	}
}

func getSum(idx int, tree []int) int {
	ret := 0
	for i := idx; i > 0; i -= lowBit(i) {
		ret += tree[i]
	}
	return ret
}

// 2. 一维: 改段求点
// 套用差分数组性质 + 树状数组
// 例如, 对 [a, b] 区间执行加 c 操作，则对应
// add(a, n, c, tree)
// add(b+1, n, -c, tree) (b+1<=n)
// getSum(idx, tree)：求解前缀和即为某个索引点 idx 对应的全部增量
//
// 对应函数调用操作: 差分数组 + 树状数组
// update(a, n, c, tree)
// update(b+1, n, -c, tree)
// getSum(idx, tree)
// 树状数组恰巧用来求解前缀和

// 3. 一维: 改段求段
// 建议此种情况最好采用线段树区间更新的方式来进行处理
// 求解公式在差分数组的基础上，进一步求解下，参考：https://oi-wiki.org/ds/fenwick/
/*
// c++ 代码如下所示
int t1[MAXN], t2[MAXN], n;

inline int lowbit(int x) { return x & (-x); }

void add(int k, int v) {
  int v1 = k * v;
  while (k <= n) {
    t1[k] += v, t2[k] += v1;
    k += lowbit(k);
  }
}

int getsum(int *t, int k) {
  int ret = 0;
  while (k) {
    ret += t[k];
    k -= lowbit(k);
  }
  return ret;
}

void add1(int l, int r, int v) {
  add(l, v), add(r + 1, -v);  // 将区间加差分为两个前缀加
}

long long getsum1(int l, int r) {
  return (r + 1ll) * getsum(t1, r) - 1ll * l * getsum(t1, l - 1) -
         (getsum(t2, r) - getsum(t2, l - 1));
}
*/

// 4. 二维树状数组: 改点求区间
// 前缀和: 一维 --> 二维
// 子矩阵求和: getSum(x2, y2) - getSum(x2, y1-1) - getSum(x1-1, y2) + getSum(x1-1,y1-1)
func getSumTwo(x, y int, mat [][]int) int {
	ret := 0
	for i := x; i > 0; i -= lowBit(i) {
		for j := y; j > 0; j -= lowBit(j) {
			ret += mat[i][j]
		}
	}
	return ret
}

// 更新操作: 一维 --> 二维
func addSumTwo(x, y, n, m, delta int, mat [][]int) {
	for i := x; i <= n; i += lowBit(i) {
		for j := y; j <= m; j += lowBit(j) {
			mat[i][j] += delta
		}
	}
}

// 5. 二维树状数组: 改区间求点

// 6. 利用树状数组求解第 k 小问题
// 1. 朴素算法：二分查找 + getSum --> O(logN * logN)
// 2. 优化算法，参考：https://oi-wiki.org/ds/fenwick/
// tree[i]: 1 ~ i 范围内的数字出现次数
func kth(tree []int, k, n int) int {
	cnt, ret := 0, 0

	// 二进制分解
	for i := math.Floor(math.Log2(float64(n))); i >= 0; i-- {
		ret += 1 << uint(i)
		if ret >= n || cnt+tree[ret] >= k {
			ret -= 1 << uint(i)
		} else {
			cnt += tree[ret]
		}
	}
	return ret + 1
}

// int kth(int k) {
// 	int cnt = 0, ret = 0;
// 	for (int i = log2(n); ~i; --i) {      // i 与上文 depth 含义相同
// 	  ret += 1 << i;                      // 尝试扩展
// 	  if (ret >= n || cnt + t[ret] >= k)  // 如果扩展失败
// 		ret -= 1 << i;
// 	  else
// 		cnt += t[ret];  // 扩展成功后 要更新之前求和的值
// 	}
// 	return ret + 1;
// }
