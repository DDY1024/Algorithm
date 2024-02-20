package main

import (
	"sort"
)

// 最大异或和：0/1 字典树求解
// https://leetcode.com/contest/weekly-contest-221/problems/maximum-xor-with-an-element-from-array/

type Trie struct {
	val      int
	children [2]*Trie
}

func calcBit(x, i int) int {
	if x&(1<<uint(i)) > 0 {
		return 1
	}
	return 0
}

func insertTrie(root *Trie, val int) {
	for i := 30; i >= 0; i-- { // 1024 * 1024 * 1024 > 10^9
		bit := calcBit(val, i)
		if root.children[bit] == nil {
			root.children[bit] = &Trie{-1, [2]*Trie{}}
		}
		root = root.children[bit]
	}
	root.val = val
}

func maxXor(root *Trie, val int) int {
	for i := 30; i >= 0; i-- {
		bit := calcBit(val, i)
		if root.children[bit^1] != nil {
			root = root.children[bit^1]
		} else {
			root = root.children[bit]
		}
	}
	return val ^ root.val
}

type Query struct {
	x, m, idx int
}

// 离线处理查询，利用字典树做最大异或和运算
func maximizeXor(nums []int, queries [][]int) []int {
	n, qL := len(nums), len(queries)
	qnds := make([]Query, 0, qL)
	for i := 0; i < qL; i++ {
		qnds = append(qnds, Query{queries[i][0], queries[i][1], i})
	}
	sort.Slice(qnds, func(i, j int) bool {
		return qnds[i].m < qnds[j].m
	})
	sort.Ints(nums)

	root := &Trie{}
	nIdx := -1
	ans := make([]int, qL)
	for i := 0; i < qL; i++ { // 离线处理查询操作，直接排序处理
		for nIdx+1 < n && nums[nIdx+1] <= qnds[i].m {
			nIdx++
			insertTrie(root, nums[nIdx])
		}
		if nIdx < 0 {
			ans[qnds[i].idx] = -1
		} else {
			ans[qnds[i].idx] = maxXor(root, qnds[i].x)
		}
	}
	return ans
}

// 在线查询处理：不处理 query 顺序，一次性遍历所有节点建立 trie。但是在利用 trie 做最大异或查询
// 我们是需要知道当前走到的子树其对应的叶子节点是否存在 <= mi 的元素的。因此在 trie 的基础上
// 我们需要维护树中每个节点所代表子树节点中的最小值来引导 mi 在每个分叉口是选择 0 或 1，这便是
// 我们理解的在线算法，C++ 代码参考如下：
// https://mp.weixin.qq.com/s/V09t-PT-RAbagtTn2RJmIw
/*
const int INF = 0x3f3f3f3f;
const int MXN = 1e5 * 32 + 11;
class Solution {
public:
    int trie[MXN][2], siz, Min[MXN];
    int newNode() {  // 卧槽，这是静态字典树的写法
        ++ siz;
        trie[siz][0] = trie[siz][1] = -1;
        Min[siz] = INF;
        return siz;
    }
    void add(int x) {
        int rt = 0;
        for(int i = 31; i >= 0; --i) {
            Min[rt] = min(Min[rt], x);
            int now = (x >> i) & 1;
            if(trie[rt][now] == -1) trie[rt][now] = newNode();
            rt = trie[rt][now];
        }
        Min[rt] = x;
    }
    int check(int x, int m) {
        int rt = 0, ans = -1;
        for(int i = 31; i >= 0; --i) {
            int now = (x >> i) & 1;
            int vis[2] = {0, 0};
            if(trie[rt][0] != -1 && Min[trie[rt][0]] <= m) vis[0] = 1;
            if(trie[rt][1] != -1 && Min[trie[rt][1]] <= m) vis[1] = 1;
            if(vis[!now]) {
                if(ans == -1) ans = 0;
                ans += (1 << i);
                rt = trie[rt][!now];
            }else if(vis[now]) {
                if(ans == -1) ans = 0;
                rt = trie[rt][now];
            }else {
                ans = -1;
                break;
            }
        }
        return ans;
    }
    vector<int> maximizeXor(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size(), m = queries.size();
        vector<int> ans(m, -1);
        siz = 0;
        trie[siz][0] = trie[siz][1] = -1;
        Min[siz] = INF;
        for(int i = 0; i < n; ++i) add(nums[i]);
        for(int i = 0; i < m; ++i) {
            ans[i] = check(queries[i][0], queries[i][1]);
        }
        return ans;
    }
};
*/
