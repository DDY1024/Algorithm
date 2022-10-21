package main

func permute(nums []int) [][]int {
	n := len(nums)

	var ans [][]int
	var do func(pos int)
	do = func(pos int) {
		if pos >= n {
			ans = append(ans, append([]int(nil), nums...)) // slice 拷贝
			return
		}

		for i := pos; i < n; i++ {
			nums[pos], nums[i] = nums[i], nums[pos]
			do(pos + 1)
			nums[pos], nums[i] = nums[i], nums[pos]
		}
	}
	do(0)

	return ans
}
