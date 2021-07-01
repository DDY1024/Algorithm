package main

import "fmt"

// Matrix 矩阵定义
type Matrix struct {
	row int
	col int
	mat [][]int
}

func (m *Matrix) Init(row, col int) {
	m.row = row
	m.col = col
	m.mat = make([][]int, row)
	for i := 0; i < row; i++ {
		m.mat[i] = make([]int, col)
	}
}

// 矩阵单位化 --> 当且仅当 row = col 才有意义
func (m *Matrix) Unit() {
	for i := 0; i < m.row; i++ {
		m.mat[i][i] = 1
	}
}

func Add(a, b Matrix) Matrix {
	var c Matrix
	c.Init(a.row, a.col)
	for i := 0; i < a.row; i++ {
		for j := 0; j < a.col; j++ {
			c.mat[i][j] = a.mat[i][j] + b.mat[i][j]
			// c.mat[i][j] = (a.mat[i][j] + b.mat[i][j]) % Mod
		}
	}
	return c
}

func Mul(a, b Matrix) Matrix {
	var c Matrix
	c.Init(a.row, b.col)
	for i := 0; i < a.row; i++ {
		for j := 0; j < b.col; j++ {
			for k := 0; k < a.col; k++ {
				c.mat[i][j] += a.mat[i][k] * b.mat[k][j]
				// c.mat[i][j] = (c.mat[i][j] + a.mat[i][k]*b.mat[k][j]) % Mod
			}
		}
	}
	return c
}

// a ^ x
func Pow(a Matrix, x int) Matrix {
	var c Matrix
	c.Init(a.row, a.col)
	c.Unit()

	for x > 0 {
		if x&1 > 0 {
			c = Mul(c, a)
		}
		a = Mul(a, a)
		x >>= 1
	}
	return c
}

func numWays(n int, relation [][]int, k int) int {
	var mat Matrix
	mat.Init(n, n)
	rNum := len(relation)
	for i := 0; i < rNum; i++ {
		u, v := relation[i][0], relation[i][1]
		mat.mat[u][v]++
	}

	result := Pow(mat, k)
	return result.mat[0][n-1]
}

// a^1 + a^2 + ... + a^x
// func PowSum(a Matrix, x int) Matrix {

// 	if x == 1 {
// 		return a
// 	}

// 	if x%2 == 0 {
//      其中 E 为单位矩阵：即主对角线元素为1，其它元素均为0
// 		return Mul(PowSum(a, x/2), Add(E, Pow(a, x/2)))
// 	}

// 	return Add(Pow(a, x), PowSum(a, x-1))
// }

func main() {
	fmt.Println("Yes")
}
