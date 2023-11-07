package main

// 1. 整数序列且相邻整数的【二进制表示】只有一位不同
// 2. 不同的构造方式生成的格雷码序列不是唯一的

//  1. 手动构造
//     从全 0 开始按照下述方式操作：
//     a. 翻转最低位得到下一个格雷码（例如，000 --> 001）
//     b. 把最右边的 1 的左边的位翻转得到下一个格雷码（例如，001 --> 011）
//     交替按照上述策略生成 2^(k-1) 次，得到 k 位格雷码序列
func generate(k int) []int {
	arr := make([]int, 0, 1<<k)
	arr = append(arr, 0)
	cnt, iterCnt := 0, 1<<(k-1)
	for cnt < iterCnt {
		x := arr[len(arr)-1]
		arr = append(arr, x^1)
		x = arr[len(arr)-1]
		arr = append(arr, x^((x&(-x))<<1))
		cnt++
	}
	return arr[:1<<k]
}

// 2. 递归构造
// 		k 位格雷码序列，可以通过 k-1 位格雷码序列构造而成
//      	a. 0 + {k-1 位格雷码正序}
//          b. 1 + {k-1 位格雷码逆序}
//      从 {0, 1} 序列开始，逐步构造 n 位格雷码序列

// 采用上述两种构造格雷码的方式，可以通过位置索引 i [0, 2^k) 计算出对应的格雷码
func G(i int) int {
	return i ^ (i >> 1)
}

// 格雷码逆变换：格雷码 --> 位置索引 i
func RG(g int) int {
	ret := 0
	for ; g > 0; g >>= 1 {
		ret ^= g
	}
	return ret
}

// [使整数变为 0 的最少操作次数] https://leetcode.cn/problems/minimum-one-bit-operations-to-make-integers-zero/
func minimumOneBitOperations(n int) int {
	var RG = func(g int) int {
		ret := 0
		for ; g > 0; g >>= 1 {
			ret ^= g
		}
		return ret
	}
	return RG(n)
}
