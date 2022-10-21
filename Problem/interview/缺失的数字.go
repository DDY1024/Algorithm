package interview

// 题目大意：寻找 1 ~ n 中消失的两个数字
// 解题思路：
// 		1. 寻找一种方法能够将消失的这两个数字分在不同组中；
//      2. 然后再利用异或性质；便可以筛选出 x1 和 x2
// 		https://leetcode.cn/problems/missing-two-lcci/
//
// 其实，还有另外一种解法；先将整个数组扩充成大小为 n 的数组，最后两位置为 1
// 然后将数组中 1 ~ n 出现的数字进行归位，进行遍历，找寻丢失的数字
//
// 扩容本身可能导致空间占用变大，便不是 O(1) 空间复杂度
//

// 1. 找到分类方法
// 2. 利用异或性质求解

func missingTwo(nums []int) []int {
	n, xor := len(nums)+2, 0

	// nums[0]、nums[1]、... 、nums[n-1]
	for i := 0; i < len(nums); i++ {
		xor ^= nums[i]
	}
	for i := 1; i <= n; i++ {
		xor ^= i
	}

	// 丢失的两个数字 xor = x1^x2
	// 按照最低为 1 的 bit 位分类，x1 和 x2 将会被划分到不同的组中
	// 由于其它数字均出现两次，因此分组后再做一遍异或便会求解出 x1 和 x2

	// 此时 xor = x1 ^ x2
	// 取最低的 bit 位为 1，进行数组元素划分
	mark := xor & (-xor)
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

// O(n) 时间复杂度，O(1) 空间复杂度解决
// 其实就是通过不断地将数字换到其对应位置即可
func missingNumber(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] < n && nums[i] != i {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	// 0 ~ n-1 存在一个缺失的
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}

	// 0 ~ n-1 均不缺失，则 n 缺失
	return n
}
