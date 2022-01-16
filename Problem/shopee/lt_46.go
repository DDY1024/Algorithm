package shopee

// 题目链接: https://leetcode-cn.com/problems/permutations/
// 题目大意

// 1. 非可重集的全排列
func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	n := len(nums)

	var do func(pos int)
	do = func(pos int) {
		if pos >= n {
			tmp := make([]int, n)
			for i := 0; i < n; i++ {
				tmp[i] = nums[i]
			}
			ret = append(ret, tmp)
			return
		}

		for i := pos; i < n; i++ {
			nums[pos], nums[i] = nums[i], nums[pos]
			do(pos + 1)
			nums[pos], nums[i] = nums[i], nums[pos]
		}
	}
	do(0)
	return ret
}

// 2. 可重集全排列
// 相对于非可重集全排列，只是在选择 pos 位置的元素时需要判重
//
