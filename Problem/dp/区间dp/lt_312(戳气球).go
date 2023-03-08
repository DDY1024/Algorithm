package main

// 解题报告：https://leetcode.cn/problems/burst-balloons/solution/chuo-qi-qiu-by-leetcode-solution/
//
// 1. 正向删除，会改变数组元素相邻的结构，难以处理
// 2. 逆向添加，每次新增一个新元素后，数组结构是确定的，方便处理
// 3. 状态定义：闭区间 --> 开区间
//   	由于闭区间 [i,j] 需要考虑填充 i 和 j 位置时，其左右相邻的元素，无法确定
//      开区间 (i,j) 只需要考虑 [i+1,j-1] 区间内的元素填充，且边界已经确定为 i 和 j
// 4. 综上所述
// 		状态定义：dp[i][j] 表示填充 (i, j) 区间内的元素能够获得的最大收益
//      状态转移：dp[i][j] = max{nums[i]*nums[k]*nums[j] + dp[i][k] + dp[k][j]}, i < k < j
//    最终结果：dp[0][n+1]

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxCoins(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	var get = func(idx int) int {
		if idx < 1 || idx > n {
			return 1 // 边界默认值为 1
		}
		return nums[idx-1]
	}

	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	// 由于我们处理的是开区间，因此对于(i,i+1)这类长度为 2 的开区间结果为 0；我们从长度为 3 的区间开始处理
	for l := 3; l <= n+2; l++ {
		for i := 0; i+l-1 < n+2; i++ {
			j := i + l - 1
			for k := i + 1; k < j; k++ {
				dp[i][j] = maxInt(dp[i][j], dp[i][k]+dp[k][j]+get(i)*get(k)*get(j))
			}
		}
	}
	return dp[0][n+1]
}
