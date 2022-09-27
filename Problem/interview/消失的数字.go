package interview

// 题目：找出 1 ~ n 中消失的两个数字
// 思路：寻找一种方法能够将消失的这两个数字分在不同类中，然后再利用异或性质，便可以筛选出 x1 和 x2
// https://leetcode.cn/problems/missing-two-lcci/
//
// 其实，还有另外一种解法；先将整个数组扩充成大小为 n 的数组，最后两位置为 1
// 然后将数组中 1 ~ n 出现的数字进行归位，进行遍历，找寻丢失的数字
//
// 扩容本身可能导致空间占用变大，便不是 O(1) 空间复杂度

func missingTwo(nums []int) []int {
	n := len(nums) + 2

	xor := 0
	for i := 0; i < len(nums); i++ {
		xor ^= nums[i]
	}
	for i := 1; i <= n; i++ {
		xor ^= i
	}

	// 丢失的两个数字 xor = x1^x2
	// 按照最低为 1 的 bit 位分类，x1 和 x2 将会被划分到不同的组中
	// 由于其它数字均出现两次，因此分组后再做一遍异或便会求解出 x1 和 x2

	mark := xor & (-xor) // 最低位的有效数字

	x1, x2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&mark > 0 {
			x1 ^= nums[i]
		} else {
			x2 ^= nums[i]
		}
	}
	for i := 1; i <= n; i++ {
		if i&mark > 0 {
			x1 ^= i
		} else {
			x2 ^= i
		}
	}

	return []int{x1, x2}
}

// 利用求和性质求解参考：https://leetcode.cn/problems/missing-two-lcci/solution/by-ac_oier-pgeh/
