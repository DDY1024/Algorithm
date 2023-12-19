package main

// 题目链接： https://leetcode.cn/problems/find-a-peak-element-ii/?envType=daily-question&envId=Invalid%20Date
// 解题报告： https://leetcode.cn/problems/find-a-peak-element-ii/solutions/2566062/xun-zhao-feng-zhi-ii-by-leetcode-solutio-y57g/?envType=daily-question&envId=Invalid+Date

func findPeakGrid(mat [][]int) []int {
	n := len(mat)

	maxElem := func(arr []int) int {
		idx := 0
		for i := 0; i < len(arr); i++ {
			if arr[idx] < arr[i] {
				idx = i
			}
		}
		return idx
	}

	// 1. 满足条件的解当中，必然存在一个元素是【所属行的最大值】，因此我们可以只考虑每一行的最大值
	// 2. 如果 mat[i-1][j] > mat[i][j]，上一行的最大值必然大于【同列】下一行的值，至于是否大于【同列】上一行的值，最坏情况在第 0 行必然满足（i < 0, val = -1），则存在【解】出现在【上部分】 --> 反证法
	// 3. 如果 mat[i+1][j] > mat[i][j]，下一行的最大值必然大于【同列】上一行的值，至于是否大于【同列】下一行的值，最坏情况在第 n-1 行必然满足（i >= n, val = -1），则存在【解】出现在【下部分】 --> 反证法
	l, r := 0, n-1
	for l <= r {
		i := l + (r-l)/2
		j := maxElem(mat[i])

		if i-1 >= 0 && mat[i-1][j] > mat[i][j] {
			r = i - 1
			continue
		}

		if i+1 < n && mat[i+1][j] > mat[i][j] {
			l = i + 1
			continue
		}

		// mat[i-1][j] < mat[i][j] > mat[i+1][j]
		return []int{i, j}
	}

	return nil // impossible
}
