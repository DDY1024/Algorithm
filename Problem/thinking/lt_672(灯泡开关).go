package thinking

// https://leetcode.cn/problems/bulb-switcher-ii/
// https://leetcode.cn/problems/bulb-switcher-ii/solution/dengp-by-capital-worker-51rb/
// https://leetcode.cn/problems/bulb-switcher-ii/solution/deng-pao-kai-guan-ii-by-leetcode-solutio-he7o/

// 规律总结
//
// 1. 灯泡的开关状态是以 6 为循环周期（循环周期）
// 2. 循环周期内每个灯泡的开关状态受操作的影响如下
// 		编号为 6k+1，受按钮 1,3,4 影响；
//      编号为 6k+2, 6k+6，受按钮 1,2 影响
//      编号为 6k+3, 6k+5，受按钮 1,3 影响；
//      编号为 6k+4，受按钮 1,2,4 影响
//
// 因此，我们只需要知道前 4 个灯泡的状态，即可以求得最终灯泡的状态

func countBit(x int) int {
	cnt := 0
	for x > 0 {
		x &= (x - 1)
		cnt++
	}
	return cnt
}

func getBit(x, i int) int {
	return (x >> i) & 1
}

func flipLights(n int, presses int) int {
	// 枚举 4 种操作的奇偶性情况，并统计前 4 个灯泡的不同状态
	mark := make(map[int]struct{}, 0)
	for i := 0; i < (1 << 4); i++ {
		if countBit(i) <= presses && countBit(i)&1 == presses&1 { // 奇偶性相同
			status := getBit(i, 0) ^ getBit(i, 2) ^ getBit(i, 3)
			if n >= 2 {
				status |= (getBit(i, 0) ^ getBit(i, 1)) << 1
			}
			if n >= 3 {
				status |= (getBit(i, 0) ^ getBit(i, 2)) << 2
			}
			if n >= 4 {
				status |= (getBit(i, 0) ^ getBit(i, 1) ^ getBit(i, 3)) << 3
			}
			mark[status] = struct{}{}
		}
	}
	return len(mark)
}
