package main

import (
	"fmt"
	"math/big"
)

// 参考资料：https://oi-wiki.org/math/bignum/
// 1. 高精度整数
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
}

// 取反操作
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

func (a *BigFloat) Add(b *BigFloat) {
	a.num = a.num.Add(a.num, b.num)
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

// 高精度整数运算模拟算法
// 	1. 低位(0) --> 高位(n) 存储
//  2. 进位、当前位
// 		cur = (a + b + carry) % 10
//      carry = (a + b + carry) / 10
//  3. 高精度减法（大数-小数）
//		cmp(a, b) 比较两数大小
//  4. 高精度乘法
//  	c[i+j] += a[i] * b[j]
//      低位 --> 高位处理 c 的进位问题
//  5. 高精度除法
//      转化成减法操作，不断试减

func main() {
	var a *big.Int   // big.Int
	var b *big.Float // big.Float
	var c *big.Rat   // big.Rat
	var d *big.Rat   // 假分数、真分数均可
	a = big.NewInt(1)
	b = big.NewFloat(0.1)
	c = big.NewRat(1, 9)
	d = big.NewRat(10, 9)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
