package heap

type Item struct {
	val int
	idx int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

// 自定义优先级比较规则
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].val == pq[j].val {
		return pq[i].idx > pq[j].idx
	}
	return pq[i].val > pq[j].val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Top() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	return (*pq)[0]
}
