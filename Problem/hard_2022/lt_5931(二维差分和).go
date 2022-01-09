package hard2022

// 题目链接: https://leetcode-cn.com/problems/stamping-the-grid/
// 题目大意: 用可重叠的邮票是否可以完全覆盖所有的空位
//
// 解题思路
// 1 <= m, n <= 10^5, 1 <= m * n <= 2 * 10^5
//
// 1. 首先利用二维子矩阵和求解出所有可以放置邮票的左上角端点
// 2. 对于一个空位是否被邮票覆盖，即判断以该点为右下端点的子矩形内是否存在可放置邮票的顶点
// 3. 1 和 2 两个条件的判断均用到了二维数组的差分和，求解方法值得总结

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

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	n, m := len(grid), len(grid[0])
	sum1 := make([][]int, n+1)
	sum2 := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		sum1[i] = make([]int, m+1)
		sum2[i] = make([]int, m+1)
	}

	var calcSum = func(i1, j1, i2, j2 int, arr [][]int) int {
		ret := arr[i2][j2]
		ret -= arr[i2][j1-1]
		ret -= arr[i1-1][j2]
		ret += arr[i1-1][j1-1]
		return ret
	}

	// 1. 求解占据位置的二维前缀和
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			sum1[i][j] = grid[i-1][j-1]
			sum1[i][j] += sum1[i][j-1]
			sum1[i][j] += sum1[i-1][j]
			sum1[i][j] -= sum1[i-1][j-1]
		}
	}

	// 2. 找出所有可以放置邮票的左上端点
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if grid[i-1][j-1] == 0 && i+stampHeight-1 <= n && j+stampWidth-1 <= m &&
				calcSum(i, j, i+stampHeight-1, j+stampWidth-1, sum1) == 0 { //
				sum2[i][j] = 1
			}
		}
	}

	// 3. 判断当前点是否可被某个左上端点覆盖
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 直接在判断位置时求解前缀和，少一次二维数组的遍历
			sum2[i][j] += sum2[i][j-1]
			sum2[i][j] += sum2[i-1][j]
			sum2[i][j] -= sum2[i-1][j-1]
			if grid[i-1][j-1] == 0 {
				i1, j1 := maxInt(1, i-stampHeight+1), maxInt(1, j-stampWidth+1)
				if calcSum(i1, j1, i, j, sum2) == 0 {
					return false
				}
			}
		}
	}
	return true
}
