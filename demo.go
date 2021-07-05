package main

import "fmt"

func powMod(a, b, c int) int {
	ret := 1
	for b > 0 {
		if b&1 > 0 {
			ret = ret * a % c
		}
		a = a * a % c
		b >>= 1
	}
	return ret
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 利用 rabin-karp 滚动哈希的思想进行求解
func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)

	// A[i]，B[i] <= 100
	mod := int(1e9 + 7) // mod 要大一些
	base := 113

	var check = func(l int) bool {
		mark := make(map[int]bool)
		hash := 0
		tmp := powMod(base, l-1, mod)
		for i := 0; i < l; i++ {
			hash = (hash*base + nums1[i]) % mod
		}
		mark[hash] = true
		for i := l; i < n; i++ {
			hash = (((hash-nums1[i-l]*tmp%mod)+mod)%mod*base + nums1[i]) % mod
			mark[hash] = true
		}

		hash = 0
		for i := 0; i < l; i++ {
			hash = (hash*base + nums2[i]) % mod
		}
		if mark[hash] {
			return true
		}

		for i := l; i < m; i++ {
			hash = (((hash-nums2[i-l]*tmp%mod)+mod)%mod*base + nums2[i]) % mod
			if mark[hash] {
				return true
			}
		}
		return false
	}

	l, r, ans := 1, minInt(n, m), 0
	for l <= r {
		mid := l + (r-l)/2
		if check(mid) {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}
	}
	return ans
}

func main() {
	// fmt.Println("hello, world!")
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
}
