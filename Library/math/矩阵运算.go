package main

const (
	Mod = 1e5 + 7
)

type Matrix struct {
	row int
	col int
	mat [][]int
}

var E *Matrix

func init() {
	E = NewMatrix(10, 10) // default 10
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

// 单位矩阵: row = col
func (m *Matrix) Unit() {
	m.Clear()
	for i := 0; i < m.row; i++ {
		m.mat[i][i] = 1
	}
}

// a + b: a.row == b.row, a.col == b.col
func Add(a, b *Matrix) *Matrix {
	c := NewMatrix(a.row, a.col)
	for i := 0; i < a.row; i++ {
		for j := 0; j < a.col; j++ {
			c.mat[i][j] = (a.mat[i][j] + b.mat[i][j]) % Mod
		}
	}
	return c
}

// a * b: a.col == b.row
// c = a * b: c.row = a.row, c.col = b.col
func Mul(a, b *Matrix) *Matrix {
	c := NewMatrix(a.row, b.col)
	for i := 0; i < a.row; i++ {
		for j := 0; j < b.col; j++ {
			for k := 0; k < a.col; k++ {
				c.mat[i][j] += (a.mat[i][k] * b.mat[k][j]) % Mod
			}
		}
	}
	return c
}

// 矩阵快速幂 a^x
func Pow(a *Matrix, x int) *Matrix {
	r := NewMatrix(a.row, a.col)
	for x > 0 {
		if x&1 > 0 {
			r = Mul(r, a)
		}
		a = Mul(a, a)
		x >>= 1
	}
	return r
}

// 矩阵快速幂和: 递归二分
// res = a^1 + a^2 + ... + a^x
// res = a^1 + a^2 + ... + a^(x/2) + a^((x/2)+1) + ... + a^(x-1) + a^x
// A = E*A = A*E
func SumPow(a *Matrix, x int) *Matrix {
	if x == 1 {
		return a
	}

	if x%2 == 0 {
		return Mul(SumPow(a, x>>1), Add(E, Pow(a, x>>1)))
	}

	return Add(Pow(a, x), SumPow(a, x-1))
}
