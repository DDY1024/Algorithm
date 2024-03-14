package main

// https://oi-wiki.org/basic/prefix-sum/
// https://www.cnblogs.com/COLIN-LIGHTNING/p/8436624.html

// 前缀和
//	一维：sum[i] = sum[i-1] + arr[i]
//  二维：sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + arr[i][j]

// 树上边权和：sum{x->y} = sum{root->x} + sum{root->y} - 2*sum{root->lca(x,y)}
// 树上点权和：sum{x->y} = sum{root->x} + sum{root->y} - sum{root->lca(x,y)} - sum{root->parent(lca(x,y))}

// 一维差分
// 		对于区间 [l, r] + val；则 d[l] += val, d[r+1] -= val；求解前缀和即为单点元素值

// 二维差分
//		对于子矩阵 [x1,y1] -> [x2,y2] + val；则 d[x1][y1] += val, d[x2+1][y1] -= val, d[x1][y2+1] -= val, d[x2+1][y2+1] += val
//		对于矩阵单个元素 arr[i][j] = sun{ d[x][y] }，其中 0 <= x <= i, 0 <= y <= j

// 差分数组前缀和
// 	一维差分数组前缀和
//  二维差分数组前缀和
