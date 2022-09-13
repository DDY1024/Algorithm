package main

// 蓄水池抽样算法
// 参考: https://leetcode-cn.com/problems/linked-list-random-node/solution/xu-shui-chi-chou-yang-suan-fa-by-jackwener/
// https://leetcode-cn.com/problems/linked-list-random-node/solution/gong-shui-san-xie-xu-shui-chi-chou-yang-1lp9d/
// 大数据处理随机算法
//
// 问题描述: 当内存无法加载全部数据时，如何从包含未知大小的数据流中随机选取k个数据，并且要保证每个数据被抽取到的概率相等

// 1. k = 1 时，即数据流含有N个数，我们知道如果要保证所有的数被抽到的概率相等，那么每个数抽到的概率应该为 1/N
//
// 方案: 遍历数据流中的所有整数，当遇到第 i 个数时，以 1/i 概率保留它，1-1/i 概率保留原来的数，这样可以保证每个数被选择的概率均为 1/n

/*
import random
class Solution:

    def __init__(self, head: ListNode):
        self.head = head

    def getRandom(self) -> int:
        count = 0
        reserve = 0
        cur = self.head
        while cur:
            count += 1
            rand = random.randint(1,count)
            if rand == count:
                reserve = cur.val
            cur = cur.next
        return reserve
*/

// 2. 当 k = m 时，由最初选择一个元素变成了最初选择 m 个元素；然后从第 m+1 元素开始，进行 random ，如果落在区间 [0,m-1] 内
// 则随机替换其中的一个元素
//
// 参考资料: https://www.jianshu.com/p/7a9ea6ece2af
//
//
/*
int[] reservoir = new int[m];

// init
for (int i = 0; i < reservoir.length; i++)
{
    reservoir[i] = dataStream[i];
}

for (int i = m; i < dataStream.length; i++)
{
    // 随机获得一个[0, i]内的随机整数
    int d = rand.nextInt(i + 1);
    // 如果随机整数落在[0, m-1]范围内，则替换蓄水池中的元素
    if (d < m)
    {
        reservoir[d] = dataStream[i];
    }
}
*/

// 3. 分布式蓄水池抽样算法
// https://www.jianshu.com/p/7a9ea6ece2af
//
// 假设存在 K 台机器，将这整个数据集分成 K 份，并最终统计每份数据集的大小为 N1, N2, ..., Nk
// a. 对于每份数据采用蓄水池抽样算法选取 m 个数，每个数被选择的概率 m/Nk
// b. 接下来的问题，便是将选出来的这些数的概率编程 m/N 的问题？
// 我们可以这样，每次从 N 中随机选择一个数，如果落在区间 [0,N1) 则从第一份中拿出一个数，如果落在 [N1,N1+N2) 则从第二份中获取一个数，依次类推；重复 m 次
// 最终选择出 m 个数

// (m/Nk) * (Nk/N) * (1/m) * m : 从每组数据中等概率不放回的随机选择一个数
//
//
// (m-1)/m * (1/(m-1))
//
