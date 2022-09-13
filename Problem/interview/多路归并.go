package main

// 优化: 小顶堆维护，进一步降低时间复杂度
func kthSmallest(matrix [][]int, k int) int {
	n, cc := len(matrix), 0
	pos, midx := make([]int, n), -1
	for cc < k {
		minVal := 0x3f3f3f3f
		for i := 0; i < n; i++ {
			if pos[i] < n && minVal > matrix[i][pos[i]] {
				minVal = matrix[i][pos[i]]
				midx = i
			}
		}
		pos[midx]++
		cc++
	}
	return matrix[midx][pos[midx]-1]
}
