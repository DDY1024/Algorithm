package main

// N 皇后【位运算】版本
// http://www.matrix67.com/blog/archives/266

// row: 行占用
// col: 列占用
// ld: 左对角线占用  \
// rd: 右对角线占用  /
func solve(n int) int {
	var (
		total = 0
		mask  = (1 << n) - 1
	)

	var dfs func(row, col, ld, rd int)
	dfs = func(row, col, ld, rd int) {
		if row >= n {
			total++
			return
		}

		// 获取第 row 行可以放置的列的位置
		can := (^(col | ld | rd)) & mask
		for can > 0 {
			p := can & (-can)
			dfs(row+1, col|p, (ld|p)>>1, (rd|p)<<1)
			can -= p
		}
	}

	dfs(0, 0, 0, 0)
	return total
}
