package main

import "container/list"

// 实现一个简单的 lru 缓存

type Item struct {
	key int
	val int
}

type LRUCache struct {
	visit *list.List
	index map[int]*list.Element // 直接利用 golang 标准库来实现双向链表
	cap   int
	size  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		visit: list.New(),
		index: make(map[int]*list.Element, 100000),
		cap:   capacity,
		size:  0,
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.index[key]
	if !ok {
		return -1
	}

	this.visit.Remove(e)
	e = this.visit.PushFront(e.Value)
	this.index[key] = e
	return e.Value.(*Item).val
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.index[key]; ok {
		e.Value.(*Item).val = value
		this.visit.Remove(e)
		e = this.visit.PushFront(e.Value)
		this.index[key] = e
		return
	}

	item := &Item{
		key: key,
		val: value,
	}

	e := this.visit.PushFront(item)
	this.index[key] = e
	this.size++

	if this.size > this.cap { // lru 驱逐策略
		this.size--
		item = this.visit.Back().Value.(*Item)
		delete(this.index, item.key)
		this.visit.Remove(this.visit.Back())
	}
}
