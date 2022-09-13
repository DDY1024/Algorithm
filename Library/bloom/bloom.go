package main

import (
	"bytes"
	"encoding/binary"
	"math"
	"sync"
)

// bloom filter 实现参考下述文档
// 1. https://juejin.cn/post/6844904007790673933
// 2. https://zh.wikipedia.org/wiki/%E5%B8%83%E9%9A%86%E8%BF%87%E6%BB%A4%E5%99%A8
// 3. go-zero bloom_filter 实现
// 4. 计算工具: https://krisives.github.io/bloom-calculator/

type BloomFilter struct {
	m      uint64 // 位数组大小
	k      uint64 // hash 函数个数
	b      *Bitmap
	locker sync.Locker
}

type Option func(bf *BloomFilter)

func WithGoroutineSafe() Option {
	return func(bf *BloomFilter) {
		bf.locker = &sync.RWMutex{}
	}
}

func NewBloomFilter(m, k uint64, opts ...Option) *BloomFilter {
	bf := &BloomFilter{
		m: m,
		k: k,
		b: NewBitMap(m),
	}
	for _, o := range opts {
		o(bf)
	}
	return bf
}

// n: 元素个数
// p: 误判率
func EstimateParameters(n uint64, p float64) (m uint64, k uint64) {
	m = uint64(math.Ceil(-1 * float64(n) * math.Log(p) / (math.Ln2 * math.Ln2)))
	k = uint64(math.Ceil(math.Ln2 * float64(m) / float64(n)))
	return m, k
}

func NewWithEstimates(n uint64, fp float64, opts ...Option) *BloomFilter {
	m, k := EstimateParameters(n, fp)
	return NewBloomFilter(m, k, opts...)
}

func (bf *BloomFilter) Add(val string) {
	bf.locker.Lock()
	defer bf.locker.Unlock()

	data := []byte(val)
	// murmur3 hash + 偏移量 i，构造 k 个 hash 函数
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

// 元信息 + bitmap byte data
func (bf *BloomFilter) EncodeData() []byte {
	bf.locker.Lock()
	defer bf.locker.Unlock()

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, bf.m)
	binary.Write(buf, binary.LittleEndian, bf.k)
	buf.Write(bf.b.Data())
	return buf.Bytes()
}

// 反序列化
func DecodeFromData(data []byte, opts ...Option) *BloomFilter {
	bf := &BloomFilter{}
	for _, o := range opts {
		o(bf)
	}
	reader := bytes.NewReader(data)
	binary.Read(reader, binary.LittleEndian, &bf.m)
	binary.Read(reader, binary.LittleEndian, &bf.k)
	bf.b = NewBitMapFromData(data[8+8:])
	return bf
}
