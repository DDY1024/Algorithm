package main

import (
	"fmt"
	"math/big"
)

// 扩展的欧几里得算法，用于计算模逆
func extendedEuclidean(a, b *big.Int) *big.Int {
	zero := big.NewInt(0)
	one := big.NewInt(1)

	x0, x1, y0, y1 := new(big.Int), new(big.Int), new(big.Int), new(big.Int)
	x0.Set(zero)
	x1.Set(one)
	y0.Set(one)
	y1.Set(zero)

	for b.Cmp(zero) != 0 {
		q, r := new(big.Int), new(big.Int)
		q.DivMod(a, b, q)
		a, b = b, r

		x0, x1 = x1, x0.Sub(x0, new(big.Int).Mul(q, x1))
		y0, y1 = y1, y0.Sub(y0, new(big.Int).Mul(q, y1))
	}

	if x0.Cmp(zero) == -1 {
		x0.Add(x0, one)
	}

	return x0
}

// 计算 RSA 密钥的私钥 d
func calculatePrivateKey(e, phiN *big.Int) *big.Int {
	return extendedEuclidean(e, phiN)
}

func main() {
	// e, _ := new(big.Int).SetString("65537", 10)
	// phi, _ := new(big.Int).SetString("413575856737258899660", 10)
	// fmt.Println(calculatePrivateKey(e, phi).String())

	// 413575856737258899661
	// 413575856693827317952
	// 413575856693827317952
	// 286493588629981024577

	// 14100036413
	// 29331545297

	// n1, _ := new(big.Int).SetString("14100036412", 10)
	// n2, _ := new(big.Int).SetString("29331545296", 10)

	// e, _ := new(big.Int).SetString("65537", 10)
	// phi := n1.Mul(n1, n2)

	// x := new(big.Int)
	// y := new(big.Int)

	// 高精度扩展欧几里得算法

	// fmt.Println(e.GCD(x, y, e, phi).String())
	// fmt.Println(phi.String())
	// fmt.Println(x.Add(x, phi).String())
	// fmt.Println(n1.Mul(n1, n2).String())

	// fmt.Println(calculatePrivateKey(e, phi))

	// fmt.Println(n1.Mul(n1, n2).String())

	n, _ := new(big.Int).SetString("65537", 10)
	n3, _ := new(big.Int).SetString("286493588629981024577", 10)
	n4, _ := new(big.Int).SetString("413575856693827317952", 10)
	n5 := n3.Mul(n, n3)
	fmt.Println(n5.Mod(n5, n4).String())
}

// 286493588629981024577
// 286493588629981024577
