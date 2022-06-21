package main

// 三指针解法，同样像双指针的方式进行统计
// 双指针 --> 三指针
// 双指针 --> 多指针
// https://leetcode.com/problems/subarrays-with-k-different-integers/
// 无论是双指针、三指针、N 指针等情况，我们只需要保证均摊下来的复杂度是 O(N) 即可
// 因此，N 指针其实是利用所有指针不断推进，求取最终结果的一个思想，属于比较巧妙的一种思路

func subarraysWithKDistinct(A []int, K int) int {
	n := len(A)
	markl := make([]int, n+1)
	markr := make([]int, n+1)
	l, r, dcl, dcr, ans := 0, 0, 0, 0, 0
	for i := 0; i < n; i++ {
		if markl[A[i]] == 0 {
			dcl++
		}
		if markr[A[i]] == 0 {
			dcr++
		}
		markl[A[i]]++
		markr[A[i]]++
		for dcl > K {  // move l
			if markl[A[l]] - 1 == 0 {
				dcl--
			}
			markl[A[l]]--
			l++
		}
		for dcr >= K && r <= i {
			if markr[A[r]] - 1 == 0 {
				dcr--
			}
			markr[A[r]]--
			r++
		}
		if dcl == K && dcr == K - 1 {
			ans += r - l
		}
	}
	return ans
}

func main() {

}