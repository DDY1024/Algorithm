package main

import "fmt"

// 原文出自: http://www.matrix67.com/blog/archives/266
// https://www.cxyxiaowu.com/8990.html

var leftShift = func(x int) int { return 1 << uint(x) }

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

		// 可摆放位置通过反转 0 --> 1
		ch := (^(col | ld | rd)) & mask

		// 方法一
		// for i := 0; i < n; i++ {
		// 	if pos&leftShift(i) > 0 {
		// 		dfs(row+1, col|leftShift(i), (ld|leftShift(i))<<1, (rd|leftShift(i))>>1)
		// 	}
		// }

		// 方法二
		for ch > 0 {
			pos := ch & (-ch)
			dfs(row+1, col|pos, (ld|pos)<<1, (rd|pos)>>1)
			ch &= ch - 1
		}
	}

	dfs(0, 0, 0, 0)
	return total
}

func main() {
	fmt.Println(NQueen(8))
}
