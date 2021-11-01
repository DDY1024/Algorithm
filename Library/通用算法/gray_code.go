package main

import "fmt"

// 参考资料：https://oi-wiki.org/misc/gray-code/
// 格雷码引用 leetcode 1611：https://leetcode.com/contest/weekly-contest-209/problems/minimum-one-bit-operations-to-make-integers-zero/
//
// 格雷码性质：一个二进制数系，其中两个相邻数的二进制位只有一位不同。例如 3 位二进制数的格雷码序列为：000,001,011,010,110,111,101,100
// 000,001,011,010,110,111,101,100
// 000,001,010,011,100,101,110,111
//
// 构造格雷码的两种方法
// 一、手动构造
// 1. 翻转最低位得到下一个格雷码，例如 000 --> 001
// 2. 把最右边的 1 的左边的位翻转得到下一个格雷码，例如 001 --> 011
// 3. 交替按照上述策略生成 2^k-1 次，即可得到 k 位的格雷码序列
//
// 二、镜像构造（递归构造）
// 讲述一种从 k-1 位格雷码序列构造 k 位格雷码序列的方法
// 1. 0gray(k-1),1gary(k-1)逆序

// 格雷码计算方法
// k 位格雷码序列第 i 个格雷码即为 G(i)
func G(n int) int {
	return n ^ (n >> 1)
}

// 格雷码逆计算，即 G(i) --> i
// i --> G(i) --> i
// 格雷码与二进制体系一一对应
func RG(g int) int {
	n := 0
	for g > 0 {
		n ^= g
		g >>= 1
	}
	return n
}

// 且记格雷码的正向计算和逆向计算

func main() {
	for i := 0; i < 8; i++ {
		fmt.Printf("%03b --> %03b\n", G(i), RG(G(i)))
	}
}
