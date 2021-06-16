package main

import (
	"fmt"
	"sort"
	"time"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minWastedSpace(packages []int, boxes [][]int) int {
	n, m := len(packages), len(boxes)
	for i := 0; i < m; i++ {
		sort.Ints(boxes[i]) // 应该不需要排序，默认是有序的
	}

	pSum := make([]int, n)
	sort.Ints(packages)
	pSum[0] = packages[0]
	for i := 1; i < n; i++ {
		pSum[i] = pSum[i-1] + packages[i]
	}

	// 1. 手写二分查找算法
	// var search = func(x int) int {
	// 	l, r, idx := 0, n-1, -1
	// 	for l <= r {
	// 		mid := l + (r-l)/2
	// 		if packages[mid] <= x {
	// 			idx = mid
	// 			l = mid + 1
	// 		} else {
	// 			r = mid - 1
	// 		}
	// 	}
	// 	return idx
	// }

	// 2. 利用标准库的二分查找能力 --> 熟悉标准库的操作能够节省不少时间
	var search = func(x int) int {
		return sort.Search(n, func(idx int) bool {
			return packages[idx] > x
		}) - 1
	}

	cost := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < len(boxes[i]); j++ {
			idx1 := search(boxes[i][j])
			if idx1 < 0 {
				continue
			}

			if j-1 >= 0 {
				idx2 := search(boxes[i][j-1])
				if idx2 < 0 {
					cost[i] += boxes[i][j]*(idx1+1) - pSum[idx1]
				} else {
					cost[i] += boxes[i][j]*(idx1-idx2) - (pSum[idx1] - pSum[idx2])
				}
			} else {
				cost[i] += boxes[i][j]*(idx1+1) - pSum[idx1]
			}
		}
	}

	mod := int(1e9 + 7)
	ans, ok := 0x3f3f3f3f3f3f3f3f, false
	for i := 0; i < m; i++ {
		if boxes[i][len(boxes[i])-1] >= packages[n-1] {
			ok = true
			ans = minInt(ans, cost[i])
		}
	}
	if !ok {
		return -1
	}
	return ans % mod
}

func main() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second)
		close(ch)
	}()
	<-ch
	fmt.Println("hello, world!")
}
