package General

// 时间复杂度 O(n^2)
// 1. 非可重集全排列
// 2. 可重集全排列

func nextPermutaion(bits []byte) {
	n, idx := len(bits), len(bits)-2
	// 1. 可重集全排列：判断标准为 >
	// 2. 非可重集全排列: 判断标准为 >=
	for idx >= 0 && bits[idx] > bits[idx+1] {
		idx--
	}
	if idx < 0 { // 最大排序
		return
	}

	// bits[idx+1:] 降序，从尾部向前查找第一个 > bits[idx] 的数
	for i := n - 1; i > idx; i-- {
		if bits[i] > bits[idx] {
			bits[i], bits[idx] = bits[idx], bits[i]
			break
		}
	}

	// 将 bit[idx+1:] 进行反转，变为升序
	for i, j := idx+1, n-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}
}
