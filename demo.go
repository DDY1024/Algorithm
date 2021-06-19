package main

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
)

func power(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 > 0 {
			res = res * a
		}
		b >>= 1
		a = a * a
	}
	return res
}

func calc(a, b int) int {
	ans := 0
	for i := 0; i <= b; i++ {
		ans += power(a, i)
	}
	return ans
}

func smallestGoodBase(n string) string {
	ni, _ := strconv.ParseInt(n, 10, 64)
	maxL := bits.Len(uint(ni)) - 1
	for i := maxL; i > 1; i-- {
		k := int(math.Pow(float64(ni), 1.0/float64(i)))
		if calc(k, i) == int(ni) {
			return strconv.FormatInt(int64(k), 10)
		}

	}
	return strconv.FormatInt(ni-1, 10)
}

// k^(i+1) - 1 / (k-1)  ni
// 2^n-1
// 2^n - 1
//
// (K+1)^m
//
// K^0 + K^1 + ... + K^x = (K^(x+1)-1)/(K-1)
// 二项式公式
// (K+1)^m = C(m, 0) * K^0 + ... + C(m, m) * K^m
//
// 二项式公式决定了其单调递增特性

// k^
// (k+1)^m

func main() {
	// fmt.Println(smallestGoodBase("13"))
	// fmt.Println(power(686286299, 3))
	fmt.Println(smallestGoodBase("470988884881403701"))
}

// 10^9
//
//
//
//
// k^1
// k^1
// a^0 + a^1 + a^2 ... + a^k = b
// a = c^0 + c^1 + ... + c^k
// c^0 + c^2 + ... + c^k
// 10^9
// 10^9 + 1
// 10^18 + 1 1,1,10^9^2
// 1000,0000, 10^14
// 1 + 10^7 + 10^14
// 1 + 10^8 + 10^16
// 1 + 10^9 + 10^18
//
// 2^64
//
//
// 10^18 --> 10^5
//
//
//
//
