package main

// 单调队列实现：考虑到 oj 对于时间的要求，不做额外的扩展操作
type Entry struct {
	val int
	idx int
}

// 自定义函数
func (e *Entry) Less(oe *Entry) bool {
	return e.val <= oe.val
}

type MQueue struct {
	entries []*Entry
	size    int
}

func NewMQueue(cap int) *MQueue {
	return &MQueue{
		entries: make([]*Entry, 0, cap),
		size:    0,
	}
}

func (mq *MQueue) Len() int {
	return mq.size
}

func (mq *MQueue) Back() *Entry {
	if mq.Len() == 0 {
		return nil
	}
	return mq.entries[mq.size-1]
}

func (mq *MQueue) PopBack() {
	if mq.Len() == 0 {
		return
	}
	mq.size--
	mq.entries = mq.entries[:mq.size]
}

func (mq *MQueue) PushBack(val, idx int) {
	if mq.Len() == 0 {
		mq.size++
		mq.entries = append(mq.entries, &Entry{
			val,
			idx,
		})
		return
	}

	e := &Entry{val, idx}
	for mq.Len() > 0 && mq.Back().Less(e) {
		mq.PopBack()
	}
	mq.entries = append(mq.entries, e)
	mq.size++
}

func (mq *MQueue) Front() *Entry {
	if mq.Len() == 0 {
		return nil
	}
	return mq.entries[0]
}

func (mq *MQueue) PopFront() {
	if mq.Len() == 0 {
		return
	}
	mq.entries = mq.entries[1:]
	mq.size--
}

// 进行窗口裁剪时调用
func (mq *MQueue) CutLessIdx(idx int) {
	for mq.Len() > 0 {
		if mq.Front().idx < idx {
			mq.RemoveFront()
		} else {
			break
		}
	}
}
