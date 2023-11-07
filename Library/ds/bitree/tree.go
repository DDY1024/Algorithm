package main

import (
	"math"
)

// 树状数组三种模型: https://oi-wiki.org/ds/fenwick/
// 		1. 改点求段
// 		2. 改段求点(差分数组)
// 		3. 改段求段(线段树)

// 一、改点求段
func lowBit(x int) int {
	return x & (-x)
}

// 树状数组性质
// a. O(nlogn) 建树
// b. O(n) 建树：每个节点 i 只需要更新其直接祖先节点 i + lowbit(i)
func build(n int, arr []int, tree []int) {
	for i := 1; i <= n; i++ {
		tree[i] += arr[i]
		j := i + lowBit(i)
		if j <= n {
			tree[j] += tree[i]
		}
	}
}

func add(pos, n, delta int, tree []int) {
	for i := pos; i <= n; i += lowBit(i) {
		tree[i] += delta
	}
}

func sum(pos int, tree []int) int {
	res := 0
	for i := pos; i > 0; i -= lowBit(i) {
		res += tree[i]
	}
	return res
}

// 子矩阵元素和（容斥原理）：sum2d(x2, y2) - sum2d(x2, y1-1) - sum2d(x1-1, y2) + sum2d(x1-1, y1-1)   --> x1 <= x2 && y1 <= y2
func sum2d(x, y int, mat [][]int) int {
	res := 0
	for i := x; i > 0; i -= lowBit(i) {
		for j := y; j > 0; j -= lowBit(j) {
			res += mat[i][j]
		}
	}
	return res
}

func add2d(x, y, n, m, delta int, mat [][]int) {
	for i := x; i <= n; i += lowBit(i) {
		for j := y; j <= m; j += lowBit(j) {
			mat[i][j] += delta
		}
	}
}

// 二、改段求点
//
//  差分数组初始化
// 		原始数组 [a1, a2, a3, a4, a5]
//      差分数组 [a1, a2-a1, a3-a2, a4-a3, a5-a4]
//  差分数组维护
// 		区间变更 [a,b] --> +c
// 			add(a, n, c, tree)
// 			add(b+1, n, -c, tree)
//		元素求值(前缀和)
//          sum(pos, tree)
//
//  二维差分
//      定义：arr[x][y] 表示 (x, y) 与矩阵右下角 (rx, ry) 构成的子矩阵的元素和
//      将子矩阵 (x1, y1) --> (x2, y2) 全部元素 + c
//  		add(x1, y1, c)
// 			add(x1, y2+1, -c)
//      	add(x2+1, y1, -c)
//      	add(x2+1, y2+1, c)
//      求解 (x, y) 的元素值
//         for i := x; i > 0; i -= lowbit(i)
//		       for j := y; j > 0; j -= lowbit(j)

//
// 三、改段求段
// 		1. 一维（区间修改，区间查询）：https://loj.ac/s/1678213
// 		2. 二维（区间修改，区间查询）：https://loj.ac/s/1678699

// 四、第 K 大/小问题
//  1. 朴素算法：二分 + 树状数组求前缀和，复杂度 O(logN*logN)
//  2. 优化参考：https://oi-wiki.org/ds/fenwick/
//     tree[i] 表示出现次数
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

// 五、常见应用
// 1. 前缀区间最值（区间加法性）
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func setMax(pos, x, n int, tree []int) {
	for i := pos; i <= n; i += lowBit(i) {
		tree[i] = maxInt(tree[i], x)
	}
}

func getMax(pos int, tree []int) int {
	ret := tree[pos]
	for i := pos; i > 0; i -= lowBit(i) {
		ret = maxInt(ret, tree[i])
	}
	return ret
}

// 2. 求解序列中 <= x 的元素个数

// 3. 求解逆序对数
// 		逆序对：(i, j)，i < j && arr[i] > arr[j]
// 	  按照顺序遍历，并利用树状数组进行统计，累加求和
