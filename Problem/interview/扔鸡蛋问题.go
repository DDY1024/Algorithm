package interview

// 提供两种不同的用动态规划求解的思路
// https://leetcode.cn/problems/super-egg-drop/solution/ji-dan-diao-luo-by-leetcode-solution-2/
//
// 第一种求解思路很直白
// dp[i][j]：i 个鸡蛋检测 j 层楼最少需要操作多少次
// 状态转移方程：分鸡蛋碎了和没碎两种情况进行讨论
// dp[i][j] = min{max{dp[i][j-x],dp[i-1][x-1]}+1}, 其中 1 <= x <= j
//
// 利用 dp[i][j] 在 i 固定时，随着 j 单调递增
// 利用 dp[i][j] 在 j 固定时，随着 i 单调递减
// 利用二分加速确定 min{max{}} 操作的最优点

func superEggDrop(k int, n int) int {
	var minInt = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var maxInt = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[1][i] = i
	}

	for i := 2; i <= k; i++ {
		dp[i][1] = 1
		for j := 2; j <= n; j++ {
			// min{max{dp[i][j-x], dp[i-1][x-1]}+1}
			// 由于存在单调性，利用二分法寻找最优点，优化状态转移复杂度
			l, r, x := 1, j, j
			for l <= r {
				mid := l + (r-l)>>1
				if dp[i-1][mid-1] >= dp[i][j-mid] {
					x = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			//dp[i][j] = minInt(maxInt(dp[i][j-x], dp[i-1][x-1])) + 1
			// if x > 1 {
			//     x--
			//     dp[i][j] = minInt(dp[i][j], maxInt(dp[][]))
			// }
			dp[i][j] = 0x3f3f3f3f // inf
			for k := x; k > 0 && k > x-2; k-- {
				dp[i][j] = minInt(dp[i][j], maxInt(dp[i][j-k], dp[i-1][k-1])+1)
			}

		}
	}
	return dp[k][n]
}

// 另外一种动态规划求解思路:
// dp[i][j]：i 次操作，j 个鸡蛋最多可以确定的楼层数
func superEggDropTwo(k int, n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][1] = i // i 次操作 1 个鸡蛋，最多能确定 i 层
	}

	for i := 1; i <= k; i++ {
		dp[1][i] = 1 // 1 次操作 i 个鸡蛋，最多只能确定 1 层
	}

	if dp[1][k] >= n {
		return 1
	}

	for i := 2; ; i++ {
		for j := 2; j <= k; j++ {
			dp[i][j] = dp[i-1][j] + // 鸡蛋没碎，上面还可以检测 dp[i-1][j] 层
				dp[i-1][j-1] + // 鸡蛋碎了，下面还可以检测 dp[i-1][j-1] 层
				1 // 加上当前检测的这一层
		}
		// dp[i][1] 本身可能已经 >= n 了，在循环内处理，可能会漏掉解
		// 所以我们只需要在循环外，直接判断 dp[i][k] >= n 即可，毕竟满足单调性嘛
		if dp[i][k] >= n {
			return i
		}
	}
}
