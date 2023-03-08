package main

import (
	"math"
)

// 树状数组三种模型
// https://oi-wiki.org/ds/fenwick/
// 		1. 改点求段
// 		2. 改段求点(差分数组)
// 		3. 改段求段(线段树更优)

func lowBit(x int) int {
	return x & (-x)
}

// 一、改点求段
// 参考：https://loj.ac/s/1678182
// 1.1 一维树状数组
// a. O(nlogn) 建树
// b. O(n) 建树：即每个节点 i 只需要更新其"直接祖先节点" i + lowbit(i) 即可
func build(n int, arr []int, tree []int) {
	for i := 1; i <= n; i++ {
		tree[i] += arr[i]
		j := i + lowBit(i)
		if j <= n {
			tree[j] += tree[i]
		}
	}
}

func add(idx, n, c int, tree []int) {
	for i := idx; i <= n; i += lowBit(i) {
		tree[i] += c
	}
}

func sum(idx int, tree []int) int {
	ret := 0
	for i := idx; i > 0; i -= lowBit(i) {
		ret += tree[i]
	}
	return ret
}

// 1.2 二维树状数组
// 参考：https://loj.ac/s/1677181
//  Add 方法
//	for i := x; i <= n; i += lowbit(i)
//     for j := y; j <= m; j += lowbit(j)
//
//  Sum 方法
//  for i := x; i > 0; i -= lowbit(i)
//	   for j := y; j > 0; j -= lowbit(j)

// 子矩阵求和（容斥原理）：getSum(x2, y2) - getSum(x2, y1-1) - getSum(x1-1, y2) + getSum(x1-1,y1-1)
func getSumTwo(x, y int, mat [][]int) int {
	ret := 0
	for i := x; i > 0; i -= lowBit(i) {
		for j := y; j > 0; j -= lowBit(j) {
			ret += mat[i][j]
		}
	}
	return ret
}

func addSumTwo(x, y, n, m, delta int, mat [][]int) {
	for i := x; i <= n; i += lowBit(i) {
		for j := y; j <= m; j += lowBit(j) {
			mat[i][j] += delta
		}
	}
}

// 二、改段求点
// 1.1 一维
// 参考实现：https://loj.ac/s/1678198
// 一维数组初始化
// 		a1, a2-a1, a3-a2, ..., a(n-1) - a(n-2)
// 利用一维差分数组的性质
// 		1. 区间修改 [a,b] --> +c
// 			add(a,n,c,tree)
// 			add(b+1,n,-c,tree)
//		2. 单点求和
// 			getSum(idx,tree)
// 1.2 二维
// 二维差分数组：
//    1. 区间修改 [x1, y1] [x2, y2] 子矩形 + c
//  	add(x1, y1, c)
// 		add(x1, y2+1, -c)
//      add(x2+1, y1, -c)
//      add(x2+1, y2+1, c)
//    2. 求和
//      for i := x; i > 0; i -= lowbit(i)
//			for j := y; j > 0; j -= lowbit(j)
//

// 三、改段求段
// 一维（区间修改，区间查询）：https://loj.ac/s/1678213
// 二维（区间修改，区间查询）：https://loj.ac/s/1678699

// 四、第 K 大/小问题
//  1. 朴素算法：二分 + getSum，复杂度 O(logN*logN)
//  2. 优化参考：https://oi-wiki.org/ds/fenwick/
//     tree[i] 表示出现次数
//
// 优化算法
func kth(tree []int, k, n int) int {
	cnt, ret := 0, 0
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

// 五、树状数组应用
// 6.1 维护区间最值
// 利用区间合并性，适应于 [1,x] 形式的区间
// update(idx, x, n)
//	  for i := idx; i <= n; i += lowbit(i)
//	      tree[i] = max(tree[i], x)
//
// getmax(idx)
//	  for i := idx; i > 0; i -= lowbit(i)
//        ret = max(ret, tree[i])
//

// 6.2 求解序列中 <= x 的元素个数
//
// update(x, n)
// 	  for i := x; i <= n; i += lowbit(i)
//        tree[i]++
//
// getcnt(x)
//    for i := x; i > 0; i -= lowbit(i)
//        ret += tree[i]

// 6.3 求解逆序对数
// 逆序对：(i, j)，i < j && arr[i] > arr[j]
// 只需从左往右遍历，寻找左侧比当前元素大的个数，累加即可
