package main

// 参考自：http://www.matrix67.com/blog/archives/266
// N 皇后位运算解法

var ls = func(x int) int { return 1 << uint(x) }

func NQueen(n int) int {
	mask := (1 << n) - 1
	total := 0

	// col: 列占用
	// ld: 左对角线占用
	// rd: 右对角线占用
	var dfs func(row, col, ld, rd int)
	dfs = func(row, col, ld, rd int) {
		if row >= n {
			total++
			return
		}

		// 反转 bit 位获取第 row 行可以放置的列
		ch := (^(col | ld | rd)) & mask

		// 方法一
		// for i := 0; i < n; i++ {
		// 	if pos&leftShift(i) > 0 {
		// 		dfs(row+1, col|leftShift(i), (ld|leftShift(i))<<1, (rd|leftShift(i))>>1)
		// 	}
		// }

		// 方法二: 直接遍历所有可以放置的位置
		for ch > 0 {
			pos := ch & (-ch)
			// 列: col|pos
			// 主对角线（撇）: (rd|pos)<<1
			// 副对角线（捺）: (ld|pos)>>1
			dfs(row+1, col|pos, (ld|pos)>>1, (rd|pos)<<1)
			ch &= ch - 1
		}
	}

	dfs(0, 0, 0, 0)
	return total
}
