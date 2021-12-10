package main

// 参考: https://oi-wiki.org/basic/prefix-sum/
// 参考: https://www.cnblogs.com/COLIN-LIGHTNING/p/8436624.html

// 1. 前缀和
// 一维前缀和: pre[i] = pre[i-1] + arr[i]
// 二维前缀和: pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + arr[i][j]
// 高维前缀和: 基于容斥原理计算
// 树上前缀和 sum(i) 为节点 i 到根节点的权值总和
// 若是点权，则路径 x --> y 权值和为 sum(x) + sum(y) - sum(lca(x,y)) - sum(parent(lca(x,y)))
// 若是边权，则路径 x --> y 权值和为 sum(x) + sum(y) - 2*sum(lca(x,y))

// 2. 差分
// 具体参考: https://oi-wiki.org/basic/prefix-sum/
//
// a. 区间差分 --> 差分数组推导区间前缀和
// 	  对于区间 [l,r] 增加 k，则对应的差分操作为 b[l] += k, b[r+1] -= k。全部操作结束后，b 的前缀和即第 i 个元素最终的值 sigma(b[1]+b[2]+...+b[i])
//    询问区间元素和: 差分数组 --> 每个元素 --> 前缀和 --> 区间和
//
// b. 树上差分 --> 比较难，遇到该类题目时再做总结
