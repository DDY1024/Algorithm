package main

import "container/list"

// 题目链接：https://leetcode.cn/problems/lru-cache/description/?envType=study-plan-v2&envId=top-100-liked

type Item struct {
	key   int
	value int
}

type LRUCache struct {
	hash  map[int]*list.Element // key-value 映射
	visit *list.List            // 维护访问顺序
	cap   int                   // 容量
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		hash:  make(map[int]*list.Element),
		visit: list.New(),
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.hash[key]
	if !ok {
		return -1
	}

	this.visit.Remove(e)
	this.hash[key] = this.visit.PushFront(e.Value)
	return e.Value.(*Item).value
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.hash[key]; ok {
		e.Value.(*Item).value = value
		this.visit.Remove(e)
		this.hash[key] = this.visit.PushFront(e.Value)
		return
	}

	item := &Item{
		key:   key,
		value: value,
	}
	this.hash[key] = this.visit.PushFront(item)

	// 缓存淘汰策略
	if len(this.hash) > this.cap {
		key := this.visit.Back().Value.(*Item).key
		delete(this.hash, key)
		this.visit.Remove(this.visit.Back())
	}
}
