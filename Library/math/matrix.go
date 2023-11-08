package main

const (
	mod = 1e9 + 7
)

type Matrix struct {
	row int
	col int
	mat [][]int
}

var E *Matrix

func init() {
	E = NewMatrix(10, 10)
	E.Unit()
}

func NewMatrix(row, col int) *Matrix {
	m := &Matrix{row: row, col: col}
	m.mat = make([][]int, row)
	for i := 0; i < row; i++ {
		m.mat[i] = make([]int, col)
	}
	return m
}

func (m *Matrix) Clear() {
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.col; j++ {
			m.mat[i][j] = 0
		}
	}
}

func (m *Matrix) Unit() {
	m.Clear()
	for i := 0; i < m.row; i++ {
		m.mat[i][i] = 1
	}
}

func add(a, b *Matrix) *Matrix {
	c := NewMatrix(a.row, a.col)
	for i := 0; i < c.row; i++ {
		for j := 0; j < c.col; j++ {
			c.mat[i][j] = (a.mat[i][j] + b.mat[i][j]) % mod
		}
	}
	return c
}

func mul(a, b *Matrix) *Matrix {
	c := NewMatrix(a.row, b.col)
	for i := 0; i < c.row; i++ {
		for j := 0; j < c.col; j++ {
			for k := 0; k < a.col; k++ {
				c.mat[i][j] += a.mat[i][k] * b.mat[k][j]
				c.mat[i][j] %= mod
			}
		}
	}
	return c
}

// 矩阵快速幂 a^x
func pow(a *Matrix, x int) *Matrix {
	r := NewMatrix(a.row, a.col)
	for x > 0 {
		if x&1 > 0 {
			r = mul(r, a)
		}
		a = mul(a, a)
		x >>= 1
	}
	return r
}

// 矩阵快速幂和
// res = a^1 + a^2 + ... + a^x
// res = a^1 + a^2 + ... + a^(x/2) + a^((x/2)+1) + ... + a^(x-1) + a^x
// A = E*A = A*E
func powSum(a *Matrix, x int) *Matrix {
	if x == 1 {
		return a
	}

	if x&1 == 0 {
		return mul(powSum(a, x/2), add(E, pow(a, x/2)))
	}
	return add(pow(a, x), powSum(a, x-1))
}
