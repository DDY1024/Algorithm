package main

// 一、 0-1 背包问题
//
// 1. 石头粉碎问题
/*
问题：有一堆石头，每块石头的重量都是正整数。每次从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x≤y。
那么粉碎的可能结果如下：
	1. 如果 x 与 y 相等，那么两块石头都会被完全粉碎；
	2. 否则，重量为 x 的石头将会完全粉碎，而重量为 y 的石头的新重量为 y−x。
最后，最多只会剩下一块石头。返回此时石头最小的可能重量。如果没有石头剩下，就返回 0。

解题思路:
	1. 我们的操作序列中，每块肯定是作为 "减数" 或 "被减数" 其中之一
	2. 因此我们求解的目标，其实可以等价为 y1 + y2 + ... + yk - x1 - x2 - xl 的差值最小
	3. 2 中问题模型的最优方案是使得 y 和 x 系列数的和最接近总和的一半
	4. 如何寻找出一系列数使得其和接近总和的一半，我们可以转化成 0-1 背包问题进行求解的
*/
