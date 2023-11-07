package main

import "fmt"

const (
	inf = 0x3f3f3f3f3f3f3f3f
)

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

//  1. 最大连续子段和
//     dp[i] 表示以 arr[i] 为结尾的最大连续子段和
func maxSubSum(arr []int) int {
	n := len(arr)
	dp := make([]int, n)
	dp[0] = arr[0]
	ret := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = maxInt(arr[i], dp[i-1]+arr[i])
		ret = maxInt(ret, dp[i])
	}
	return ret
}

func maxSubSum2(arr []int) int {
	ret, n, tmp := 0, len(arr), 0
	for i := 0; i < n; i++ {
		if tmp < 0 {
			tmp = 0
		}
		tmp += arr[i]
		ret = maxInt(ret, tmp)
	}
	return ret
}

// 2. 最大 M 子段和
//   	dp[i][j] : 前 i 个数，分成 j 段且 arr[i] 结尾
//
//	  状态转移
// 		  dp[i][j] = dp[i-1][j] + arr[i] --> arr[i] 与 arr[i-1] 同段
// 		  dp[i][j] = max{ dp[t][j-1] + arr[i] } --> arr[i] 单独一段
//
// 	  最终结果
// 	      result = max{ dp[i][m] }, i >= m

func maxMSubSum(arr []int, m int) int {
	n := len(arr)

	// 滚动数组
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
	}

	// 正常情况下，针对段数为 0 的情况，结果为 0
	// 由于采用滚动数组进行压缩，此处 dp[0][i] 的初始赋值为 0，会对后续结果计算产生影响，因此初始赋值需要设置为 -inf
	// [-1, -2, -3, -1], 3 --> -4
	for i := 1; i <= n; i++ {
		dp[0][i] = -inf
	}

	for i := 1; i <= m; i++ {
		idx := i & 1
		tmp := dp[idx^1][i-1]
		for j := i; j <= n; j++ {
			// 1. arr[j] 与 arr[j-1] 为同一段
			dp[idx][j] = dp[idx][j-1] + arr[j-1] // arr 下标从 0 开始

			// 2. arr[j] 单独一段
			dp[idx][j] = maxInt(dp[idx][j], tmp+arr[j-1])
			tmp = maxInt(tmp, dp[idx^1][j]) // O(1) 状态转移优化
		}
	}

	ans := -inf
	for i := m; i <= n; i++ {
		ans = maxInt(ans, dp[m&1][i])
	}
	return ans
}

func main() {
	arr := []int{1, 2, 3, 4, 5} // 15
	fmt.Println(maxMSubSum(arr, 3))

	arr = []int{1, 2, -100, 4, 5} // 12
	fmt.Println(maxMSubSum(arr, 3))

	arr = []int{5, -100, -100, -100, -10, 1}
	fmt.Println(maxMSubSum(arr, 3))

	arr = []int{-1, -2, -3, -1} // -1, -2, -1 --> -4
	fmt.Println(maxMSubSum(arr, 3))
}
