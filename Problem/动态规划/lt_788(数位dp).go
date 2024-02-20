package main

// 题目链接：https://leetcode.cn/problems/rotated-digits/
//
//
// 数位 dp 进行求解，具体参考代码实现

func rotatedDigits(n int) int {
	all := make([]int, 6)  // 全部可旋转的方案数
	same := make([]int, 6) // 全部可旋转且相同的方案数
	all[0] = 1
	same[0] = 1

	for i := 1; i <= 5; i++ {
		all[i] = all[i-1] * 7   // 2, 5, 6, 9
		same[i] = same[i-1] * 3 // 0, 1, 8
	}

	bits := make([]int, 0)
	for n > 0 {
		bits = append(bits, n%10)
		n /= 10
	}

	// bit 位反转
	blen := len(bits)
	for i, j := 0, blen-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}

	choose := []int{0, 1, 2, 5, 6, 8, 9}
	var dfs func(pos int, diff, bound bool) int
	dfs = func(pos int, diff, bound bool) int {
		if pos >= blen { //  边界处理
			if diff { // 存在 diff bit 位
				return 1
			}
			return 0
		}

		if !bound {
			if diff {
				return all[blen-pos]
			}
			return all[blen-pos] - same[blen-pos]
		}

		ret := 0
		for i := 0; i < len(choose) && choose[i] <= bits[pos]; i++ {
			tdiff, tbound := diff, bound
			if i == 2 || i == 3 || i == 4 || i == 6 {
				tdiff = true // bit 位 diff
			}

			if choose[i] < bits[pos] {
				tbound = false // 前缀 < 关系
			}

			ret += dfs(pos+1, tdiff, tbound)
		}
		return ret
	}

	return dfs(0, false, true)
}
