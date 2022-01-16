package hard2022

// 题目链接: https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/
// 题目大意
// 计算右侧小于当前元素的个数，由于题目中元素数据范围 [-10000,10000]，做下预处理后，可以利用树状数组求解
//
// 卧槽，这是个逆序对问题，即可以利用树状数组进行求解，也可以利用归并排序进行求解
//
// 如何利用归并排序的过程求解逆序对数???

func lowbit(x int) int {
	return x & (-x)
}

func add(x, n, c int, arr []int) {
	for i := x; i <= n; i += lowbit(i) {
		arr[i] += c
	}
}

func sum(x int, arr []int) int {
	ret := 0
	for i := x; i > 0; i -= lowbit(i) {
		ret += arr[i]
	}
	return ret
}

func countSmaller(nums []int) []int {
	n, delta, upper := len(nums), 20000, 30000
	arr := make([]int, 40010)
	ret := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		nums[i] += delta
		ret[i] = sum(nums[i]-1, arr)
		add(nums[i], upper, 1, arr)
	}
	return ret
}
