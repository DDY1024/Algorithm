package main

// 1. n * n 矩阵顺时针旋转 90 度
// https://leetcode.cn/problems/rotate-image/solution/xuan-zhuan-tu-xiang-by-leetcode-solution-vu3m/
//
//	  坐标变换 (row, col) --> (col, n - row - 1)
//
//   由上述坐标变换可以发现，每四个点会构成一个循环
//   针对 n 为 奇数 和 偶数 两种情况，我们选择左上角的部分点进行循环交换即可
//
//
//   另外一种变换思路
//   a. 按照水平轴上下翻转
//   b. 按照主对角线进行翻转

func rotate1(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			// 四个坐标点循环交换即可
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

func rotate2(matrix [][]int) {
	n := len(matrix)

	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}



