package main

// 题目链接: https://leetcode-cn.com/problems/super-egg-drop/
// 题目大意
// 经典扔鸡蛋问题，动态规划进行求解
// k 个鸡蛋 n 层楼，使用最少的操作次数求解出鸡蛋的楼层上限数 f
//
// 状态转移方程
// dp[i][j]: i 层楼 j 个鸡蛋测试，需要的最少操作次数
// dp[i][1] = i
// dp[i][0] = inf
// dp[0][i] = 0
//
// dp[i][j] = min{max{dp[k-1][j-1],dp[i-k][j]}} + 1, 其中 1 <= k <= i
// 选择每个楼层的情况为碎和不碎两种情况，我们需要求解两种情况中的最坏情况是什么
//
// 接下来是如何优化状态转移方程？？？
// dp[x][j]: j 固定 x 单调递增
// dp[i][y]: i 固定 y 单调递减
//
//
// https://leetcode-cn.com/problems/super-egg-drop/solution/ji-dan-diao-luo-by-leetcode-solution-2/
// 通过官方的解题报告，我们学习到了一种优化状态转移的方法。
// 两个函数一个单调递增，一个单调递减；则 min{max{f1, f2}} 的最优值是出现在两个函数的交点处的
// 我们可以通过二分的方法查找 f1 >= f2 的 x0 和 f2 >= f1 的 x1，取 x0、x1 处的最优值即可
//
// 一个单调递增函数 f1，一个单调递减函数 f2
// min{max{f1, f2}} 操作如何利用二分法进行优化
// 两个单调函数的交点附近是最优结果

func superEggDrop(k int, n int) int {

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][1] = i
	}

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

	for j := 2; j <= k; j++ {
		for i := 1; i <= n; i++ {
			l, r, x0 := 1, i, i
			for l <= r {
				mid := l + (r-l)/2
				if dp[mid-1][j-1] >= dp[i-mid][j] {
					x0 = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			l, r, x1 := 1, i, 1
			for l <= r {
				mid := l + (r-l)/2
				if dp[i-mid][j] >= dp[mid-1][j-1] {
					x1 = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			dp[i][j] = minInt(maxInt(dp[x1-1][j-1], dp[i-x1][j]), maxInt(dp[x0-1][j-1], dp[i-x0][j])) + 1
		}
	}
	return dp[n][k]
}

// 另外一种比较难以想到的方法，但是更优，具体参考官方解题报告
// 如果我们可以做 t 次操作，而且有 k 个鸡蛋，那么我们能找到答案的最高的 n 是多少?
// dp[t][1] = t
// dp[1][k] = 1
// dp[t][k] = 1 + dp[t-1][k-1](鸡蛋碎了，在下方) + dp[t-1][k](鸡蛋没碎，在上方)

func superEggDropTwo(k int, n int) int {
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][1] = i // i 次操作一个鸡蛋
	}

	for i := 1; i <= k; i++ {
		dp[1][i] = 1 // 1 次操作 i 个鸡蛋
	}

	// 寻找最小的 i 使得 dp[i][k] >= n
	if dp[1][k] >= n {
		return 1
	}

	for i := 2; i <= n; i++ { // 上线是 n，n 次操作肯定是可以试验出 n 层楼的
		for j := 2; j <= k; j++ {
			dp[i][j] = 1 + dp[i-1][j-1] + dp[i-1][j]
		}
		if dp[i][k] >= n {
			return i
		}
	}
	return -1 // not arrive
}
