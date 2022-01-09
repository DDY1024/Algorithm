package main

import "container/list"

// leetcode 146 关于 lru cache 的实现，非并发数据结构实现
// 题目链接: https://leetcode-cn.com/problems/lru-cache/
// 题目大意
// LRU Cache 设计

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

func Constructor(capacity int) LRUCache {
	return LRUCache{
		visit: list.New(),
		index: make(map[int]*list.Element, 100000), // 可以随便给一个初始容量
		cap:   capacity,
		size:  0,
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.index[key]
	if !ok {
		return -1
	}

	this.visit.Remove(e)              // element
	e = this.visit.PushFront(e.Value) // value
	this.index[key] = e
	return e.Value.(*Item).val
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.index[key]; ok { // 同样影响 visit 中元素的顺序
		e.Value.(*Item).val = value
		this.visit.Remove(e)              // element
		e = this.visit.PushFront(e.Value) // value
		this.index[key] = e
		return
	}

	item := &Item{
		key: key,
		val: value,
	}

	e := this.visit.PushFront(item) // value
	this.index[key] = e
	this.size++

	if this.size > this.cap {
		this.size--
		item = this.visit.Back().Value.(*Item)
		delete(this.index, item.key)
		this.visit.Remove(this.visit.Back()) // element
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
