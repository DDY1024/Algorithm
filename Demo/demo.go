package main

import (
	"fmt"
)

// func main() {
// 	// arr := []int{1, 2, 3}
// 	// slices.Reverse(arr)
// 	// fmt.Println(arr)

// 	switch 1 {
// 	case 1:
// 		fmt.Println(1)
// 	case 2:
// 		fmt.Println(2)
// 	}
// 	// slices.Clone[]()
// 	slices.c
// }

// 根节点、左子树、右子树
// 左子树、根节点、右子树
// 左子树、右子树、根节点

func main() {
	// return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
	// fmt.Println(C(99, 99))
	// return C(m+n-2, n-1)
}

func C(m, n int) int {
	n = min(n, m-n)
	fz, fm := 1, 1
	for i := 1; i <= n; i++ {
		fz *= m - i + 1
		fm *= i
		d := gcd(fz, fm)
		fz /= d
		fm /= d
		fmt.Println(fz, fm)
	}
	return fz / fm
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
