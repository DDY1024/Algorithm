package main

func main() {
	// a := new(big.Int)
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	if n == 1 {
		return []int{arr[0], arr[n-1]}
	}

	var calc = func(i1, j1 int) int {
		cnt := 0
		for i2 := 0; i2 < n-1; i2++ {
			l2, r2, ret := i2+1, n-1, -1
			for l2 <= r2 {
				m2 := l2 + (r2-l2)/2
				if arr[i1]*arr[m2] <= arr[i2]*arr[j1] {
					l2 = m2 + 1
				} else {
					ret = m2
					l2 = m2 - 1
				}
			}
			if ret != -1 {
				cnt += ret - i2
			}
		}
		return cnt
	}

	for i := 0; i < n-1; i++ {
		l1, r1 := i+1, n-1
		for l1 <= r1 {
			m1 := l1 + (r1-l1)/2
			rr := calc(i, m1)
			if rr == k-1 {
				return []int{arr[i], arr[m1]}
			}
			if rr > k-1 {
				l1 = m1 + 1
			} else {
				l1 = m1 - 1
			}
		}
	}
	return []int{-1, -1}
}
