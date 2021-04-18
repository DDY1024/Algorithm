package main

import "fmt"

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minOperations(nums []int) int {
	n, cnt := len(nums), 0
	if n == 1 {
		return 0
	}
	cur := nums[0]
	for i := 1; i < n; i++ {
		cnt += maxInt(cur+1, nums[i]) - nums[i]
		cur = maxInt(cur+1, nums[i])
	}
	return cnt
}

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	ans := make([]int, 0, n)
	xor, mask := 0, (1<<uint(maximumBit))-1
	for i := 0; i < n; i++ {
		xor ^= nums[i]
	}
	for i := n - 1; i >= 0; i-- {
		ans = append(ans, (xor|mask)^xor)
		xor ^= nums[i]
		return ans
	}
}

func main() {
	fmt.Println(3 + 4>>1)
}

type Item struct {
	enter int
	cost  int
	idx   int
}

type MinPQ []*Item

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Less(i, j int) bool {
	if pq[i].cost == pq[j].cost {
		return pq[i].idx < pq[j].idx
	}
	return pq[i].cost < pq[j].cost
}

func (pq MinPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinPQ) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *MinPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *MinPQ) Top() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	return (*pq)[0]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getOrder(tasks [][]int) []int {
	n := len(tasks)
	items := make([]*Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, &Item{
			enter: tasks[i][0],
			cost:  tasks[i][1],
			idx:   i,
		})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].enter <= items[j].enter
	})
	ans := make([]int, 0, n)
	hp := make(MinPQ, 0, n)
	cur, idx := items[0].enter, 0
	for idx < n && items[idx].enter <= cur {
		heap.Push(&hp, items[idx])
		idx++
	}
	for hp.Len() > 0 {
		item := heap.Pop(&hp).(*Item)
		ans = append(ans, item.idx)
		cur += item.cost
		for idx < n && items[idx].enter <= cur {
			heap.Push(&hp, items[idx])
			idx++
		}
	}
	return ans
}
