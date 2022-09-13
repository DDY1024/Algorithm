package main

// 参考资料：https://zhuanlan.zhihu.com/p/106118909

const maxn = 10010 // 一般为 4*n + 10，即 n<<2 + 10

var lchild = func(idx int) int { return idx << 1 }
var rchild = func(idx int) int { return (idx << 1) | 1 }

func pushUp(idx int, sum []int) {
	sum[idx] = sum[lchild(idx)] + sum[rchild(idx)]
}

func buildTree(l, r, idx int, arr []int, sum []int) {
	if l == r {
		sum[idx] = arr[l]
		return
	}

	mid := (l + r) >> 1
	buildTree(l, mid, lchild(idx), arr, sum)
	buildTree(mid+1, r, rchild(idx), arr, sum)
	pushUp(idx, sum)
}

// [l, r]: 查询区间
// [s, t]: 当前线段树节点代表的区间
// idx: 节点索引编号
// 例如: query(l, r, 1, n, 1, sum)
func query(l, r, s, t, idx int, sum []int) int {
	if l <= s && t <= r { // 查询区间包含节点区间，直接返回该节点区间的统计结果
		return sum[idx]
	}

	mid, ret := (s+t)>>1, 0
	if l <= mid {
		ret += query(l, r, s, mid, lchild(idx), sum)
	}
	if r > mid {
		ret += query(l, r, mid+1, t, rchild(idx), sum)
	}

	return ret
}

// 单点更新
// 位置参数: pos
// 增量: d
func update(pos, d, s, t, idx int, sum []int) {
	if s == t { // 叶子节点
		sum[idx] += d
		return
	}

	mid := (s + t) >> 1
	if pos <= mid {
		update(pos, d, s, mid, lchild(idx), sum)
	} else {
		update(pos, d, mid+1, t, rchild(idx), sum)
	}
	pushUp(idx, sum)
}
