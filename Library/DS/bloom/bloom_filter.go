package main

import (
	"bytes"
	"encoding/binary"
	"math"
	"sync"
)

// 参考文档
// 1. https://juejin.cn/post/6844904007790673933
// 2. https://zh.wikipedia.org/wiki/%E5%B8%83%E9%9A%86%E8%BF%87%E6%BB%A4%E5%99%A8
// 3. go-zero bloom_filter 实现

type Options struct {
	locker sync.Locker
}

type Option func(opt *Options)

// func WithGoroutineSafe() Option {
// 	return func(opt *Options) {
// 		opt.locker = &sync.RWMutex{}
// 	}
// }

type BloomFilter struct {
	m      uint64 // bitmap 容量
	k      uint64 // hash 函数个数
	b      *Bitmap
	locker sync.Locker
}

func NewBloomFilter(m, k uint64, opts ...Option) *BloomFilter {
	opt := Options{
		locker: &sync.RWMutex{},
	}

	for _, o := range opts {
		o(&opt)
	}

	return &BloomFilter{
		m:      m,
		k:      k,
		b:      NewBitMap(m),
		locker: opt.locker,
	}
}

// EstimateParameters estimates m and k from n and p
func EstimateParameters(n uint64, p float64) (m uint64, k uint64) {
	m = uint64(math.Ceil(-1 * float64(n) * math.Log(p) / (math.Ln2 * math.Ln2)))
	k = uint64(math.Ceil(math.Ln2 * float64(m) / float64(n)))
	return m, k
}

// NewWithEstimates creates a new BloomFilter with n and fp.
// n is the capacity of the BloomFilter
// fp is the tolerated error rate of the BloomFilter（允许错误率）
func NewWithEstimates(n uint64, fp float64, opts ...Option) *BloomFilter {
	m, k := EstimateParameters(n, fp)
	return NewBloomFilter(m, k, opts...)
}

func (bf *BloomFilter) Add(val string) {
	bf.locker.Lock()
	defer bf.locker.Unlock()

	data := []byte(val)
	// k 个槽位 hash 计算，参考 go-zero 实现, murmur3 加偏移量 i
	for i := uint64(0); i < bf.k; i++ {
		pos := Hash(append(data, byte(i)))
		bf.b.Set(pos % bf.m)
	}
}

func (bf *BloomFilter) Contains(val string) bool {
	bf.locker.Lock()
	defer bf.locker.Unlock()

	data := []byte(val)
	for i := uint64(0); i < bf.k; i++ {
		pos := Hash(append(data, byte(i)))
		if !bf.b.IsSet(pos % bf.m) {
			return false
		}
	}
	return true
}

// 布隆过滤器序列化
func (bf *BloomFilter) EncodeData() []byte {
	bf.locker.Lock()
	defer bf.locker.Unlock()

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, bf.m)
	binary.Write(buf, binary.LittleEndian, bf.k)
	buf.Write(bf.b.Data())
	return buf.Bytes()
}

func DecodeFromData(data []byte, opts ...Option) *BloomFilter {
	opt := Options{
		locker: &sync.Mutex{},
	}
	for _, o := range opts {
		o(&opt)
	}
	b := &BloomFilter{
		locker: opt.locker,
	}

	reader := bytes.NewReader(data)
	binary.Read(reader, binary.LittleEndian, &b.m) // 8byte
	binary.Read(reader, binary.LittleEndian, &b.k) // 8byte
	b.b = NewBitMapFromData(data[8+8:])
	return b
}
