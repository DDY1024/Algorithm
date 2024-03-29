package main

import (
	"math/big"
)

// big.Int
// big.Float
// big.Rat

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

func (a *BigInt) Neg() {
	a.num = a.num.Neg(a.num)
}

type BigFloat struct {
	num *big.Float
}

func NewBigFloat(x float64) *BigFloat {
	return &BigFloat{num: big.NewFloat(x)}
}

func (a *BigFloat) Add(b *BigFloat) {
	a.num = a.num.Add(a.num, b.num)
}

type BigRat struct {
	num *big.Rat
}

func NewBigRat(a, b int64) *BigRat {
	return &BigRat{
		num: big.NewRat(a, b),
	}
}
