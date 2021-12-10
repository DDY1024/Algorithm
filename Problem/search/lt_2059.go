package main

// 题目链接：https://leetcode-cn.com/problems/minimum-operations-to-convert-number/
// 理解题目几个关键点才容易想到解题思路
// 1. 对于超出 [0, 1000] 范围内的整数，如果不等于 goal 值直接 pass 掉
// 2. 求解最少操作次数的性质与 bfs 搜索性质相同，且在进行 [0, 1000] 范围内判重的基础上实质上搜索范围不大
// 3. 不要被题目数据范围吓到，仔细分析题目本质
func minimumOperations(nums []int, start int, goal int) int {
	if start == goal {
		return 0
	}

	opList := []func(x, y int) int{
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
		func(x, y int) int { return x ^ y },
	}

	vis := make([]bool, 1001)
	n := len(nums)
	dis := make([]int, 1001)
	que := make([]int, 1001)
	for i := 0; i < n; i++ {
		dis[i] = 0x3f3f3f3f
	}
	dis[start] = 0
	vis[start] = true
	front, rear := 0, 1
	que[0] = start
	// 1000 * 3000 = 300w 次运算
	for front < rear {
		x := que[front]
		front++
		for i := 0; i < len(nums); i++ {
			for j := 0; j < 3; j++ {
				y := opList[j](x, nums[i])
				if y == goal {
					return dis[x] + 1
				}
				if y >= 0 && y <= 1000 && !vis[y] {
					que[rear] = y
					rear++
					dis[y] = dis[x] + 1
					vis[y] = true
				}
			}
		}
	}
	return -1
}
