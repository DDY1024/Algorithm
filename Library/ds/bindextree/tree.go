package main

import (
	"math"
)

// 树状数组三类模型
// https://blog.csdn.net/b735098742/article/details/52198579
// https://oi-wiki.org/ds/fenwick/
// https://zhuanlan.zhihu.com/p/93795692
//
// 1. 改点求段
// 2. 改段求点(差分数组)
// 3. 改段求段(还是线段树吧)

func lowBit(x int) int {
	return x & (-x)
}

// 一、改点求段模型
// 1.1 一维
// 	a. O(nlogn) 建树
//  b. O(n) 建树：即每个节点 i 只需要更新其"直接祖先节点" i + lowbit(i) 即可
// 		O(n) 建树，只需要更新自己的直接父亲节点即可
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

// 1.2 二维
//
//  Add 操作
//	for i := x; i <= n; i += lowbit(i)
//     for j := y; j <= m; j += lowbit(j)
//
//  Sum 操作
//  for i := x; i > 0; i -= lowbit(i)
//	   for j := y; j > 0; j -= lowbit(j)

// 二、改段求点模型
// 	套用差分数组的性质，针对 [l,r] 区间增加 c，相当于 d[l] += c, d[r+1] -= c，其中 d 数组为差分数组;
//  在差分数组形式下，如果要求解 arr[i] 具体的值，则直接求解差分数组的前缀和即可
//
// 例如，针对 [a,b] 区间内所有元素 +c
// add(a,n,c,tree)
// add(b+1,n,-c,tree)
//
// getSum(idx,tree)

// 三、改段求段模型
// 建议直接使用线段树，不要用树状数组来处理了

// 四、多维模型
// 1.1 子矩阵求和
//	 getSum(x2, y2) - getSum(x2, y1-1) - getSum(x1-1, y2) + getSum(x1-1,y1-1)  --> 容斥原理
//
func getSumTwo(x, y int, mat [][]int) int {
	ret := 0
	for i := x; i > 0; i -= lowBit(i) {
		for j := y; j > 0; j -= lowBit(j) {
			ret += mat[i][j]
		}
	}
	return ret
}

// 更新操作
func addSumTwo(x, y, n, m, delta int, mat [][]int) {
	for i := x; i <= n; i += lowBit(i) {
		for j := y; j <= m; j += lowBit(j) {
			mat[i][j] += delta
		}
	}
}

// 1.2 二维改段求点
// 一维差分&&二维差分: https://blog.csdn.net/qq_44691917/article/details/109350635

// 五、利用树状数组求解第 k 小问题
// 1. 朴素算法：二分 + getSum，复杂度 O(logN*logN)
// 2. 优化算法，参考：https://oi-wiki.org/ds/fenwick/
// 		tree[i]: 1 ~ i 范围内的数字出现次数
// 优化后的求第 kth 算法
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

// 六、树状数组经典应用
// 6.1 维护区间最值（利用区间合并性）
//
// update(idx, x, n)
//	  for i := idx; i <= n; i += lowbit(i)
//	      tree[i] = max(tree[i], x)
//
// getmax(idx)
//	  for i := idx; i > 0; i -= lowbit(i)
//        ret = max(ret, tree[i])

// 6.2 求解序列中 <= x 的元素个数（下表表示元素值）
//
// update(x, n)
// 	  for i := x; i <= n; i += lowbit(i)
//        tree[i]++
//
// getcnt(x)
//    for i := x; i > 0; i -= lowbit(i)
//        ret += tree[i]

// 6.3 求解逆序对数
// 逆序对定义：i < j 且 arr[i] > arr[j]
// 只需从左往右遍历，寻找左侧比当前元素大的元素个数，叠加即可
