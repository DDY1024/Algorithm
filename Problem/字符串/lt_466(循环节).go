package main

// 题目链接：https://leetcode.cn/problems/count-the-repetitions/?envType=daily-question&envId=2024-01-02
//
// 1. 处理【字符串构造】的循环节
// 2. 具体看注释，比较考察思维

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	l1, l2 := len(s1), len(s2)
	if l1 == 0 || l2 == 0 || l1*n1 < l2*n2 {
		return 0
	}

	// idxMark: 第 i 个 s1 最初匹配到 s2 的索引 j
	// cntMark: 每个 s1 内匹配 s2 的第一个索引位置 i 对应的第 j 个 s1（判断循环节）
	idxMark, cntMark := make(map[int]int, 0), make(map[int]int, 0)

	// cnt: 当前 s1 个数
	// idx: 当前 s2 匹配的索引
	cnt, idx := 0, 0

	// cycleCnt：一个循环节需要占用多少个 s1
	// cycleS2Cnt：一个循环节可以构造多少个 s2
	cycleCnt, cycleS2Cnt := -1, -1

	// left : 去除【循环节】和【初始构造】，剩余多少个 s1
	ret, left := 0, 0
	// 最坏情况下 l2 个 s1 才能构造出一个 s2；因此【循环节】验证最多迭代【l2+1】次即可
	for cnt <= l2+1 { // 最多迭代 l2+1 次
		cnt++
		// 循环节判断
		if pos, ok := cntMark[idx%l2]; ok {
			cycleCnt = cnt - pos                   // 计算循环节长度
			cycleS2Cnt = (idx - idxMark[pos]) / l2 // 根据 idxMark 维护的索引信息计算可以构造多少个 s2
			// 结果计算【三部分】: [1, p), 根据 [pos, cnt) 计算循环节部分, left 剩余部分
			ret = (idxMark[pos]-1+1)/l2 + (n1-pos+1)/cycleCnt*cycleS2Cnt
			left = (n1 - pos + 1) % cycleCnt
			break
		}

		idxMark[cnt] = idx
		cntMark[idx%l2] = cnt
		for i := 0; i < l1; i++ {
			if s1[i] == s2[idx%l2] {
				idx++
			}
		}
	}

	// 无法构造 s2，不存在循环节
	if cycleCnt < 0 {
		return 0
	}

	// 计算剩余的 left 个 s1 能够构造多少个 s2
	// 起始匹配索引：idxMark[cnt]
	for i := 0; i < left; i++ {
		for j := 0; j < l1; j++ {
			if s1[j] == s2[idx%l2] {
				idx++
				if idx%l2 == 0 { // 可以构造一个 s2
					ret++
				}
			}
		}
	}

	// 总共可以构造 ret 个 s1
	return ret / n2
}
