package main

import (
	"math/rand"
	"time"
)

// 蓄水池抽样算法
// https://leetcode-cn.com/problems/linked-list-random-node/solution/gong-shui-san-xie-xu-shui-chi-chou-yang-1lp9d/

// 解决问题
//      当内存无法加载全部数据时，如何从包含未知大小的数据流中随机选取 k 个数据，并且要保证每个数据被抽取到的概率相等

// 问题一
//      当 k = 1 时，数据流含有 N 个数，如果要保证所有的数被抽到的概率相等，那么每个数抽到的概率应该为 1/N
//      这个问题的解决方案本质上和洗牌算法是相同的，洗牌算法本身保证了每个数在某一位置出现的概率相等

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	// rand.Seed(time.Now().UnixNano())
}

func solveOne(data []int) int {
	ret := data[0]
	for i := 1; i < len(data); i++ {
		if r.Intn(i+1) == 0 {
			ret = data[i]
		}
	}
	return ret
}

// 问题二
//      当 k = m 时，由最初选择一个元素变成了最初选择 m 个元素；然后从第 m+1 元素开始，进行 random ，如果落在区间 [0,m-1] 内
// 则随机替换其中的一个元素

func solveTwo(data []int, k int) []int {
	n := len(data)
	for i := k; i < n; i++ {
		p := r.Intn(i + 1) //
		if p < k {
			data[i], data[p] = data[p], data[i]
		}
	}
	return data[:k]
}

// 问题三
//      分布式蓄水池抽样算法：https://www.jianshu.com/p/7a9ea6ece2af
//
//      假设存在 K 台机器，将整个数据集分成 K 份，并最终统计每份数据集的大小为 N1, N2, ..., Nk
//      a. 对于每份数据集采用蓄水池抽样算法选取 m 个数，每个数被选择的概率 m/Nk
//      b. 接下来的问题，便是将选出来的这些数的概率变成 m/N ？
//          每次从 N 中随机选择一个数，如果落在区间 [0,N1) 则从第一份中拿出一个数，如果落在 [N1,N1+N2) 则从第二份中获取一个数，依次类推;
//          重复 m 次, 最终选择出 m 个数
