package timewheel

import (
	"sync"
	"time"
)

// 一个简单时间轮实现
// 注意点
// 1. add 操作 remain_time 不能大于时间轮窗口大小
// 2. 元素 expire < ticksPerWheel；例如 60 秒过期，申请 ticksPerWheel = 61

const (
	defaultTaskChanBufferSize = 300
	defaultSlotSize           = 300
)

type slot struct {
	elements map[interface{}]interface{}
}

type entry struct {
	k      interface{}
	v      interface{}
	expire int
}

func newSlot() *slot {
	s := &slot{}
	s.elements = make(map[interface{}]interface{}, defaultSlotSize)
	return s
}

func (s *slot) add(k, v interface{}) {
	s.elements[k] = v
}

func (s *slot) remove(k interface{}) {
	delete(s.elements, k)
}

type handler func(k, v interface{})

type TimeWheel struct {
	tickDuration     time.Duration         // 设置 tick 递增的时间间隔
	ticksPerWheel    int                   // 时间轮 slot 数量
	currentTickIndex int                   // 当前 tick 索引下标
	ticker           *time.Ticker          // timer
	onTick           handler               // 处理函数
	wheel            []*slot               // slot list
	indicator        map[interface{}]*slot // 记录每个 key 所属 slot
	sync.RWMutex

	taskChan chan *entry
	quitChan chan struct{}
}

func New(tickDuration time.Duration, ticksPerWheel int, f handler) *TimeWheel {
	if tickDuration < 1 || ticksPerWheel < 1 || nil == f {
		return nil
	}

	t := &TimeWheel{
		tickDuration:     tickDuration,
		ticksPerWheel:    ticksPerWheel,
		onTick:           f,
		currentTickIndex: 0,
		// 带缓冲区的 channel
		// taskChan:         make(chan *entry, defaultTaskChanBufferSize),
		// 不带缓冲区的 channel
		taskChan: make(chan *entry),
		quitChan: make(chan struct{}),
	}
	t.indicator = make(map[interface{}]*slot, t.ticksPerWheel*defaultSlotSize)
	t.wheel = make([]*slot, ticksPerWheel)
	for i := 0; i < ticksPerWheel; i++ {
		t.wheel[i] = newSlot()
	}
	return t
}

// 异步处理元素添加
func (t *TimeWheel) Add(k, v interface{}) {
	t.taskChan <- &entry{k, v, t.ticksPerWheel - 1} // 不能向当前 slot 添加元素，会导致误删，因此默认要 -1
}

// 代码实现并没有针对超过一轮的过期时间进行特殊处理，建议每次 add 要小于 < t.ticksPerWheel
func (t *TimeWheel) AddWithRemainingTime(k, v interface{}, remainingTime int) {
	t.taskChan <- &entry{k, v, remainingTime}
}

func (t *TimeWheel) slotRemove(k interface{}) {
	if v, ok := t.indicator[k]; ok {
		v.remove(k)
	}
}

func (t *TimeWheel) getCurrentTickIndex() int {
	t.RLock()
	defer t.RUnlock()
	return t.currentTickIndex
}

func (t *TimeWheel) Stop() {
	close(t.quitChan)
}

func (t *TimeWheel) run() {
	ticker := time.NewTicker(t.tickDuration)
	defer ticker.Stop()

	for {
		select {
		case <-t.quitChan:
			break
		case <-t.ticker.C:
			slot := t.wheel[t.currentTickIndex]
			for k, v := range slot.elements {
				delete(slot.elements, k)
				delete(t.indicator, k)
				t.onTick(k, v)
			}
			t.currentTickIndex = (t.currentTickIndex + 1) % t.ticksPerWheel
		case element := <-t.taskChan:
			t.slotRemove(element.k)
			idx := (t.getCurrentTickIndex() + element.expire) % t.ticksPerWheel
			slot := t.wheel[idx]
			slot.add(element.k, element.v)
			t.indicator[element.k] = slot
		}
	}
}
