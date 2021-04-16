package Library

// 参考: https://blog.csdn.net/b735098742/article/details/52198579
// 树状数组常见的三种模型
// 1. 改点求段
// 2. 改段求点
// 3. 改段求段

// 1. 一维：改点求段
func lowBit(x int) int {
	return x & (-x)
}

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
//
// update(a, n, c, tree)
// update(b+1, n, -c, tree)
// getSum(idx, tree) 前缀和表示针对 idx 索引下标元素的整体 delta


// 3. 一维: 改段求段
// 建议此种情况最好采用线段树区间更新的方式来进行处理

// 4. 二维树状数组
// a. 二维树状数组相比于一维树状数组在更新节点的方向是一致的，只是由原先的一维变成了现在的二维操作
// b. 二维（更多维）树状数组在计算区间和时需要用到容斥原理，例如 getSum(x2, y2) - getSum(x2, y1-1) - getSum(x1-1, y2) + getSum(x1-1,y1-1)
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