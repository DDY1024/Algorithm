package main

import "math/big"

type BigInt struct {
	num *big.Int
}

func NewBigInt(x int64) *BigInt {
	return &BigInt{num: big.NewInt(x)}
}

func (a *BigInt) Add(b *BigInt) {
	a.num = a.num.Add(a.num, b.num)
}

func (a *BigInt) Sub(b *BigInt) {
	a.num = a.num.Sub(a.num, b.num)
}

func (a *BigInt) Mul(b *BigInt) {
	a.num = a.num.Mul(a.num, b.num)
}
func (a *BigInt) Div(b *BigInt) {
	a.num = a.num.Div(a.num, b.num)
}

func (a *BigInt) Mod(b *BigInt) {
	a.num = a.num.Rem(a.num, b.num)
	// a.num.Mod()
	// a.num.Mod()
}

func (a *BigInt) Neg() {
	a.num = a.num.Neg(a.num)
}

// 高精度浮点数
type BigFloat struct {
	num *big.Float
}

func NewBigFloat(x float64) *BigFloat {
	return &BigFloat{num: big.NewFloat(x)}
}

// 高精度分数
type BigRat struct {
	num *big.Rat
}

func NewBigRat(a, b int64) *BigRat {
	return &BigRat{
		num: big.NewRat(a, b), // 分子/分母
	}
}

// func main() {
// 	var a big.Int   // 高精度整数
// 	var b big.Float // 高精度浮点数
// 	var c big.Rat   // 高精度分数
// }

// 高精度整数运算
// 1. 低位 --> 高位存储
// 2. 处理进位问题
// 		(a + b + carry) % 10 --> 当前位
//      (a + b + carry) / 10 --> 进位
//
// 3. 高精度减法 a - b
//	    处理好符号位，一般做 大数 - 小数 的减法运算
//      (1 + (-91))
// 4. 高精度乘法
//      11 * 9
//      arr[i+j] += arr[i] * arr[j]
//      处理进位问题
// 5. 高精度除法
//    转化成减法操作，一遍一遍试着减去对应的数q
