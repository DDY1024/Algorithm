package main

import "fmt"

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDifference(nums []int, queries [][]int) []int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 101)
	}
	dp[0][nums[0]]++
	for i := 1; i < n; i++ {
		dp[i][nums[i]]++
		for j := 1; j <= 100; j++ {
			dp[i][j] += dp[i-1][j]
		}
	}

	m := len(queries)
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		l, r := queries[i][0], queries[i][1]
		preV, minV := -1000, 1000
		for j := 1; j <= 100; j++ {
			cnt := dp[r][j]
			if l-1 >= 0 {
				cnt -= dp[l-1][j]
			}
			if cnt > 0 && j != preV {
				minV = minInt(minV, j-preV)
				preV = j
			}
		}
		if minV >= 1000 {
			ans[i] = -1
		} else {
			ans[i] = minV
		}
	}
	return ans
}

func doPrint() bool {
	fmt.Println("hello, world!")
	return true
}

func main() {
	fmt.Println(false && doPrint())
	fmt.Println(true && doPrint())
	fmt.Println(true || doPrint())
	fmt.Println(false || doPrint())
}

// "||"、"&&" 均是短路运算符
