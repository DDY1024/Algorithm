package main

// 参考: https://oi-wiki.org/basic/prefix-sum/
// 参考: https://www.cnblogs.com/COLIN-LIGHTNING/p/8436624.html

// 1. 前缀和
// 一维前缀和: pre[i] = pre[i-1] + arr[i]
// 二维前缀和: pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + arr[i][j]
//
// 树上前缀和 sum(i) 为节点 i 到根节点的权值总和
// 若是点权，则路径 x --> y 权值和为 sum(x) + sum(y) - sum(lca(x,y)) - sum(parent(lca(x,y)))
// 若是边权，则路径 x --> y 权值和为 sum(x) + sum(y) - 2*sum(lca(x,y))

// 2. 差分（一维差分、二维差分）
// 具体参考: https://oi-wiki.org/basic/prefix-sum/
//
// a. 区间差分 --> 差分数组推导区间前缀和
// 	  对于区间 [l,r] 增加 k，则对应的差分操作为 b[l] += k, b[r+1] -= k。全部操作结束后，b 的前缀和即第 i 个元素最终的值 sigma(b[1]+b[2]+...+b[i])
//    询问区间元素和: 差分数组 --> 每个元素 --> 前缀和 --> 区间和
//
// b. 树上差分 --> 比较难，遇到该类题目时再做总结

// 3. 二维差分
// 参考：https://blog.csdn.net/qq_44691917/article/details/109350635
//
// 二维差分数组最终效果: [0,0] --> [x,y] 子矩阵的元素和即为原先 arr[x][y] 的值
//
//
// 3.1 由原二维数组 a[i][j] 构建差分数组 d[i][j]
func build() {
	var a [][]int
	var d [][]int
	var n, m int
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			d[i][j] = a[i][j] - a[i-1][j] - a[i][j-1] + a[i-1][j-1]
			// 前缀子矩阵和
			// sum[i][j] = sum[i][j-1] + sum[i-1][j] - sum[i-1][j-1] + a[i][j]
		}
	}
}

// 3.2 对子矩阵 [i,j] --> [x,y] 所有元素加 val
func add(i, j, x, y, val int, d [][]int) {
	d[i][j] += val   // [i,j] --> [n,m] 子矩阵
	d[x+1][j] -= val // [x+1,j] --> [n,m] 子矩阵
	d[i][y+1] -= val
	d[x+1][y+1] += val
}

// 3.3 对于 [x,y] 元素值的求解，直接计算 [1,1] --> [x,y] 子矩阵元素和即可
