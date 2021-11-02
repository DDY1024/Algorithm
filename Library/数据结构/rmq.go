package main

import (
	"fmt"
	"math"
)

//
// 换一种角度考虑，其实其应该理解成 ”倍增 DP“
// 区间范围最值，实际上如果两个区间满足加法性质，均可以采用 rmq 思想来求解
//

var ls = func(x int) int { return 1 << uint(x) }

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calcExp(n int) int {
	return int(math.Floor(math.Log2(float64(n))))
}

func getMax(l, r int, dp [][]int) int {
	exp := calcExp(r - l + 1)
	return maxInt(dp[l][exp], dp[r-ls(exp)+1][exp])
}

func getMin(l, r int, dp [][]int) int {
	exp := calcExp(r - l + 1)
	return minInt(dp[l][exp], dp[r-ls(exp)+1][exp])
}

func rmqInit(n int, arr []int) [][]int {
	exp := calcExp(n)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, exp+1)
		dp[i][0] = arr[i]
	}

	for l := 1; l <= exp; l++ {
		for i := 0; i+ls(l)-1 < n; i++ {
			dp[i][l] = maxInt(dp[i][l-1], dp[i+ls(l-1)][l-1])
			// dp[i][l] = minInt(dp[i][l-1], dp[i+ls(l-1)][l-1])
		}
	}
	return dp
}

func main() {
	arr := []int{1, 2, 3, 4}
	dp := rmqInit(4, arr)
	fmt.Println(getMax(0, 3, dp))
	fmt.Println(getMax(1, 2, dp))
	fmt.Println(getMax(2, 2, dp))
	fmt.Println(getMax(0, 1, dp))
	fmt.Println(getMax(0, 0, dp))
	//fmt.Println(math.Floor(math.Log2(4.00,000001)))
	//fmt.Println(math.Floor(math.Log2(3.0)))
}
