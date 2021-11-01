package main

// 线段树优化点
// 1. 在叶子节点处无需下放懒惰标记（其实完全下放以后 add 和 sum 是相等的），所以懒惰标记可以不下传到叶子节点
// 2. 下放懒惰标记单独抽出函数 push_down，向上更新父节点单独抽出函数 push_up
// TODO: 3. 标记永久化？？？ 可持久化的写法

// 改段求段模型
var lchild = func(idx int) int { return idx << 1 }
var rchild = func(idx int) int { return (idx << 1) | 1 }

func buildTwo(l, r, idx int, arr []int, add []int, sum []int) {
	if l == r {
		sum[idx] = arr[l]
		add[idx] = 0
		return
	}
	mid := (l + r) >> 1
	buildTwo(l, mid, lchild(idx), arr, add, sum)
	buildTwo(mid+1, r, rchild(idx), arr, add, sum)
	pushUpTwo(idx, sum)
}

func pushUpTwo(idx int, sum []int) {
	sum[idx] = sum[lchild(idx)] + sum[rchild(idx)]
}

func pushDownTwo(l, idx int, add []int, sum []int) {
	if add[idx] > 0 {
		add[lchild(idx)] += add[idx]
		add[rchild(idx)] += add[idx]
		sum[lchild(idx)] += add[idx] * (l - l>>1) // [l, (l+r)>>1]
		sum[rchild(idx)] += add[idx] * (l >> 1)   // [(l+r)>>1+1, r]
		add[idx] = 0
	}
}

// query 时只需要做 push_down 操作
// 由于 update 操作时祖先节点的 sum 值已经更新过了，query 时只处理子区间，因此只需要做 push_down 操作
func queryTwo(L, R, l, r, idx int, add []int, sum []int) int {
	if L <= l && r <= R {
		return sum[idx]
	}
	pushDownTwo(r-l+1, idx, add, sum)
	ret, mid := 0, (l+r)>>1
	if L <= mid {
		ret += queryTwo(L, R, l, mid, lchild(idx), add, sum)
	}
	if R > mid {
		ret += queryTwo(L, R, mid+1, r, rchild(idx), add, sum)
	}
	return ret
}

// update 时需要同时做 push_down 和 push_up 操作
// update 操作时已经更新了根节点路径上所有祖先节点的 sum 值
func updateTwo(L, R, l, r, delta, idx int, add []int, sum []int) {
	if L <= l && r <= R {
		add[idx] += delta
		sum[idx] += delta * (r - l + 1)
		return
	}
	pushDownTwo(r-l+1, idx, add, sum)
	mid := (l + r) >> 1
	if L <= mid {
		updateTwo(L, R, l, mid, delta, lchild(idx), add, sum)
	}
	if R > mid {
		updateTwo(L, R, mid+1, r, delta, rchild(idx), add, sum)
	}
	pushUpTwo(idx, sum)
}
