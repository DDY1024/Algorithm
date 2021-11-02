package main

const (
	N   = 10    // 矩阵大小
	Mod = 10007 // 模
)

// Matrix 矩阵定义
type Matrix struct {
	mat [N][N]int
}

var E Matrix

func init() {
	E.Unit()
}

func (m *Matrix) Zero() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			m.mat[i][j] = 0
		}
	}
}

func (m *Matrix) Unit() {
	for i := 0; i < N; i++ {
		m.mat[i][i] = 1
		for j := 0; j < N; j++ {
			if i != j {
				m.mat[i][j] = 0
			}
		}
	}
}

func Add(a, b Matrix) Matrix {
	var c Matrix
	c.Zero()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			c.mat[i][j] = (a.mat[i][j] + b.mat[i][j]) % Mod
		}
	}
	return c
}

func Mul(a, b Matrix) Matrix {
	var c Matrix
	c.Zero()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				c.mat[i][j] = (c.mat[i][j] + a.mat[i][k]*b.mat[k][j]) % Mod
			}
		}
	}
	return c
}

// a ^ x
// 二分求解
func Pow(a Matrix, x int) Matrix {
	var c Matrix
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

// a^1 + a^2 + ... + a^x
// 二分求解
func PowSum(a Matrix, x int) Matrix {

	if x == 1 {
		return a
	}

	if x%2 == 0 {
		return Mul(PowSum(a, x/2), Add(E, Pow(a, x/2)))
	}

	return Add(Pow(a, x), PowSum(a, x-1))
}
