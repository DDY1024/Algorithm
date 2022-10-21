package main

import "container/list"

type Item struct {
	key int
	val int
}

type LRUCache struct {
	visit *list.List
	index map[int]*list.Element
	cap   int
	size  int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		visit: list.New(),
		index: make(map[int]*list.Element, capacity),
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

	if this.size > this.cap { // evict 元素驱逐
		this.size--
		item = this.visit.Back().Value.(*Item)
		delete(this.index, item.key)
		this.visit.Remove(this.visit.Back())
	}
}
