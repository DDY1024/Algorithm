package main

// 官方解题报告：https://leetcode.cn/problems/burst-balloons/solution/chuo-qi-qiu-by-leetcode-solution/
//
//  记忆化搜索（自顶向下）  vs  动态规划（自底向上）
//
//
// 思考过程
//
//       正向删除 --> 逆向添加
//
// 1. 考虑正向删除过程，将 i 删除后，会使 i-1 和 i+1 变成相邻，难以处理，无法用动态规划求解
// 2. 考虑逆向添加过程，例如将 [i,j] 区间一次添加完元素后，能够获得的最大分数值，这样可以通过枚举 [i,j] 区间内第一个添加的数字
//  	来划分成子区间进行求解
// 4. 提供一种巧妙的处理思路
// 		dp[i][j] 表示 i, j 已经被添加过的情况下，填充完该开区间 (i, j) 内所有元素后能够获得的最大收益值
//      因为我们已经有了边界 i 和 j 具体取值 nums[i] 和 nums[j], 那这样一来在 (i,j) 开区间内首次添加的元素，我们是可以计算其收益值的
//      很容易得到
//      dp[i][j] = max{ nums[i] * nums[k] * nums[j] + dp[i][k] + dp[k][j] }，其中 k 为 (i, j) 取值
//   为了得到最终结果，我们需要特殊处理下边界，即 nums[0], nums[n+1] 表示已经添加完的元素了，均为 1，[1,n] 表示我们待填充的全部元素
//   因此最终结果 dp[0][n+1] 便是可以正确地表示我们这一填充的过程，问题得以解决，比较巧妙。
//
//
//   闭区间  ---->  开区间
//
//

func maxCoins(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var get = func(idx int) int {
		if idx > 0 && idx <= n {
			return nums[idx-1]
		}
		return 1 // 特殊边界: 全部填充为 1
	}

	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	for l := 3; l <= n+2; l++ {
		for i := 0; i+l-1 < n+2; i++ {
			j := i + l - 1
			// (i,j) 开区间填充完所有元素的最大收益
			for k := i + 1; k < j; k++ {
				dp[i][j] = maxInt(dp[i][j], dp[i][k]+dp[k][j]+get(i)*get(k)*get(j))
			}
		}
	}

	return dp[0][n+1]
}
