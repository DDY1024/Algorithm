package main

// 总结树状数组三类模型
// https://blog.csdn.net/b735098742/article/details/52198579
// 1. 改点求段
// 2. 改段求点
// 3. 改段求段

func lowBit(x int) int {
	return x & (-x)
}

// 1. 一维: 改点求段
func update(idx, n, c int, tree []int) {
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
// 求解前缀和

// 3. 一维: 改段求段
// 此类模型用线段树来处理更佳

// 4. 二维树状数组
// 每一维单独分开计算，高维求解区间元素和时需要用到容斥原理
// 同理类比更高维度的树状数组
func getSumTwo(x, y int, mat [][]int) int {
	ret := 0
	for i := x; i > 0; i -= lowBit(i) {
		for j := y; j > 0; j -= lowBit(j) {
			ret += mat[i][j]
		}
	}
	return ret
}

func updateTwo(x, y, n, m, delta int, mat [][]int) {
	for i := x; i <= n; i += lowBit(i) {
		for j := y; j <= m; j += lowBit(j) {
			mat[i][j] += delta
		}
	}
}

// getSum(x2, y2) - getSum(x2, y1-1) - getSum(x1-1, y2) + getSum(x1-1,y1-1)
