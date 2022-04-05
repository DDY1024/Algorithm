package bit

// 计算正整数二进制表示中 1 的个数

func countBit(x int) int {
	cnt := 0
	for x > 0 {
		x &= (x - 1)
		cnt++
	}
	return cnt
}

// 快速统计二进制中 1 的个数
// O(log log n)
func countBitTwo(n uint32) uint32 {
	// 01,01,01,...,01
	// 0011,...,0011
	// 00001111,...,00001111
	// 0000000011111111,...,0000000011111111
	// 00000000000000001111111111111111
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	n = (n & 0x0f0f0f0f) + ((n >> 4) & 0x0f0f0f0f)
	n = (n & 0x00ff00ff) + ((n >> 8) & 0x00ff00ff)
	n = (n & 0x0000ffff) + ((n >> 16) & 0x0000ffff)
	return n
}

// 同理求解 uint64
