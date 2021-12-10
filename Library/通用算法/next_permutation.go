package General

// 下一个排列: 非可重集、可重集
// 算法复杂度: O(n^2)
func nextPermutaion(bits []byte) {
	n, idx := len(bits), len(bits)-2
	for idx >= 0 && bits[idx] >= bits[idx+1] { // >= : 可重集、非可重集
		idx--
	}
	if idx < 0 {
		return
	}

	for i := n - 1; i > idx; i-- {
		if bits[i] > bits[idx] {
			bits[i], bits[idx] = bits[idx], bits[i]
			break
		}
	}
	for i, j := idx+1, n-1; i < j; i, j = i+1, j-1 {
		bits[i], bits[j] = bits[j], bits[i]
	}
}
