package General

// 求解下一个排列算法
//  1. 非可重集全排列
//  2. 可重集全排列
func nextPermutaion(bits []byte) {
	n, idx := len(bits), len(bits)-2
	for idx >= 0 && bits[idx] > bits[idx+1] { // >=: 非可重集
		idx--
	}
	if idx < 0 { // 最后一个排列
		return
	}

	// bits[idx+1:n] 逆序，从大到小进行排列
	for i := n - 1; i > idx; i-- {
		if bits[i] > bits[idx] {
			bits[i], bits[idx] = bits[idx], bits[i] // 1. swap
			break
		}
	}

	// 2. reverse
	for i, j := idx+1, n-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}
}

// 1,2,3,4
// 1,2,4,3
// 1,3,2,4
// 1,3,4,2
